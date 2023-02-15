package types

import (
	"crypto/sha256"

	protobuf "github.com/golang/protobuf/proto"
	pb "github.com/rauschp/nexis-chain/proto"
)

func HashBlock(block *pb.Block) []byte {
	bl, err := protobuf.Marshal(block)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(bl)

	return hash[:]
}
