import { expect } from "chai";
import { network } from "hardhat";
const { ethers } = await network.connect();

describe("SkillChain Marketplace - Comprehensive Tests", function () {
  let skillToken, stakingNFT, licenseNFT;
  let owner, dev, user, attacker, backendSigner, treasury;

  // Constants for precision
  const MIN_STAKE = ethers.parseEther("10");
  const LICENSE_PRICE = ethers.parseEther("0.01");

  beforeEach(async function () {
    [owner, dev, user, attacker, backendSigner, treasury] = await ethers.getSigners();

    // Deploy SkillToken
    const SkillToken = await ethers.getContractFactory("SkillToken");
    skillToken = await SkillToken.deploy(owner.address, ethers.parseEther("1000000"), ethers.parseEther("10000000"));
    await skillToken.waitForDeployment();

    // Deploy StakingNFT
    const StakingNFT = await ethers.getContractFactory("StakingNFT");
    stakingNFT = await StakingNFT.deploy(await skillToken.getAddress(), MIN_STAKE);
    await stakingNFT.waitForDeployment();

    // Deploy LicenseNFT
    const LicenseNFT = await ethers.getContractFactory("LicenseNFT");
    
    licenseNFT = await LicenseNFT.deploy(backendSigner.address, "https://api.skillchain.xyz/license/");
    await licenseNFT.waitForDeployment();

    // Fund accounts
    await skillToken.connect(owner).transfer(dev.address, ethers.parseEther("1000"));
    await skillToken.connect(owner).transfer(user.address, ethers.parseEther("500"));
    await skillToken.connect(owner).transfer(attacker.address, ethers.parseEther("500"));
  });

  /******************************************************************
   * STAKINGNFT CONTRACT - EDGE CASE & SECURITY TESTS
   ******************************************************************/
  describe("StakingNFT - Edge Cases & Security", function () {
    it("Should revert if listing without token approval", async function () {
      // DO NOT call approve
      await expect(
        stakingNFT.connect(dev).listTool(MIN_STAKE, "ipfs://tool1")
      ).to.be.revertedWithCustomError(skillToken, "ERC20InsufficientAllowance");
    });

    it("Should revert if listing with insufficient token balance", async function () {
      const poorDev = attacker; // Has 500 tokens, but we try to stake 1000
      await skillToken.connect(poorDev).approve(await stakingNFT.getAddress(), ethers.parseEther("1000"));
      await expect(
        stakingNFT.connect(poorDev).listTool(ethers.parseEther("1000"), "ipfs://tool1")
      ).to.be.revertedWithCustomError(skillToken, "ERC20InsufficientBalance");
    });

    it("Should correctly emit ToolListed event with all parameters", async function () {
      await skillToken.connect(dev).approve(await stakingNFT.getAddress(), MIN_STAKE);
      const tx = await stakingNFT.connect(dev).listTool(MIN_STAKE, "ipfs://QmUniqueHash123");
      await expect(tx)
        .to.emit(stakingNFT, "ToolListed")
        .withArgs(1, dev.address, MIN_STAKE, "ipfs://QmUniqueHash123");
    });

    it("Should allow ONLY owner to propose slash", async function () {
      await skillToken.connect(dev).approve(await stakingNFT.getAddress(), MIN_STAKE);
      await stakingNFT.connect(dev).listTool(MIN_STAKE, "ipfs://tool1");
      // Attacker tries to propose slash
      await expect(
        stakingNFT.connect(attacker).proposeSlash(0, ethers.parseEther("5"))
      ).to.be.revertedWithCustomError(stakingNFT, "OwnableUnauthorizedAccount");
    });

    it("Should revert if proposing slash exceeding current stake", async function () {
      await skillToken.connect(dev).approve(await stakingNFT.getAddress(), MIN_STAKE);
      await expect(stakingNFT.connect(dev).listTool(MIN_STAKE, "ipfs://tool1"))
      .to.emit(stakingNFT, "ToolListed");
      const excessSlash = MIN_STAKE + ethers.parseEther("1");
      await expect(
        stakingNFT.connect(owner).proposeSlash(1, excessSlash)
      ).to.be.revertedWith("Slash amount exceeds stake");
    });

    it("Should require timelock before executing slash", async function () {
      await skillToken.connect(dev).approve(await stakingNFT.getAddress(), MIN_STAKE);
      await stakingNFT.connect(dev).listTool(MIN_STAKE, "ipfs://tool1");
      const slashAmount = ethers.parseEther("3");
      // Propose slash
      await stakingNFT.connect(owner).proposeSlash(1, slashAmount);
      // Try to execute immediately (should fail - timelock not expired)
      await expect(
        stakingNFT.connect(owner).executeSlash(1, treasury.address)
      ).to.be.revertedWith("Timelock not expired");
    });

    it("Should allow delistTool ONLY by NFT owner", async function () {
      await skillToken.connect(dev).approve(await stakingNFT.getAddress(), MIN_STAKE);
      await stakingNFT.connect(dev).listTool(MIN_STAKE, "ipfs://tool1");
      // Attempt delisting by non-owner
      await expect(
        stakingNFT.connect(attacker).delistTool(1)
      ).to.be.revertedWith("Not token owner");
    });

    it("Should return stake on delistTool and emit ToolDelisted event", async function () {
      await skillToken.connect(dev).approve(await stakingNFT.getAddress(), MIN_STAKE);
      await stakingNFT.connect(dev).listTool(MIN_STAKE, "ipfs://tool1");
      const devBalanceBefore = await skillToken.balanceOf(dev.address);
      const tx = await stakingNFT.connect(dev).delistTool(1);
      const devBalanceAfter = await skillToken.balanceOf(dev.address);
      expect(devBalanceAfter - devBalanceBefore).to.equal(MIN_STAKE);
      expect(await stakingNFT.stakeAmount(1)).to.equal(0);
      await expect(tx)
        .to.emit(stakingNFT, "ToolDelisted")
        .withArgs(1, dev.address, MIN_STAKE);
    });

    it("Should emit StakeWithdrawn event when stake is returned", async function () {
      await skillToken.connect(dev).approve(await stakingNFT.getAddress(), MIN_STAKE);
      await stakingNFT.connect(dev).listTool(MIN_STAKE, "ipfs://tool1");
      const tx = await stakingNFT.connect(dev).delistTool(1);
      await expect(tx)
        .to.emit(stakingNFT, "StakeWithdrawn")
        .withArgs(1, dev.address, MIN_STAKE);
    });
  });

  /******************************************************************
   * LICENSENFT CONTRACT - EDGE CASE & SECURITY TESTS
   ******************************************************************/
  describe("LicenseNFT - Signature & Expiry Logic", function () {
    let toolId, expiresAt, futureExpiresAt, signature, expiredSignature;

    beforeEach(async function () {
      toolId = 42; // Arbitrary tool ID
      // Valid expiration (1 day from now)
      const currentTime = await latest()
      expiresAt = currentTime + 86400;
      // Far future expiration for another test
      futureExpiresAt = currentTime + 86400 * 30;
      // Expired timestamp (1 day ago)
      const expiredAt = currentTime - 86400;
const { chainId } = await ethers.provider.getNetwork();
	
      // Helper to create signature
      const createSignature = async (expiry, nonce, signer = backendSigner) => {
        const domain = {
          name: 'SkillChainLicense',
          version: '1',
          chainId,
          verifyingContract: await licenseNFT.getAddress(),
        };
        const types = {
          MintLicense: [
            { name: 'user', type: 'address' },
            { name: 'toolId', type: 'uint256' },
            { name: 'expiresAt', type: 'uint256' },
            { name: 'nonce', type: 'uint256' },
          ],
        };
        const value = { user: user.address, toolId, expiresAt: expiry, nonce };
        return await backendSigner.signTypedData(domain, types, value);
      };


      signature = await createSignature(expiresAt, 12345);
      expiredSignature = await createSignature(expiredAt, 99999);
    });

    it("Should revert with an expired backend signature", async function () {
      await expect(
        licenseNFT.connect(user).mintLicense(toolId, expiresAt, 12345, expiredSignature, { value: LICENSE_PRICE })
      ).to.be.revertedWith("Invalid signature"); // Because expired timestamp in signature
    });

    it("Should revert if signature is from wrong signer (impersonation attack)", async function () {
      const fakeSignature = await (await ethers.getSigner(attacker.address)).signMessage("I am not the backend");
      await expect(
        licenseNFT.connect(user).mintLicense(toolId, expiresAt, 12345, fakeSignature, { value: LICENSE_PRICE })
      ).to.be.revertedWith("Invalid signature");
    });

    it("Should revert if user tries to reuse a nonce (replay attack prevention)", async function () {
      // First mint succeeds
      await licenseNFT.connect(user).mintLicense(toolId, expiresAt, 12345, signature, { value: LICENSE_PRICE });
      // Attempt to mint again with same nonce and signature
      await expect(
        licenseNFT.connect(user).mintLicense(toolId, expiresAt, 12345, signature, { value: LICENSE_PRICE })
      ).to.be.revertedWith("Nonce already used");
    });

    it("Should allow minting a new license after old one expires", async function () {
      // Mint first license
      await licenseNFT.connect(user).mintLicense(toolId, expiresAt, 12345, signature, { value: LICENSE_PRICE });
      expect(await licenseNFT.isLicenseValid(user.address, toolId)).to.be.true;
      // Fast-forward past expiration
      await increase(86400 + 1); // Just over 1 day
      expect(await licenseNFT.isLicenseValid(user.address, toolId)).to.be.false;
      // New signature with new expiration and nonce
      const newExpiresAt = (await latest()) + 86400;
      const { chainId } = await ethers.provider.getNetwork();
      const newSignature = await (async () => {
        const domain = {
          name: 'SkillChainLicense',
          version: '1',
          chainId,
          verifyingContract: await licenseNFT.getAddress(),
        };
        const types = {
          MintLicense: [
            { name: 'user', type: 'address' },
            { name: 'toolId', type: 'uint256' },
            { name: 'expiresAt', type: 'uint256' },
            { name: 'nonce', type: 'uint256' },
          ],
        };
        const value = { user: user.address, toolId, expiresAt: newExpiresAt, nonce: 67890 };
        return await backendSigner.signTypedData(domain, types, value);
      })();
      // Should succeed
      await expect(
        licenseNFT.connect(user).mintLicense(toolId, newExpiresAt, 67890, newSignature, { value: LICENSE_PRICE })
      ).to.emit(licenseNFT, "LicenseMinted");
      expect(await licenseNFT.isLicenseValid(user.address, toolId)).to.be.true;
    });

    it("Should correctly report license validity at exact expiry moment", async function () {
      await licenseNFT.connect(user).mintLicense(toolId, expiresAt, 12345, signature, { value: LICENSE_PRICE });
      // Move to 1 second before expiry
      await increase(86400 - 1);
      expect(await licenseNFT.isLicenseValid(user.address, toolId)).to.be.true;
      // Move to exact expiry timestamp
      await increase(1);
      expect(await licenseNFT.isLicenseValid(user.address, toolId)).to.be.false;
    });

    it("Should only allow owner to update the trusted signer", async function () {
      await expect(
        licenseNFT.connect(attacker).updateSigner(attacker.address)
      ).to.be.revertedWithCustomError(licenseNFT, "OwnableUnauthorizedAccount");
      // Owner can update
      await expect(licenseNFT.connect(owner).updateSigner(attacker.address))
        .to.emit(licenseNFT, "SignerChanged")
        .withArgs(backendSigner.address, attacker.address);
      expect(await licenseNFT.trustedSigner()).to.equal(attacker.address);
    });
  });

  /******************************************************************
   * INTEGRATION & PRECISION TESTS
   ******************************************************************/
  describe("Integration & Precision", function () {
    it("Should maintain precise stake accounting after timelock slash execution", async function () {
      const largeStake = ethers.parseEther("123.456789");
      await skillToken.connect(dev).approve(await stakingNFT.getAddress(), largeStake);
      await stakingNFT.connect(dev).listTool(largeStake, "ipfs://precise");
      const initialStake = await stakingNFT.stakeAmount(1);
      expect(initialStake).to.equal(largeStake);
      // Propose slash
      const slashAmt = ethers.parseEther("12.345678");
      await stakingNFT.connect(owner).proposeSlash(1, slashAmt);
      // Fast-forward 3 days + 1 second to exceed timelock
      await increase(3 * 24 * 60 * 60 + 1);
      // Execute slash
      const treasuryBalanceBefore = await skillToken.balanceOf(treasury.address);
      await stakingNFT.connect(owner).executeSlash(1, treasury.address);
      const treasuryBalanceAfter = await skillToken.balanceOf(treasury.address);
      // Verify slash transferred correctly
      expect(treasuryBalanceAfter - treasuryBalanceBefore).to.equal(slashAmt);
      expect(await stakingNFT.stakeAmount(1)).to.equal(largeStake - slashAmt);
      // Delist remaining stake
      const devBalanceBefore = await skillToken.balanceOf(dev.address);
      await stakingNFT.connect(dev).delistTool(1);
      const devBalanceAfter = await skillToken.balanceOf(dev.address);
      // Check balance increase matches remaining stake precisely
      expect(devBalanceAfter - devBalanceBefore).to.equal(largeStake - slashAmt);
      expect(await stakingNFT.stakeAmount(0)).to.equal(0);
    });

    it("Should handle complete slash workflow: propose -> timelock -> execute", async function () {
      await skillToken.connect(dev).approve(await stakingNFT.getAddress(), MIN_STAKE);
      await stakingNFT.connect(dev).listTool(MIN_STAKE, "ipfs://tool0");
      // Propose slash
      const slashAmount = ethers.parseEther("5");
      await expect(stakingNFT.connect(owner).proposeSlash(1, slashAmount))
        .to.emit(stakingNFT, "SlashProposed")
        .withArgs(1, slashAmount);
      // Execute immediately should fail
      await expect(
        stakingNFT.connect(owner).executeSlash(1, treasury.address)
      ).to.be.revertedWith("Timelock not expired");
      // Fast-forward time
      await increase(3 * 24 * 60 * 60 + 1);
      // Now execution should succeed
      const treasuryBalanceBefore = await skillToken.balanceOf(treasury.address);
      await stakingNFT.connect(owner).executeSlash(1, treasury.address);
      const treasuryBalanceAfter = await skillToken.balanceOf(treasury.address);
      expect(treasuryBalanceAfter - treasuryBalanceBefore).to.equal(slashAmount);
      expect(await stakingNFT.stakeAmount(1)).to.equal(MIN_STAKE - slashAmount);
    });
    it("Should ListTotal and capture event", async function () {
      await skillToken.connect(dev).approve(await stakingNFT.getAddress(), MIN_STAKE);
      await expect(stakingNFT.connect(dev).listTool(MIN_STAKE, "ipfs://tool0"))
      .to.emit(stakingNFT, "ToolListed")
      // Propose slash
      const slashAmount = ethers.parseEther("5");
      await expect(stakingNFT.connect(owner).proposeSlash(1, slashAmount))
        .to.emit(stakingNFT, "SlashProposed")
        .withArgs(1, slashAmount);
    })
  });
});

async function increase(time: Number){
  // Fast-forward time to satisfy hold period
    await ethers.provider.send("evm_increaseTime", [time]);
    await ethers.provider.send("evm_mine", []);
}
async function latest(): Promise<number>{
  return await ethers.provider.getBlock("latest").then(block => block.timestamp);
}