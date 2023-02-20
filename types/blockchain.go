package types

import pb "github.com/rauschp/nexis-chain/proto"

type Blockchain struct {
	BlockStore  *BlockStore
	WalletStore *WalletStore
}

type BlockStore interface {
	Height() int
	Set(block *pb.Block) error
	Get(hash string) (*pb.Block, error)
}

type WalletStore interface {
	GetByPublicKey(hash string)
	GetByAddress(address string)
	Set()
}
