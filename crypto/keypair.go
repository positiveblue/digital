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
func GenerateNewKeypair() *Keypair {

	curve, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {

	}

	// xy := bigJoin(28, curve.PublicKey.X, curve.PublicKey.Y)

	public := common, BigIntToBase64(x)
	private := common.BigIntToBase64(curve.D)

	kp := Keypair{Public: public, Private: private}

	return &kp
}
