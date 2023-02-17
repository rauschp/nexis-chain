package server

import (
	"context"

	pb "github.com/rauschp/nexis-chain/proto"
)

func (n *Node) BroadcastEvent(msg any) error {
	n.PeerManager.Lock.RLock()
	defer n.PeerManager.Lock.RUnlock()

	for _, peer := range n.PeerManager.Peers {
		switch msgType := msg.(type) {
		case *pb.Transaction:
			c := peer.Connection

			client := pb.NewNodeServiceClient(c)

			_, err := client.HandleTransaction(context.Background(), msgType)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
