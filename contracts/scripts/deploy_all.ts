import { expect } from "chai";
import { log } from "console";
import { network } from "hardhat";
const { ethers } = await network.connect();

 // Constants for precision
  const MIN_STAKE = ethers.parseEther("10");
  const LICENSE_PRICE = ethers.parseEther("0.01");
async function deployAll() {
    let [owner, dev, user, attacker, backendSigner, treasury] = await ethers.getSigners();

    // Deploy SkillToken
    const SkillToken = await ethers.getContractFactory("SkillToken");
    let skillToken = await SkillToken.deploy(owner.address, ethers.parseEther("1000000"), ethers.parseEther("10000000"));
    await skillToken.waitForDeployment();

    // Deploy StakingNFT
    const StakingNFT = await ethers.getContractFactory("StakingNFT");
    let stakingNFT = await StakingNFT.deploy(await skillToken.getAddress(), MIN_STAKE);
    await stakingNFT.waitForDeployment();

    // Deploy LicenseNFT
    const LicenseNFT = await ethers.getContractFactory("LicenseNFT");
    
    let licenseNFT = await LicenseNFT.deploy(backendSigner.address, "https://api.skillchain.xyz/license/");
    await licenseNFT.waitForDeployment();
    log(backendSigner.address)
    // Fund accounts
    await skillToken.connect(owner).transfer(dev.address, ethers.parseEther("1000"));
    await skillToken.connect(owner).transfer(user.address, ethers.parseEther("500"));
    await skillToken.connect(owner).transfer(attacker.address, ethers.parseEther("500"));
  }  deployAll()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });