package bchain

import (
	"fmt"
	"reflect"

	"github.com/Anirudh9794/myblockchain/internal/pkg/hash"
)

type Blockchain struct {
	Chain []Block
}

func CreateBlockchain() *Blockchain {
	bchain := Blockchain{
		Chain: []Block{genesisBlock},
	}

	return &bchain

}

func (bchain *Blockchain) AppendBlock(data string) {
	lastBlock := bchain.Chain[len(bchain.Chain)-1]

	newBlock := mineBlock(data, lastBlock)

	bchain.Chain = append(bchain.Chain, newBlock)
}

func IsValidChain(chain []Block) bool {
	// should start with genesis block
	if !reflect.DeepEqual(chain[0], genesisBlock) {
		fmt.Println("genesis block is not valid")
		return false
	}

	previousBlock := chain[0]

	for _, block := range chain[1:] {
		// lastHash is not correct
		if block.LastHash != previousBlock.Hash {
			fmt.Println("LastHash doesnt match for block ", block.Data)
			return false
		}

		// hash is not correct
		actualHash := block.Hash
		expectedHash := hash.CreateHash(block.Data, fmt.Sprintf("%d", block.Nonce), previousBlock.Hash, block.Timestamp)

		if actualHash != expectedHash {
			fmt.Println("Hash doesn't match for block ", block.Data)
			fmt.Println("Expected: ", expectedHash, " Actual:", actualHash)
			return false
		}

		previousBlock = block
	}

	return true
}

func (bchain *Blockchain) ReplaceChain(chain []Block) {
	// do not replace if new chain is shorter or if the new chain is invalid
	if len(chain) < len(bchain.Chain) || !IsValidChain(chain) {
		return
	}

	bchain.Chain = chain
}
