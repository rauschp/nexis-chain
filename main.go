package main

import (
	"encoding/hex"
	"fmt"

	"github.com/rauschp/nexis-chain/crypto"
)

func main() {
	privateKey := crypto.GenerateNewPrivateKey()

	fmt.Println(hex.EncodeToString(privateKey.Key))
	fmt.Println(hex.EncodeToString(privateKey.Public().Key))
}
