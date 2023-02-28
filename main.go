package main

import (
	"github.com/rauschp/nexis-chain/storage"
	"github.com/rs/zerolog"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	app := &App{
		Blockchain: &Blockchain{
			BlockStore:  storage.CreatePersistentBlockstore(),
			WalletStore: storage.CreatePersistentWalletstore(),
		},
	}

	app.createServer(":3050", []string{})
	time.Sleep(2 * time.Second)

	app.createServer(":3060", []string{":3050"})
	time.Sleep(2 * time.Second)

	n3 := app.createServer(":3070", []string{":3060"})
	time.Sleep(3 * time.Second)

	app.sendTransaction(n3, ":3050")
	app.sendTransaction(n3, ":3070")

	// Stop program from exiting
	select {}
}
