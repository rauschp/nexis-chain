package storage

import (
	"encoding/json"
	"github.com/cockroachdb/pebble"
	pb "github.com/rauschp/nexis-chain/proto"
	"github.com/rs/zerolog/log"
)

type PersistentBlockstore struct {
	DB *pebble.DB
}

func CreatePersistentBlockstore() *PersistentBlockstore {
	db, err := pebble.Open("data/block-store", &pebble.Options{})
	if err != nil {
		log.Error().Err(err).Msg("Error opening store")
	}

	bs := &PersistentBlockstore{
		DB: db,
	}

	return bs
}

func (m *PersistentBlockstore) Height() int {
	value, closer, err := m.DB.Get([]byte("Height"))
	if err != nil {
		log.Error().Err(err).Msg("Error getting height value")
	}

	var height = int(value[0])
	if err := closer.Close(); err != nil {
		log.Error().Err(err).Msg("Error closing reader")
		return 0
	}

	return height
}

func (m *PersistentBlockstore) Get(hash string) (*pb.Block, error) {
	value, closer, err := m.DB.Get([]byte(hash))
	if err != nil {
		log.Error().Err(err).Msg("Error getting block")
		return nil, err
	}

	var block = &pb.Block{}
	if err := json.Unmarshal(value, block); err != nil {
		log.Error().Stack().Err(err).Msgf("Unable to marshal persistent store for key %s", hash)
		return nil, err
	}

	if err := closer.Close(); err != nil {
		log.Error().Err(err).Msg("Error closing reader")
		return nil, err
	}

	return block, nil
}

func (m *PersistentBlockstore) Set(block *pb.Block) error {
	key := []byte(block.Header.Hash)
	value, err := json.Marshal(block)
	if err != nil {
		return err
	}

	if err := m.DB.Set(key, value, pebble.Sync); err != nil {
		log.Error().Err(err).Msg("Error persisting key")
		return err
	}

	return nil
}
