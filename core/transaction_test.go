package core

import (
	"testing"

	"github.com/jomsdev/digital/common"
)

func MockTransactions() []Transaction {
	var transactions []Transaction

	keys := MockKeyPairs()

	tr0 := NewTransaction(keys[0].Public, keys[0].Public, CoinCreation, []byte("First CoinCreation"))
	tr0.SignTransaction(&keys[0])

	tr1 := NewTransaction(keys[0].Public, keys[0].Public, CoinCreation, []byte("Second CoinCreation"))
	tr1.SignTransaction(&keys[0])

	tr2 := NewTransaction(keys[0].Public, keys[1].Public, Payment, []byte("Giving one to the first user"))
	tr2.SignTransaction(&keys[0])

	tr3 := NewTransaction(keys[0].Public, keys[2].Public, Payment, []byte("Giving one to the second user"))
	tr3.SignTransaction(&keys[0])

	tr4 := NewTransaction(keys[1].Public, keys[2].Public, Payment, []byte("Giving one to the second user"))
	tr4.SignTransaction(&keys[1])
	transactions = append(transactions, *tr0, *tr1, *tr2, *tr3, *tr4)

	return transactions
}
func TestTransactionVerification(t *testing.T) {

	kp := MockKeyPairs()[0]
	randomPayload := []byte(common.RandStringBytesMaskImprSrc(1024))

	tr := NewTransaction(kp.Public, nil, Payment, randomPayload)

	tr.SignTransaction(&kp)

	res, err := tr.VerifyTransaction()

	if err != nil || !res {
		t.Errorf("error verifying a correct transaction")
	}
}

func TestTransactionVerificationFalseSignature(t *testing.T) {

	kp := MockKeyPairs()[0]
	kp2 := MockKeyPairs()[1]
	randomPayload := []byte(common.RandStringBytesMaskImprSrc(1024))

	tr := NewTransaction(kp.Public, nil, Payment, randomPayload)

	tr.SignTransaction(&kp2)

	res, err := tr.VerifyTransaction()

	if err != nil || res {
		t.Errorf("error verifying a fake transaction")
	}
}
