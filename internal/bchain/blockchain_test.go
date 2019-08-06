package bchain

import (
	"testing"
)

func TestCreateBlockchain(t *testing.T) {

	testchain := createBlockchain()

	t.Run("blockchain length is not zero", func (t *testing.T){
		if len(testchain.Chain) != 1 {
			t.Fail()
		}
	})

	t.Run("starts with genesis block", func (t *testing.T){
		if testchain.Chain[0] != genesisBlock {
			t.Fail()
		}
	})

}

func TestAppendBlock(t *testing.T) {

	testData := "test"

	testchain := createBlockchain()

	testchain.appendBlock(testData)

	t.Run("length should be greater than 1", func (t *testing.T) {
		if len(testchain.Chain) < 2 {
			t.Fail()
		}
	})

	t.Run("starts with genesis block", func (t *testing.T) {
		if testchain.Chain[0] != genesisBlock {
			t.Fail()
		}
	})

	t.Run("ends with new block", func (t *testing.T) {
		if testchain.Chain[len(testchain.Chain)-1].Data != testData {
			t.Errorf("Expected: %s Actual: %s",testData, testchain.Chain[len(testchain.Chain)-1].Data)
		}
	})
}
