package main

import (
	"context"
	"fmt"
	"github.com/rauschp/nexis-chain/crypto"
	pb "github.com/rauschp/nexis-chain/proto"
	"github.com/rauschp/nexis-chain/server"
	"github.com/rauschp/nexis-chain/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
)

type App struct {
	Node       *server.Node
	Blockchain *types.Blockchain
	PrivateKey *crypto.PrivateKey
}

func (a *App) createServer(addr string, nodes []string) *server.Node {
	privKey := crypto.GenerateNewPrivateKey()
	n := server.NewNode(addr, nodes, privKey)

	go n.StartNodeServer()

	return n
}

func (a *App) sendTransaction(node *server.Node, destAddr string) {
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
