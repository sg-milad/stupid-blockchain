package cli

import (
	"fmt"
	"log"

	"github.com/sg-milad/stupid-blockchain/internal/blockchain"
	"github.com/sg-milad/stupid-blockchain/internal/wallet"
	"github.com/sg-milad/stupid-blockchain/pkg"
)

func (cli *CLI) getBalance(address, nodeID string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := blockchain.UTXOSet{bc}
	defer bc.DB.Close()

	balance := 0
	pubKeyHash := pkg.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
