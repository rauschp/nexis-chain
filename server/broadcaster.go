package server

import (
	"context"

	pb "github.com/rauschp/nexis-chain/proto"
	"google.golang.org/grpc"
)

func (n *Node) BroadcastEvent(msg any) error {
	n.PeerManager.Lock.RLock()
	defer n.PeerManager.Lock.RUnlock()

	for host, _ := range n.PeerManager.Peers {
		switch msgType := msg.(type) {
		case *pb.Transaction:
			c, err := grpc.Dial(host, grpc.WithInsecure())
			if err != nil {
				return err
			}

			client := pb.NewNodeServiceClient(c)

			_, err = client.HandleTransaction(context.Background(), msgType)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
