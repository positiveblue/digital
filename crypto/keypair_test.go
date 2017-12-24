package crypto

import (
	"testing"

	"github.com/jomsdev/digital/common"
)

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
	for testCases := 10; testCases > 0; testCases-- {
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
