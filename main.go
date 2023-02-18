package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
	"time"

	pb "github.com/rauschp/nexis-chain/proto"
	"github.com/rauschp/nexis-chain/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	createServer(":3050", []string{})
	time.Sleep(2 * time.Second)
	createServer(":3060", []string{":3050"})
	time.Sleep(2 * time.Second)
	n3 := createServer(":3070", []string{":3060"})
	time.Sleep(3 * time.Second)

	sendTransaction(n3, ":3050")

	// Stop program from exiting
	select {}
}

func createServer(addr string, nodes []string) *server.Node {
	n := server.NewNode(addr, nodes)

	go n.StartNodeServer()

	return n
}

func sendTransaction(node *server.Node, destAddr string) {
	c, err := grpc.Dial(destAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Err(err).Msg("Error dialing")
	}

	grpcClient := pb.NewNodeServiceClient(c)

	transInput := &pb.TransactionInput{
		Address:   make([]byte, 32),
		Amount:    10,
		PublicKey: make([]byte, 32),
		Signature: make([]byte, 32),
	}

	transOutput := &pb.TransactionOutput{
		Address: make([]byte, 32),
		Amount:  10,
	}

	_, err = grpcClient.HandleTransaction(context.Background(), &pb.Transaction{
		Version: fmt.Sprintf("test-%d", rand.Intn(500000)),
		Inputs:  []*pb.TransactionInput{transInput},
		Outputs: []*pb.TransactionOutput{transOutput},
	})
	if err != nil {
		log.Error().Err(err).Msg("Unable to send message")
	}

}
