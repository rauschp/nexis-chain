package main

import (
	"github.com/rauschp/nexis-chain/server"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	createServer(":3050", []string{})
	createServer(":3060", []string{":3050"})
	createServer(":3070", []string{":3060"})

	select {}
}

func createServer(addr string, nodes []string) {
	n := server.NewNode(addr, nodes)

	go n.StartNodeServer()
}
