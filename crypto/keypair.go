package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

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

func (kp *Keypair) Sign(hash []byte) ([]byte, error) {
	curve := kp.curve()
	R, S, err := ecdsa.Sign(rand.Reader, curve, hash)
	if err != nil {
		return nil, err
	}

	signature := append(R.Bytes(), S.Bytes()...)

	return signature, nil
}

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
	var x, y, D *big.Int
	x.SetBytes(kp.Public[:32])
	y.SetBytes(kp.Public[32:])
	D.SetBytes(kp.Private)

	return &ecdsa.PrivateKey{ecdsa.PublicKey{elliptic.P256(), x, y}, D}
}

func splitPoint(bytes []byte) (x, y *big.Int) {
	x.SetBytes(bytes[:32])
	y.SetBytes(bytes[32:])

	return x, y
}

func joinPoints(x, y *big.Int) []byte {
	xBytes := x.Bytes()
	yBytes := y.Bytes()

	return append(xBytes, yBytes...)
}
