package storage

import (
	"github.com/rauschp/nexis-chain/crypto"
	pb "github.com/rauschp/nexis-chain/proto"
	"github.com/rauschp/nexis-chain/types"
)

type BlockStore interface {
	Height() int
	Set(block *pb.Block) error
	Get(hash string) (*pb.Block, error)
}

type WalletStore interface {
	GetByPublicKey(pc *crypto.PublicKey) (*types.Wallet, error)
	GetByAddress(address crypto.Address) (*types.Wallet, error)
	AddUnspentCurrency(address crypto.Address, transaction *pb.TransactionOutput) error
	WithdrawCurrency(address crypto.Address, amount int64) ([]*pb.TransactionOutput, error)
}
