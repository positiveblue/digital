package core

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/jomsdev/digital/common"
)

type Transaction struct {
	Header    TransactionHeader
	Signature []byte
	Payload   []byte
}

type TransactionType uint32

const (
	Payment TransactionType = 0
)

type TransactionHeader struct {
	From        []byte
	To          []byte
	Type        TransactionType
	Timestamp   uint32
	PayloadHash []byte
}

func NewTransaction(from, to []byte, transactionType TransactionType, payload []byte) *Transaction {
	header := TransactionHeader{
		From:        from,
		To:          to,
		Type:        transactionType,
		Timestamp:   uint32(time.Now().Unix()),
		PayloadHash: common.Sha256(payload)}

	return &Transaction{Header: header, Payload: payload}
}

func (t *Transaction) Hash() ([]byte, error) {
	th := t.Header
	buf := new(bytes.Buffer)

	fromKey, err := common.Pkcs7Pad(th.From, NetworkKeySize)
	if err != nil {
		return nil, err
	}
	toKey, err := common.Pkcs7Pad(th.To, NetworkKeySize)
	if err != nil {
		return nil, err
	}

	buf.Write(fromKey)
	buf.Write(toKey)
	binary.Write(buf, binary.LittleEndian, th.Type)
	binary.Write(buf, binary.LittleEndian, th.Timestamp)
	buf.Write(th.PayloadHash)

	return common.Sha256(buf.Bytes()), nil
}

func (t *Transaction) SignTransaction(kp *Keypair) error {
	signature, err := t.sign(kp)
	if err != nil {
		return err
	}

	t.Signature = signature

	return nil
}

func (t *Transaction) VerifyTransaction() (bool, error) {
	hash, err := t.Hash()
	if err != nil {
		return false, err
	}
	return Verify(t.Header.From, t.Signature, hash), nil
}

func (t *Transaction) sign(kp *Keypair) ([]byte, error) {
	hash, err := t.Hash()
	if err != nil {
		return nil, err
	}
	s, err := kp.Sign(hash)
	if err != nil {
		return nil, err
	}

	return s, nil
}
