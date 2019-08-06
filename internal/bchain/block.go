package bchain

import (
	"github.com/Anirudh9794/myblockchain/internal/pkg/hash"
	"time"
)

type Block struct {
	Hash string
	LastHash string
	Data string
	Timestamp string
}

func createBlock(data string, lastBlock Block) Block {

	newBlock := Block{
		Data: data,
		LastHash: lastBlock.Hash,
		Timestamp: time.Now().String(),
		Hash: hash.CreateHash(data, lastBlock.Hash),
	}

	return newBlock
}
