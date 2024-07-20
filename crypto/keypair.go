package crypto

import (
	"Starcoin/types"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return PrivateKey{
		key: key,
	}
}

func (k PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		key: &k.key.PublicKey,
	}
}

type PublicKey struct {
	key *ecdsa.PublicKey
}

func (k PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(k.key, k.key.X, k.key.Y)
}

func (k PublicKey) Address() types.Address {
	h := sha256.Sum256(k.ToSlice())

	return types.AddressFromBytes(h[len(h)-20:])
}

type Signature struct {
}
