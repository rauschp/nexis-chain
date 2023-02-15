package server

import (
	"context"
	"net"

	pb "github.com/rauschp/nexis-chain/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type Node struct {
	Version string
	Peers   map[string]*grpc.ClientConn

	pb.UnimplementedNodeServiceServer
}

type PeerNode struct {
	Version    string
	Hostname   string
	Connection *grpc.ClientConn
}

func NewNode() *Node {
	return &Node{
		Version: "nexis-0.0.1",
		Peers:   make(map[string]*grpc.ClientConn),
	}
}

func (n *Node) StartNodeServer(addr string) {
	log.Debug().Msgf("Starting server on %s", addr)

	//grpcOpt := []grpc.ServerOption{grpc.WithInsecure}
	grpcServer := grpc.NewServer()
	ln, err := net.Listen("tcp", addr)
	log.Debug().Msgf("Listening on %s", addr)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Unable to create server")
	}

	log.Debug().Msgf("Registering on %s", addr)
	pb.RegisterNodeServiceServer(grpcServer, n)
	grpcServer.Serve(ln)

	log.Info().Msgf("Node started on host %s", addr)
}

func (n *Node) HandleTransaction(ctx context.Context, t *pb.Transaction) (*pb.EmptyAckResponse, error) {
	p, _ := peer.FromContext(ctx)
	log.Debug().Msgf("Transaction received from %s" + p.Addr.Network())

	return &pb.EmptyAckResponse{}, nil
}

// func (n *Node) Initialize(ctx context.Context, m *pb.InitMessage) (*pb.InitMessage, error) {

// }
