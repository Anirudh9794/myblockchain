package bchain

import (
	"fmt"
	"time"
)

func init() {
	genesisBlock = Block{
		Hash: "newHash",
		LastHash: "lastHash",
		Data: "dummydata",
		Timestamp: time.Now().String(),
	}

	// add genesis block
	blockchain = createBlockchain()
}

func Start() {
	blockchain.appendBlock("twoBlock")
	blockchain.appendBlock("threeBlock")
	blockchain.appendBlock("oneBlock")

	fmt.Println("blockchain: ", blockchain)
}
