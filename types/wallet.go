package types

import "github.com/rauschp/nexis-chain/crypto"

type Wallet struct {
	PublicKey *crypto.PublicKey
	Address   crypto.Address
	Balance   int64
}
