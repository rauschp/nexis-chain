package server

import (
	"context"
	"github.com/rauschp/nexis-chain/crypto"
	"google.golang.org/grpc/credentials/insecure"
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

type PeerNode struct {
	Version    string
	Host       string
	Connection *grpc.ClientConn
}

type NodeConfig struct {
	Host       string
	PrivateKey *crypto.PrivateKey
}

type Node struct {
	Version     string
	PeerManager *PeerManager
	Mempool     *Mempool
	NodeConfig  NodeConfig

	pb.UnimplementedNodeServiceServer
}

func NewNode(addr string, peerNodes []string, privateKey *crypto.PrivateKey) *Node {
	newNode := &Node{
		Version: "nexis7-0.0.1",
		NodeConfig: NodeConfig{
			Host:       addr,
			PrivateKey: privateKey,
		},
		PeerManager: CreatePeerManager(),
		Mempool:     CreateMempool(),
	}

	if len(peerNodes) > 0 {
		for _, nodeHost := range peerNodes {
			if nodeHost == addr {
				continue
			}

			c, err := grpc.Dial(nodeHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	log.Debug().Msgf("Starting server on %s", n.NodeConfig.Host)

	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", n.NodeConfig.Host)

	log.Debug().Msgf("Listening on %s", n.NodeConfig.Host)

	if err != nil {
		log.Error().Stack().Err(err).Msg("Unable to create server")
	}

	pb.RegisterNodeServiceServer(grpcServer, n)
	log.Info().Msgf("Node started on host %s", n.NodeConfig.Host)

	grpcServer.Serve(lis)
}

func (n *Node) addPeer(node *PeerNode) {
	client, err := grpc.Dial(node.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Err(err).Msg("Error dialing peer")
	}

	node.Connection = client

	n.PeerManager.AddPeer(node)
}
