package server

import (
	"context"
	"net"
	"sync"

	pb "github.com/rauschp/nexis-chain/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type Node struct {
	Version  string
	peerLock sync.RWMutex
	Host     string
	Peers    map[string]*PeerNode

	pb.UnimplementedNodeServiceServer
}

type PeerNode struct {
	Version    string
	Host       string
	Connection *grpc.ClientConn
}

func NewNode(addr string, peerNodes []string) *Node {
	newNode := &Node{
		Version: "nexis7-0.0.1",
		Host:    addr,
		Peers:   make(map[string]*PeerNode),
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
			newNode.Peers[peerNode.Host] = peerNode
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

func (n *Node) HandleTransaction(ctx context.Context, t *pb.Transaction) (*pb.EmptyAckResponse, error) {
	p, _ := peer.FromContext(ctx)

	log.Debug().Msgf("Transaction received from %s" + p.Addr.Network())

	return &pb.EmptyAckResponse{}, nil
}

func (n *Node) Initialize(ctx context.Context, m *pb.InitMessage) (*pb.InitMessage, error) {
	_, ok := n.Peers[m.Address]
	if !ok {
		// Peer doesn't exist in map, add it :)
		p := &PeerNode{
			Version: m.Version,
			Host:    m.Address,
		}

		n.addPeer(p)

	}

	var hosts []string
	n.peerLock.RLock()
	defer n.peerLock.RUnlock()

	for hostnameString, _ := range n.Peers {
		hosts = append(hosts, hostnameString)
	}

	return &pb.InitMessage{
		Version:   n.Version,
		Height:    0,
		Address:   n.Host,
		NodeHosts: hosts,
	}, nil
}

func (n *Node) addPeer(node *PeerNode) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()

	client, err := grpc.Dial(node.Host, grpc.WithInsecure())
	if err != nil {
		log.Error().Err(err).Msg("Error dialing peer")
	}

	node.Connection = client

	n.Peers[node.Host] = node
}
