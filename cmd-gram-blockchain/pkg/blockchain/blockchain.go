package blockchain

import (
	"cmd-gram-blockchain/internal/models"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
)

const (
	dbFile       = "dbfile.db"
	blocksBucket = "blockBucket"
)

type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

func (bc *Blockchain) AddBlock(data models.MessageDTO) error {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		return err
	}

	newBlock, err := NewBlock(data, lastHash)
	if err != nil {
		return err
	}

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		res, err := json.Marshal(newBlock)
		if err != nil {
			return err
		}

		err = b.Put(newBlock.Hash, res)
		if err != nil {
			return err
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			return err
		}

		bc.tip = newBlock.Hash

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func NewGenesisBlock() (*Block, error) {
	var msg = models.MessageDTO{Body: "Genesis Block"}
	return NewBlock(msg, []byte{})
}

func New() (*Blockchain, error) {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis, err := NewGenesisBlock()
			if err != nil {
				return err
			}
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				return err
			}

			res, err := json.Marshal(genesis)
			if err != nil {
				return err
			}

			err = b.Put(genesis.Hash, res)
			if err != nil {
				return err
			}

			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				return err
			}

			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	bc := Blockchain{tip, db}

	return &bc, nil

}

func (bc *Blockchain) PrintChain() {
	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
