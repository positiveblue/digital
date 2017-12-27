package core

import (
	"errors"

	"github.com/jomsdev/digital/common"
)

func GenerateMerkelRoot(transactions []Transaction) ([]byte, error) {
	if len(transactions) == 0 {
		return nil, errors.New("merkleTree: Hashing and empty slice of transactions")
	} else if len(transactions) == 1 {
		hash, err := transactions[0].Hash()
		if err != nil {
			return nil, err
		}
		return hash, nil
	} else {
		mid := len(transactions) / 2
		leftHash, _ := GenerateMerkelRoot(transactions[:mid])
		rightHash, _ := GenerateMerkelRoot(transactions[mid:])
		nodeHash := common.Sha256(append(leftHash, rightHash...))
		return nodeHash, nil
	}
}
