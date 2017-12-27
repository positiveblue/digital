package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"math/big"
)

// Key generation with proof of work
type Keypair struct {
	Public  []byte
	Private []byte
}

// GenerateNewKeypair retunrs a key pair for the P256 Curve (which implments the P-256, see FIPS 186-3, section D.2.3)
// Cryptographic operations for P256 are implemented using constant-time algorithms.
func GenerateNewKeypair() (*Keypair, error) {
	curve, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	kp := Keypair{nil, nil}
	kp.init(curve)

	return &kp, nil
}

func (kp *Keypair) FromBase64(public, private string) {
	kp.Public, _ = base64.StdEncoding.DecodeString(public)
	kp.Private, _ = base64.StdEncoding.DecodeString(private)
}

func (kp *Keypair) ToBase64() (public, private string) {
	public = base64.StdEncoding.EncodeToString(kp.Public)
	private = base64.StdEncoding.EncodeToString(kp.Private)
	return public, private
}

//TODO: Document function
func (kp *Keypair) Sign(hash []byte) ([]byte, error) {
	curve := kp.curve()
	R, S, err := ecdsa.Sign(rand.Reader, curve, hash)
	if err != nil {
		return nil, err
	}

	signature := append(R.Bytes(), S.Bytes()...)

	return signature, nil
}

//TODO: Document function
func Verify(pk, sig, hash []byte) bool {
	x, y := splitPoint(pk)
	r, s := splitPoint(sig)
	publicKey := ecdsa.PublicKey{elliptic.P256(), x, y}
	return ecdsa.Verify(&publicKey, hash, r, s)
}

func (kp *Keypair) init(curve *ecdsa.PrivateKey) {
	kp.Public = joinPoints(curve.PublicKey.X, curve.PublicKey.Y)
	kp.Private = curve.D.Bytes()
}

func (kp *Keypair) curve() (curve *ecdsa.PrivateKey) {
	x := new(big.Int).SetBytes(kp.Public[:32])
	y := new(big.Int).SetBytes(kp.Public[32:])
	D := new(big.Int).SetBytes(kp.Private)

	return &ecdsa.PrivateKey{ecdsa.PublicKey{elliptic.P256(), x, y}, D}
}

func splitPoint(bytes []byte) (x, y *big.Int) {
	x = new(big.Int).SetBytes(bytes[:32])
	y = new(big.Int).SetBytes(bytes[32:])

	return x, y
}

func joinPoints(x, y *big.Int) []byte {
	xBytes := x.Bytes()
	yBytes := y.Bytes()

	return append(xBytes, yBytes...)
}
