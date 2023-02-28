package types

import "github.com/rauschp/nexis-chain/storage"

type Blockchain struct {
	BlockStore  storage.BlockStore
	WalletStore storage.WalletStore
}
