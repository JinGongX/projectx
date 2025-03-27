package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestGeneratePrivateKey(t *testing.T) {
// 	PrivKey := GeneratePrivateKey()
// 	pubkey := PrivKey.PublicKey()
// 	//address := Pubkey.Address()

// 	msg := []byte("hello world")
// 	sig, err := PrivKey.Sign(msg)
// 	assert.Nil(t, err)

// 	b := sig.Verify(pubkey, msg)
// 	assert.True(t, b)

// }
func TestKeypairSignVerify(t *testing.T) {
	privKey := GeneratePrivateKey()
	publicKey := privKey.PublicKey()
	msg := []byte("hello world")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	assert.True(t, sig.Verify(publicKey, msg))
}

func TestKeypairSignVerifyFail(t *testing.T) {
	privKey := GeneratePrivateKey()
	publicKey := privKey.PublicKey()
	msg := []byte("hello world")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	otherPrivKey := GeneratePrivateKey()
	otherPublicKey := otherPrivKey.PublicKey()

	assert.False(t, sig.Verify(otherPublicKey, msg))
	assert.False(t, sig.Verify(publicKey, []byte("xxxxxx")))
}
