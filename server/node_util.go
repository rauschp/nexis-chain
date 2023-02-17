package server

import (
	"encoding/hex"

	pb "github.com/rauschp/nexis-chain/proto"
	"github.com/rauschp/nexis-chain/types"
)

// Mempool Functions
func CreateMempool() *Mempool {
	return &Mempool{
		Pool: make(map[string]*pb.Transaction),
	}
}

func (m *Mempool) ContainsTransaction(t *pb.Transaction) bool {
	m.Lock.RLock()
	defer m.Lock.RUnlock()

	hash := hex.EncodeToString(types.HashTransaction(t))
	_, hasTransaction := m.Pool[hash]

	return hasTransaction
}

func (m *Mempool) AddTransaction(t *pb.Transaction) bool {
	if m.ContainsTransaction(t) {
		return false
	}

	m.Lock.Lock()
	defer m.Lock.Unlock()

	hash := hex.EncodeToString(types.HashTransaction(t))
	m.Pool[hash] = t

	return true
}

// Peer Manager Functions
func CreatePeerManager() *PeerManager {
	return &PeerManager{
		Peers: make(map[string]*PeerNode),
	}
}

func (pm *PeerManager) ContainsPeer(p *PeerNode) bool {
	pm.Lock.RLock()
	defer pm.Lock.RUnlock()

	_, hasPeer := pm.Peers[p.Host]

	return hasPeer
}

func (pm *PeerManager) ContainsPeerByString(host string) bool {
	pm.Lock.RLock()
	defer pm.Lock.RUnlock()

	_, hasPeer := pm.Peers[host]

	return hasPeer
}

func (pm *PeerManager) AddPeer(p *PeerNode) {
	if pm.ContainsPeer(p) {
		return
	}

	pm.Lock.Lock()
	defer pm.Lock.Unlock()

	pm.Peers[p.Host] = p
}
