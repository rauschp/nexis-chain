package storage

import (
	"encoding/json"
	"errors"
	"github.com/cockroachdb/pebble"
	"github.com/rauschp/nexis-chain/crypto"
	pb "github.com/rauschp/nexis-chain/proto"
	"github.com/rauschp/nexis-chain/types"
	"github.com/rs/zerolog/log"
)

type PersistentWalletStore struct {
	DB *pebble.DB
}

type WalletRecord struct {
	Transactions []*pb.TransactionOutput
	PublicKey    *crypto.PublicKey
}

func CreatePersistentWalletstore() *PersistentWalletStore {
	db, err := pebble.Open("data/wallet-store", &pebble.Options{})
	if err != nil {
		log.Error().Err(err).Msg("Error opening store")
	}

	return &PersistentWalletStore{
		DB: db,
	}
}

func (w *PersistentWalletStore) GetByPublicKey(pk *crypto.PublicKey) (*types.Wallet, error) {
	addr := pk.GetAddress()

	return w.GetByAddress(addr)
}

func (w *PersistentWalletStore) GetByAddress(address crypto.Address) (*types.Wallet, error) {
	record, err := w.getWalletRecord(address)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching wallet")
		return nil, err
	}

	bal := calculateBalance(record.Transactions)

	return &types.Wallet{
		Address:   address,
		PublicKey: record.PublicKey,
		Balance:   bal,
	}, nil
}

func (w *PersistentWalletStore) AddUnspentCurrency(address crypto.Address, transaction *pb.TransactionOutput) error {
	record, err := w.getWalletRecord(address)
	if err != nil {
		log.Error().Err(err).Msg("Error getting wallet")
		return err
	}

	newTransactions := append(record.Transactions, transaction)
	newRecord := WalletRecord{
		Transactions: newTransactions,
		PublicKey:    record.PublicKey,
	}

	if err := w.saveWalletRecord(address, newRecord); err != nil {
		return err
	}

	return nil
}

func (w *PersistentWalletStore) WithdrawCurrency(address crypto.Address, amount int64) ([]*pb.TransactionOutput, error) {
	record, err := w.getWalletRecord(address)
	if err != nil {
		log.Error().Err(err).Msg("Error getting wallet")
		return nil, err
	}

	bal := calculateBalance(record.Transactions)

	if bal < amount {
		err = errors.New("insufficent funds")
		log.Error().Err(err)
		return nil, err
	}

	// Mark all as spent
	var newTransList []*pb.TransactionOutput

	for _, trans := range record.Transactions {
		t := &pb.TransactionOutput{
			Address: trans.Address,
			Amount:  trans.Amount,
			Spent:   true,
		}
		newTransList = append(newTransList, t)
	}

	newTransaction := &pb.TransactionOutput{
		Address: address.ToBytes(),
		Amount:  bal - amount,
		Spent:   false,
	}

	newTransList = append(newTransList, newTransaction)

	newRecord := WalletRecord{
		Transactions: newTransList,
		PublicKey:    record.PublicKey,
	}

	if err := w.saveWalletRecord(address, newRecord); err != nil {
		log.Error().Err(err).Msg("error saving wallet")
		return nil, err
	}

	return newTransList, nil
}

func (w *PersistentWalletStore) getWalletRecord(address crypto.Address) (*WalletRecord, error) {
	value, closer, err := w.DB.Get(address.ToBytes())
	if err != nil {
		log.Error().Err(err).Msg("Error getting height value")
		return nil, err
	}

	if err := closer.Close(); err != nil {
		log.Error().Err(err).Msg("Error closing reader")
		return nil, err
	}

	var record = &WalletRecord{}
	if err = json.Unmarshal(value, record); err != nil {
		log.Error().Err(err).Msg("Error parsing kvs result")
		return nil, err
	}

	return record, nil
}

func (w *PersistentWalletStore) saveWalletRecord(address crypto.Address, record WalletRecord) error {
	key := address.ToBytes()
	value, err := json.Marshal(record)
	if err != nil {
		log.Error().Err(err).Msg("Error marshalling wallet record to save")
		return err
	}

	if err := w.DB.Set(key, value, pebble.Sync); err != nil {
		log.Error().Err(err).Msg("Error saving record to kvs")
		return err
	}

	return nil
}

func calculateBalance(transactions []*pb.TransactionOutput) int64 {
	var bal int64 = 0
	for _, trans := range transactions {
		bal += trans.Amount
	}

	return bal
}
