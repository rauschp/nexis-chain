package server

import (
	"fmt"
	"math/rand"
	"testing"

	pb "github.com/rauschp/nexis-chain/proto"
	"github.com/stretchr/testify/assert"
)

func TestCreateMempool(t *testing.T) {
	test := CreateMempool()

	assert.NotNil(t, test)
	assert.NotNil(t, test.Pool)
}

func TestMempoolAddTransaction(t *testing.T) {
	var (
		test        = CreateMempool()
		transaction = createDummyTransaction()
	)

	test.AddTransaction(transaction)

	assert.True(t, test.ContainsTransaction(transaction))
}

func TestMempoolContainsTransaction(t *testing.T) {
	var (
		test                  = CreateMempool()
		transaction           = createDummyTransaction()
		notPresentTransaction = createDummyTransaction()
	)

	test.AddTransaction(transaction)

	assert.True(t, test.ContainsTransaction(transaction))
	assert.False(t, test.ContainsTransaction(notPresentTransaction))

	// It shouldn't allow you to add a transaction that's already added
	// This ensures we don't dupe transactions which would be silly
	assert.False(t, test.AddTransaction(transaction))
}

func TestCreatePeerManager(t *testing.T) {
	test := CreatePeerManager()

	assert.NotNil(t, test)
	assert.NotNil(t, test.Peers)
}

func TestAddPeer(t *testing.T) {
	pm := CreatePeerManager()
	peer := createDummyPeerNode()

	pm.AddPeer(peer)

	assert.True(t, pm.ContainsPeer(peer))
}

func TestContainsPeer(t *testing.T) {
	pm := CreatePeerManager()
	peer := createDummyPeerNode()
	notPresentPeer := createDummyPeerNode()

	pm.AddPeer(peer)

	assert.True(t, pm.ContainsPeer(peer))
	assert.False(t, pm.ContainsPeer(notPresentPeer))

	assert.True(t, pm.ContainsPeerByString(peer.Host))
	assert.False(t, pm.ContainsPeerByString(notPresentPeer.Host))

	assert.False(t, pm.AddPeer(peer))
}

func createDummyPeerNode() *PeerNode {
	return &PeerNode{
		Version:    "test",
		Host:       fmt.Sprintf("test-%d", rand.Intn(500000)),
		Connection: nil, // Not necessary for these tests
	}
}

func createDummyTransaction() *pb.Transaction {
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

	return &pb.Transaction{
		Version: fmt.Sprintf("test-%d", rand.Intn(500000)),
		Inputs:  []*pb.TransactionInput{transInput},
		Outputs: []*pb.TransactionOutput{transOutput},
	}
}
