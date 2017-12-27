package core

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/jomsdev/digital/common"
)

type Block struct {
	Header       BlockHeader
	Signature    []byte
	Transactions []Transaction
}

type BlockHeader struct {
	From       []byte
	PrevHash   []byte
	MarkleRoot []byte
	Timestamp  uint32
	Nonce      uint32
}

func GenerateBlock(from, prevHash []byte, transactions []Transaction) Block {
	block := newBlock(from, prevHash, transactions)
	block.Header.MarkleRoot, _ = GenerateMerkleRoot(block.Transactions)
	generateProofOfWork(&block)
	block.Header.Timestamp = uint32(time.Now().Unix())
	return block
}

func (block *Block) Hash() ([]byte, error) {

	bh := block.Header
	buf := new(bytes.Buffer)

	fromKey, err := common.Pkcs7Pad(bh.From, NetworkKeySize)
	if err != nil {
		return nil, err
	}

	buf.Write(fromKey)
	buf.Write(bh.PrevHash)
	buf.Write(bh.MarkleRoot)
	binary.Write(buf, binary.LittleEndian, bh.Nonce)

	return common.Sha256(buf.Bytes()), nil
}

func newBlock(from, prevHash []byte, transactions []Transaction) Block {
	header := BlockHeader{From: from, PrevHash: prevHash, Nonce: 0}
	return Block{Header: header, Transactions: transactions}
}
