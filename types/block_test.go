package types

import (
	"testing"

	pb "github.com/rauschp/nexis-chain/proto"
	"github.com/stretchr/testify/assert"
)

func TestHashBlock(t *testing.T) {
	block := &pb.Block{
		Header: &pb.Header{
			Version: 1,
			Height:  0,
		},
		Transactions: []*pb.Transaction{},
	}

	test := HashBlock(block)

	assert.True(t, len(test) == 32)
}
