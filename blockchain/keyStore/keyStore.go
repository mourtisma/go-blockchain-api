package keyStore

import (
	"crypto/dsa"
	"crypto/rand"
	"errors"
)

type IkeyStore interface {
	GenerateKeys()
}

type KeyStore struct {
	PublicKey  *dsa.PublicKey
	PrivateKey *dsa.PrivateKey
}

func (ks *KeyStore) GenerateKeys() error {
	params := new(dsa.Parameters)
	if err := dsa.GenerateParameters(params, rand.Reader, dsa.L1024N160); err != nil {
		return errors.New("Error generating the parameters")
	}
	privateKey := new(dsa.PrivateKey)
	privateKey.PublicKey.Parameters = *params
	dsa.GenerateKey(privateKey, rand.Reader)
	ks.PrivateKey = privateKey
	ks.PublicKey = &privateKey.PublicKey

	return nil
}
