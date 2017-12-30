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
	MerkleRoot []byte
	Timestamp  uint32
	Nonce      uint32
}

func NewBlock(from, prevHash []byte, transactions []Transaction) Block {
	block := newBlock(from, prevHash, transactions)
	block.Header.MerkleRoot, _ = GenerateMerkleRoot(block.Transactions)
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
	buf.Write(bh.MerkleRoot)
	binary.Write(buf, binary.LittleEndian, bh.Nonce)

	return common.Sha256(buf.Bytes()), nil
}

func (block *Block) Verify(prevHash []byte) bool {
	// Verify signature
	pk := block.Header.From
	sig := block.Signature
	hash, _ := block.Hash()

	if !Verify(pk, sig, hash) {
		return false
	}

	// Verify Markle root
	MerkleRoot, _ := GenerateMerkleRoot(block.Transactions)
	if !bytes.Equal(block.Header.MerkleRoot, MerkleRoot) {
		return false
	}

	// Verify time
	if block.Header.Timestamp >= uint32(time.Now().Unix()) {
		return false
	}

	// Verify ProofOfWork
	if !checkProofOfWork(hash) {
		return false
	}

	return true
}

func newBlock(from, prevHash []byte, transactions []Transaction) Block {
	header := BlockHeader{From: from, PrevHash: prevHash, Nonce: 0}
	return Block{Header: header, Transactions: transactions}
}
