package core

import (
	"projectx/crypto"
	types "projectx/type"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}
	tx := Transaction{
		Data: []byte("foo"),
	}
	return NewBlock(header, []Transaction{tx})
}
func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)

	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())

}

// func TestHeader_Encode_Decode(t *testing.T) {
// 	h := &Header{
// 		Version:   1,
// 		PrevBlock: types.RandomHash(),
// 		Timestamp: uint64(time.Now().UnixNano()),
// 		Height:    10,
// 		Nonce:     1234567890,
// 	}

// 	buf := &bytes.Buffer{}
// 	assert.Nil(t, h.EncodeBinary(buf))

// 	hDecode := &Header{}
// 	assert.Nil(t, hDecode.DecodeBinary(buf))
// 	assert.Equal(t, h, hDecode)
// }

// func TestBlock_Encode_Decode(t *testing.T) {
// 	b := &Block{
// 		Header: Header{
// 			Version:   1,
// 			PrevBlock: types.RandomHash(),
// 			Timestamp: uint64(time.Now().UnixNano()),
// 			Height:    10,
// 			Nonce:     1234567890,
// 		},
// 		Transactions: nil,
// 	}
// 	buf := &bytes.Buffer{}
// 	assert.Nil(t, b.EncodeBinary(buf))

// 	bDecode := &Block{}
// 	assert.Nil(t, bDecode.DecodeBinary(buf))
// 	assert.Equal(t, b, bDecode)
// 	//fmt.Printf("%+v", bDecode)
// }

// func TestBlockHash(t *testing.T) {
// 	b := &Block{
// 		Header: Header{
// 			Version:   1,
// 			PrevBlock: types.RandomHash(),
// 			Timestamp: uint64(time.Now().UnixNano()),
// 			Height:    10,
// 		},
// 		Transactions: []Transaction{},
// 	}
// 	h := b.Hash()
// 	fmt.Println(h)
// 	assert.False(t, h.IsZero())
// }
