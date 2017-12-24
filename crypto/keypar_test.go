package crypto

import (
	"testing"
)

func TestGenerateNewKeypair(t *testing.T) {
	kp, err := GenerateNewKeypair()
	if err != nil {
		t.Errorf("crypto: error generating a public keypair. %s", err.Error())
	}
}
