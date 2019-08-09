package bchain

import (
	"fmt"
	"testing"
	"time"
)

func TestMineBlock(t *testing.T)  {
	data := "dummyData"
	lastBlock := genesisBlock
	createdBlock := mineBlock(data, lastBlock)

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
			fmt.Println(createdBlock.Hash)
			t.Fail()
		}
	})
}

func TestAdjustDifficulty(t *testing.T) {
	difficulty := 10

	t.Run("increases difficulty if lastBlock is mined in less time", func(t *testing.T) {
		lastBlock := Block{
			Timestamp: time.Now().UTC().Format(time.ANSIC),
			Difficulty: difficulty,
		}

		adjustedDifficulty := adjustDifficulty(lastBlock)

		if adjustedDifficulty < 10 {
			t.Error("Expected: ", difficulty+1, " Got:", adjustedDifficulty)
		}
	})

	t.Run("decreases difficulty if lastBlock is mined in more time", func(t *testing.T) {
		lastBlock := Block{
			Timestamp: time.Now().UTC().Format(time.ANSIC),
			Difficulty: difficulty,
		}

		time.Sleep(3000 * time.Millisecond)

		adjustedDifficulty := adjustDifficulty(lastBlock)

		if adjustedDifficulty > 10 {
			t.Error("Expected: ", difficulty-1, " Got:", adjustedDifficulty)
		}
	})
}
