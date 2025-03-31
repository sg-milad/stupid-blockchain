package cli

import (
	"fmt"
	"log"

	"github.com/sg-milad/stupid-blockchain/internal/blockchain"
	"github.com/sg-milad/stupid-blockchain/internal/wallet"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := blockchain.CreateBlockchain(address, nodeID)
	defer bc.DB.Close()

	UTXOSet := blockchain.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
