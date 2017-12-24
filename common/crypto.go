package common

import (
	"crypto/ecdsa"
	"encoding/base64"
	"math/big"
)

func BigIntToBase64(i *big.Int) string {
	bytes := i.Bytes()
	return base64.StdEncoding.EncodeToString(bytes)
}

func Base64ToBigInt(s string) (*big.Int, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	i := new(big.Int)
	i.SetBytes(data)
	return i, nil
}

func EncodeCurve(curve *ecdsa.PrivateKey) (pk, sk []byte) {

	xBytes := curve.PublicKey.X.Bytes()
	yBytes := curve.PublicKey.Y.Bytes()
	publicKeyBytes := append(xBytes, yBytes...)

	pk = publicKeyBytes
	sk = curve.D.Bytes()

	return pk, sk
}
