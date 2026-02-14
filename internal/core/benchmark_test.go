package core

import (
	"testing"
	"time"
)

func BenchmarkVerifyServiceProvenanceHash(b *testing.B) {
	service := &VerificationService{}

	b.ResetTimer()

	// Benchmark provenance hash generation
	for i := 0; i < b.N; i++ {
		service.generateProvenanceHash("license123", "tool456", "0xUser")
	}

	b.ReportAllocs()
}

func BenchmarkVerifyServiceSignatureValidation(b *testing.B) {
	service := &VerificationService{}

	b.ResetTimer()

	// Benchmark signature validation
	for i := 0; i < b.N; i++ {
		timestamp := time.Now().Unix() - int64(i%300)
		service.ValidateSignature("0xUser", timestamp, "sig")
	}

	b.ReportAllocs()
}

func BenchmarkGenerateProvenanceHash(b *testing.B) {
	service := &VerificationService{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.generateProvenanceHash("license123", "tool456", "0xUser")
	}

	b.ReportAllocs()
}