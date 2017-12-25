package core

import (
	"testing"

	"github.com/jomsdev/digital/common"
)

func TestTransactionVerification(t *testing.T) {

	kp, _ := GenerateNewKeypair()
	randomPayload := []byte(common.RandStringBytesMaskImprSrc(1024))

	tr := NewTransaction(kp.Public, nil, Payment, randomPayload)

	tr.SignTransaction(kp)

	res, err := tr.VerifyTransaction()

	if err != nil || !res {
		t.Errorf("error verifying a correct transaction")
	}
}

func TestTransactionVerificationFalseSignature(t *testing.T) {

	kp, _ := GenerateNewKeypair()
	kp2, _ := GenerateNewKeypair()
	randomPayload := []byte(common.RandStringBytesMaskImprSrc(1024))

	tr := NewTransaction(kp.Public, nil, Payment, randomPayload)

	tr.SignTransaction(kp2)

	res, err := tr.VerifyTransaction()

	if err != nil || res {
		t.Errorf("error verifying a fake transaction")
	}
}
