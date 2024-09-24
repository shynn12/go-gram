package blockchain

import (
	"bytes"
	"cmd-gram-blockchain/internal/models"
	"crypto/sha256"
	"encoding/json"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int64
}

// for safety we need to sethash thats why blocks immutable
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(msg models.MessageDTO, prevBlockHash []byte) (*Block, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block, nil
}
