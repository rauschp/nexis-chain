package server

import (
	"context"
	"net"
	"sync"

	pb "github.com/rauschp/nexis-chain/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type PeerManager struct {
	Peers map[string]*PeerNode
	Lock  sync.RWMutex
}

type Mempool struct {
	Pool map[string]*pb.Transaction
	Lock sync.RWMutex
}

type Node struct {
	Version     string
	PeerManager *PeerManager
	Mempool     *Mempool
	Host        string

	pb.UnimplementedNodeServiceServer
}

type PeerNode struct {
	Version    string
	Host       string
	Connection *grpc.ClientConn
}

func NewNode(addr string, peerNodes []string) *Node {
	newNode := &Node{
		Version:     "nexis7-0.0.1",
		Host:        addr,
		PeerManager: CreatePeerManager(),
		Mempool:     CreateMempool(),
	}

	if len(peerNodes) > 0 {
		for _, nodeHost := range peerNodes {
			if nodeHost == addr {
				continue
			}

			c, err := grpc.Dial(nodeHost, grpc.WithInsecure())
			if err != nil {
				log.Logger.Error().Err(err).Msg("Error dialing")
			}

			nodeClient := pb.NewNodeServiceClient(c)
			res, err := nodeClient.Initialize(context.Background(), &pb.InitMessage{
				Version: newNode.Version,
				Height:  0,
				Address: addr,
				Success: true,
			})
			if err != nil {
				log.Error().Err(err).Msg("Unable to initalize node")
			}

			peerNode := &PeerNode{
				Version:    res.Version,
				Host:       res.Address,
				Connection: c,
			}

			log.Info().Msgf("Adding client (%s) to node %s", nodeHost, addr)
			newNode.PeerManager.AddPeer(peerNode)
		}
	}

	return newNode
}

func (n *Node) StartNodeServer() {
	log.Debug().Msgf("Starting server on %s", n.Host)

	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", n.Host)

	log.Debug().Msgf("Listening on %s", n.Host)

	if err != nil {
		log.Error().Stack().Err(err).Msg("Unable to create server")
	}

	pb.RegisterNodeServiceServer(grpcServer, n)
	log.Info().Msgf("Node started on host %s", n.Host)

	grpcServer.Serve(lis)
}

func (n *Node) addPeer(node *PeerNode) {
	client, err := grpc.Dial(node.Host, grpc.WithInsecure())
	if err != nil {
		log.Error().Err(err).Msg("Error dialing peer")
	}

	node.Connection = client

	n.PeerManager.AddPeer(node)
}
