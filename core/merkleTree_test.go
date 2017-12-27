package core

import (
	"testing"
)

func TestGenerateMerkelRoot(t *testing.T) {
	transactions := MockTransactions()

	hash, err := GenerateMerkelRoot(transactions)
	if err != nil {

	}

}
