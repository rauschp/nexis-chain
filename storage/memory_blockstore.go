package storage

import (
	"encoding/hex"
	"errors"
	pb "github.com/rauschp/nexis-chain/proto"
	"github.com/rauschp/nexis-chain/types"
	"sync"
)

type MemoryBlockstore struct {
	Blocks map[string]*pb.Block
	Lock   sync.RWMutex
}

func CreateMemoryBlockstore() *MemoryBlockstore {
	bs := &MemoryBlockstore{
		Blocks: make(map[string]*pb.Block),
	}

	return bs
}

func (m *MemoryBlockstore) Height() int {
	m.Lock.RLock()
	defer m.Lock.RUnlock()

	return len(m.Blocks)
}

func (m *MemoryBlockstore) Get(hash string) (*pb.Block, error) {
	m.Lock.RLock()
	defer m.Lock.RUnlock()

	block, exists := m.Blocks[hash]

	if !exists {
		return nil, errors.New("block not found by hash")
	}

	return block, nil
}

func (m *MemoryBlockstore) Set(block *pb.Block) error {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	hash := types.HashBlock(block)
	m.Blocks[hex.EncodeToString(hash)] = block

	return nil
}
