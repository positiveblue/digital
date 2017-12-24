package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/jomsdev/digital/common"
)

// Key generation with proof of work
type Keypair struct {
	Public  []byte
	Private []byte
}

// GenerateNewKeypair retunrs a key pair for the P256 Curve (which implments the P-256, see FIPS 186-3, section D.2.3)
// Cryptographic operations for P256 are implemented using constant-time algorithms.
func GenerateNewKeypair() (kp *Keypair, err error) {

	curve, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	pk, sk := common.EncodeCurve(curve)

	kp = &Keypair{Public: pk, Private: sk}

	return kp, nil
}
