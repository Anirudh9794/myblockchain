package bchain

import "testing"

func TestCreateBlock(t *testing.T)  {
	data := "dummyData"
	lastBlock := genesisBlock

	createdBlock := createBlock(data, lastBlock)

	t.Run("data is correct", func(t *testing.T) {
		if createdBlock.Data != data {
			t.Fail()
		}
	})

	t.Run("LastHash is correct", func(t *testing.T) {
		if createdBlock.LastHash != lastBlock.Hash {
			t.Fail()
		}
	})

	t.Run("Timestamp is populated", func(t *testing.T) {
		if createdBlock.Timestamp == "" {
			t.Fail()
		}
	})

	t.Run("Hash is populated", func(t *testing.T) {
		if createdBlock.Hash == "" {
			t.Fail()
		}
	})
}
