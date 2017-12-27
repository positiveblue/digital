package core

import (
	"testing"
)

func TestGenerateMerkleRoot(t *testing.T) {
	transactions := MockTransactions()

	hash, err := GenerateMerkleRoot(transactions)
	if err != nil {

	}

}
