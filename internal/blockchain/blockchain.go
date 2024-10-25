package blockchain

import (
	"github.com/sg-milad/stupid-blockchain/internal/types"
)

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	Blocks []*types.Block
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*types.Block{NewGenesisBlock()}}
}
