package main

import (
	"encoding/hex"
	"fmt"
	"net"

	"github.com/rauschp/nexis-chain/crypto"
	"github.com/rauschp/nexis-chain/node"
)

func main() {
	privateKey := crypto.GenerateNewPrivateKey()

	fmt.Println(hex.EncodeToString(privateKey.Key))
	fmt.Println(hex.EncodeToString(privateKey.Public().Key))
}

func createServerNode(addr string) (*node.ServerNode, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	node := &node.ServerNode{
		HostAddress: addr,
		Listener:    listener,
	}

	return node, nil
}
