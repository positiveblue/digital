package crypto

import (
	"testing"
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
