package core

import "bytes"

func powMask() []byte {
	buf := new(bytes.Buffer)
	return buf.Bytes()
}

func generateProofOfWork(block *Block) {
	for {
		hash, _ := block.Hash()
		if checkProofOfWork(hash) {
			break
		}
		block.Header.Nonce++
	}
}

func checkProofOfWork(hash []byte) bool {
	return true
}
