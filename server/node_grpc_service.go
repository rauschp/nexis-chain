package server

import (
	"context"

	pb "github.com/rauschp/nexis-chain/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/peer"
)

func (n *Node) HandleTransaction(ctx context.Context, t *pb.Transaction) (*pb.EmptyAckResponse, error) {
	p, _ := peer.FromContext(ctx)

	if n.Mempool.AddTransaction(t) {
		log.Debug().Msgf("Transaction added to mempool on %s from %s", n.Host, p.Addr)
		// Broadcasting to peer nodes
		go func() {
			if err := n.BroadcastEvent(t); err != nil {
				log.Error().Err(err).Msg("Error broadcasting event")
			}
		}()
	}

	return &pb.EmptyAckResponse{}, nil
}

func (n *Node) Initialize(ctx context.Context, m *pb.InitMessage) (*pb.InitMessage, error) {
	hasValue := n.PeerManager.ContainsPeerByString(m.Address)
	if !hasValue {
		// Peer doesn't exist in map, add it :)
		p := &PeerNode{
			Version: m.Version,
			Host:    m.Address,
		}

		n.addPeer(p)

		log.Debug().Msgf("Adding peer (%s) to %s", p.Host, n.Host)
	}

	n.PeerManager.Lock.Lock()
	defer n.PeerManager.Lock.Unlock()

	var hosts []string
	for hostnameString, _ := range n.PeerManager.Peers {
		hosts = append(hosts, hostnameString)
	}

	return &pb.InitMessage{
		Version:   n.Version,
		Height:    0,
		Address:   n.Host,
		NodeHosts: hosts,
	}, nil
}
