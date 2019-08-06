package bchain

type Blockchain struct {
	Chain []Block
}

func createBlockchain() *Blockchain{

	bchain := Blockchain{
		Chain: []Block{genesisBlock},
	}

	return &bchain

}

func (bchain *Blockchain) appendBlock(data string) {

	lastBlock := bchain.Chain[len(bchain.Chain)-1]

	newBlock := createBlock(data, lastBlock)

	bchain.Chain = append(bchain.Chain, newBlock)
}
