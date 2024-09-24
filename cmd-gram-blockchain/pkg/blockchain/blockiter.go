package blockchain

import (
	"encoding/json"

	"github.com/boltdb/bolt"
)

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		err := json.Unmarshal(encodedBlock, &block)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil
	}

	i.currentHash = block.PrevBlockHash

	return block
}
