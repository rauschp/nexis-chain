package server

import (
	"net"
)

type ServerNode struct {
	HostAddress string
	Listener    net.Listener
}

func New(addr string) (*ServerNode, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	node := &ServerNode{
		HostAddress: addr,
		Listener:    listener,
	}

	return node, nil
}
