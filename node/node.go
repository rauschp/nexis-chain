package node

import (
	"net"
)

type ServerNode struct {
	HostAddress string
	Listener    net.Listener
}
