package blockchain

import (
	"time"

	"github.com/sg-milad/stupid-blockchain/internal/types"
	"github.com/sg-milad/stupid-blockchain/pkg/pow"
)

// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash []byte) *types.Block {
	block := &types.Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := pow.NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *types.Block {
	return NewBlock("Genesis Block", []byte{})
}
