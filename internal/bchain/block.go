package bchain

import (
	"fmt"
	"strings"
	"time"

	"github.com/Anirudh9794/myblockchain/internal/pkg/hash"
)

type Block struct {
	Hash       string
	LastHash   string
	Data       string
	Timestamp  string
	Nonce      int
	Difficulty int
}

func adjustDifficulty(lastBlock Block) int {
	parsedTime, _ := time.Parse(time.ANSIC, lastBlock.Timestamp)
	timeDiff := time.Since(parsedTime).Seconds() * 1e3
	difficulty := lastBlock.Difficulty

	if timeDiff > AVERAGE_MINIG_TIME {
		return difficulty - 1
	}

	return difficulty + 1
}

func mineBlock(data string, lastBlock Block) Block {
	timestamp := time.Now().UTC().Format(time.ANSIC)
	nonce := 0
	h := ""
	difficulty := lastBlock.Difficulty

	for {
		timestamp = time.Now().UTC().Format(time.ANSIC)

		h = hash.CreateHash(data, fmt.Sprintf("%d", nonce), lastBlock.Hash, timestamp)
		if strings.HasPrefix(hash.HexToBinary(h), strings.Repeat("0", difficulty)) {
			break
		}

		nonce++
		difficulty = adjustDifficulty(lastBlock)
	}

	newBlock := Block{
		Data:       data,
		LastHash:   lastBlock.Hash,
		Timestamp:  timestamp,
		Hash:       h,
		Nonce:      nonce,
		Difficulty: difficulty,
	}

	return newBlock
}
