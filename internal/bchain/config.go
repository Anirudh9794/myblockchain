package bchain

const (
	INIITIAL_DIFFICULTY = 1
	AVERAGE_MINIG_TIME  = 1000
	NUM_BLOCKS_TO_MINE  = 10000
)

var (
	genesisBlock    Block
	LocalBlockchain *Blockchain
)
