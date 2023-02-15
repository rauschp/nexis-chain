package types

import (
	"testing"

	"github.com/rauschp/nexis-chain/crypto"
	pb "github.com/rauschp/nexis-chain/proto"
	"github.com/stretchr/testify/assert"
)

func TestHashBlock(t *testing.T) {
	block := createBlock()
	test := HashBlock(block)

	assert.True(t, len(test) == 32)
}

func TestBlockVerify_Success(t *testing.T) {
	block := createBlock()
	pk := crypto.GenerateNewPrivateKey()
	hash := HashBlock(block)

	sig := pk.Sign(hash)

	assert.True(t, pk.Public().Verify(hash, sig))
}

func TestBlockVerify_InvalidSigniature(t *testing.T) {
	block := createBlock()
	pk := crypto.GenerateNewPrivateKey()
	hash := HashBlock(block)

	sig := pk.Sign(hash)

	diffPk := crypto.GenerateNewPrivateKey()

	assert.False(t, diffPk.Public().Verify(hash, sig))
}

func createBlock() *pb.Block {
	block := &pb.Block{
		Header: &pb.Header{
			Version: 1,
			Height:  0,
		},
		Transactions: []*pb.Transaction{},
	}

	return block
}
