package types

import (
	"github.com/sg-milad/stupid-blockchain/internal/transaction"
)

// Transaction represents a Bitcoin transaction
type Transaction struct {
	ID   []byte
	Vin  []transaction.TXInput
	Vout []transaction.TXOutput
}

// Block represents a block in the blockchain
type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Height        int
}
