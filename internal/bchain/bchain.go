package bchain

import (
	"fmt"
	"strings"
	"time"
)

func init() {
	genesisBlock = Block{
		Hash: strings.Repeat("0", INIITIAL_DIFFICULTY) + "gen",
		Timestamp: time.Now().UTC().Format(time.ANSIC),
		Difficulty: INIITIAL_DIFFICULTY,
		Nonce: 0,
		Data: "foo",
	}
}

func Start() {
	// add genesis block
	blockchain = createBlockchain()
	times := []float64{}

	for i := 0 ; i < 2; i += 1 {
		start := time.Now().UTC()
		blockchain.appendBlock(fmt.Sprintf("block %d", i))
		elapsed := time.Since(start).Seconds() * 1e3

		times =append(times, elapsed)

		sum := 0.0
		for _, v := range times {
			sum += v
		}

		addedBlock := blockchain.Chain[len(blockchain.Chain)-1]

		fmt.Println("----")
		fmt.Println("Data:\t", addedBlock.Data)
		fmt.Println("Difficulty:\t", addedBlock.Difficulty)
		fmt.Println("Nonce:\t", addedBlock.Nonce)
		fmt.Println("Average time: \t", sum / float64(len(times)))
		fmt.Println("Current time: \t", elapsed, " ms")
		fmt.Println("---")

		for _, blk := range blockchain.Chain {
			fmt.Println("hash: ", blk.Hash, "LastHash: ", blk.LastHash)
		}
		fmt.Println(blockchain)
	}
}
