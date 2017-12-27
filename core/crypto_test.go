package core

import (
	"testing"

	"github.com/jomsdev/digital/common"
)

func MockKeyPairs() []Keypair {
	var keys []Keypair
	var kp1, kp2, kp3, kp4, kp5 Keypair
	kp1.FromBase64("JSbuw9jo56ilt+EFJb9j1W2zZjmVPmNqqYYGdRKVcPppYHdoGXUuCGovZ+SeT++S1n6vOHTY2VSs1sCiGNegzg==", "yQyi13MkAcnEm++yUNrfo21OwBzPiKhNvrjZLrFQGOU=")
	kp2.FromBase64("W2TBZ4r9jutVE42OV+MnwnVnxNy+cClso2Cp5EKQdskdreK9xOZwW9e7qsmBi4osTNaRjR7YXHKX2wNaho8osw==", "Vi59U7sTgzqr4tEsp+oDh1R0yjPPah9ztv3Vg6G+fLU=")
	kp3.FromBase64("6NgAreHtGZDR0T9qUtKibzcXEO0PGXIzWOB8bJuzYmO1sCUSicPfW4NPlj12USwTZ0obYDs7yppmixdobniS6w==", "BK15D6kgLvUl0FLM8xYBRyAdvs283cLMF6cSiG+rjbE=")
	kp4.FromBase64("rg0tZ8jjpz7Rf6QB/PgNeZxU4ajJ6gDJRN/6rGBs9w6L49z43o1cbez/CIGh+qcflXa7BgXf+JksTuW0enxmZw==", "L9c4TXRwsxj0rL6inoiy6h5wfpk3fLvvgaSN5cVZSoI=")
	kp5.FromBase64("x+YazUImI0DdU9pOqfwMiUKNud9PMCdUTEzhxfl5oPDTZ7U3tVQrZx54dXybx2X5FLlCd891PU4qYgHnbCFTIA==", "JXaCpzCFBpgZr/Je2kMjqTT40r+saNtqYmRmC+jOOig=")
	keys = append(keys, kp1, kp2, kp3, kp4, kp5)

	return keys
}

func TestGenerateNewKeypair(t *testing.T) {
	kp, err := GenerateNewKeypair()
	if err != nil {
		t.Errorf("crypto: error generating a public keypair. %s", err.Error())
	}

	//Check the number of bytes of the two keys
	if len(kp.Public) != 64 || len(kp.Private) != 32 {
		t.Errorf("crypto: error generating keys for P256, keys with bad length")
	}
}

func TestSignAndVerify(t *testing.T) {
	// True Signatures
	for testCases := 5; testCases > 0; testCases-- {
		kp, err := GenerateNewKeypair()
		message := common.RandStringBytesMaskImprSrc(64)
		hash := []byte(message)

		signature, err := kp.Sign(hash)
		if err != nil {
			t.Errorf("crypto: error signing a message")
		}

		if !Verify(kp.Public, signature, hash) {
			t.Errorf("crypto: error verifying a valid signature")
		}
	}
}

func TestSignAndVerifyFalseSignatures(t *testing.T) {
	// Fake Signatures
	for testCases := 10; testCases > 0; testCases-- {
		kp, _ := GenerateNewKeypair()
		kp2, _ := GenerateNewKeypair()
		message := common.RandStringBytesMaskImprSrc(64)

		signature, err := kp.Sign([]byte(message))
		if err != nil {
			t.Errorf("crypto: error signing a message")
		}

		if Verify(kp2.Public, signature, []byte(message)) {
			t.Errorf("crypto: error verifying a valid signature")
		}
	}
}
