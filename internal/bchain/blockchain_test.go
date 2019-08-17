package bchain

import (
	"reflect"
	"testing"
	"time"
)

func generateBlockchain() *Blockchain {
	blkchain := CreateBlockchain()

	blkchain.AppendBlock("myData1")
	blkchain.AppendBlock("myData2")
	blkchain.AppendBlock("myData3")
	blkchain.AppendBlock("myData4")

	return blkchain
}

func TestCreateBlockchain(t *testing.T) {
	testchain := CreateBlockchain()

	t.Run("blockchain length is not zero", func(t *testing.T) {
		if len(testchain.Chain) != 1 {
			t.Fail()
		}
	})

	t.Run("starts with genesis block", func(t *testing.T) {
		if testchain.Chain[0] != genesisBlock {
			t.Fail()
		}
	})
}

func TestAppendBlock(t *testing.T) {

	testData := "test"
	testchain := CreateBlockchain()

	testchain.AppendBlock(testData)

	t.Run("length should be greater than 1", func(t *testing.T) {
		if len(testchain.Chain) < 2 {
			t.Fail()
		}
	})

	t.Run("starts with genesis block", func(t *testing.T) {
		if testchain.Chain[0] != genesisBlock {
			t.Fail()
		}
	})

	t.Run("ends with new block", func(t *testing.T) {
		if testchain.Chain[len(testchain.Chain)-1].Data != testData {
			t.Errorf("Expected: %s Actual: %s", testData, testchain.Chain[len(testchain.Chain)-1].Data)
		}
	})
}

func TestIsValidChain(t *testing.T) {
	t.Run("Does not start with genesis block", func(t *testing.T) {
		blkchain := generateBlockchain()
		invalidGenesis := Block{
			Hash:      "my-dummy-hash",
			LastHash:  "my-dummy-last-hash",
			Data:      "invalid-data",
			Timestamp: time.Now().String(),
		}

		blkchain.Chain[0] = invalidGenesis

		if IsValidChain(blkchain.Chain) {
			t.Fail()
		}
	})

	t.Run("A block contains invalid LastHash", func(t *testing.T) {
		blkchain := generateBlockchain()

		blkchain.Chain[2].LastHash = "wrong-hash"

		if IsValidChain(blkchain.Chain) {
			t.Fail()
		}
	})

	t.Run("Malicious block in the chain", func(t *testing.T) {
		blkchain := generateBlockchain()
		blkchain.Chain[2].Data = "manipulated-data"

		if IsValidChain(blkchain.Chain) {
			t.Fail()
		}
	})

	t.Run("All blocks are valid", func(t *testing.T) {
		blkchain := generateBlockchain()

		if !IsValidChain(blkchain.Chain) {
			t.Fail()
		}
	})
}

func TestBlockchain_ReplaceChain(t *testing.T) {
	t.Run("Fail if input chain is smaller", func(t *testing.T) {
		blkchain := generateBlockchain()
		smallerchain := CreateBlockchain()

		smallerchain.AppendBlock("small-data1")

		blkchain.ReplaceChain(smallerchain.Chain)

		if reflect.DeepEqual(smallerchain.Chain, blkchain.Chain) {
			t.Fail()
		}
	})

	t.Run("Fail if input chain is invalid", func(t *testing.T) {
		blkChain := generateBlockchain()
		newChain := generateBlockchain()

		newChain.AppendBlock("newData2")
		newChain.Chain[2].Data = "invalid data"

		blkChain.ReplaceChain(newChain.Chain)

		if reflect.DeepEqual(newChain.Chain, blkChain.Chain) {
			t.Fail()
		}
	})

	t.Run("Succeed if input chain is longer and valid", func(t *testing.T) {
		blkChain := generateBlockchain()
		newChain := generateBlockchain()

		newChain.AppendBlock("newData2")

		blkChain.ReplaceChain(newChain.Chain)

		if !reflect.DeepEqual(newChain.Chain, blkChain.Chain) {
			t.Fail()
		}
	})
}
