package cli

import (
	"fmt"
	"log"

	"github.com/sg-milad/stupid-blockchain/internal/blockchain"
	"github.com/sg-milad/stupid-blockchain/internal/network"
	"github.com/sg-milad/stupid-blockchain/internal/transaction"
	"github.com/sg-milad/stupid-blockchain/internal/wallet"
)

func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !wallet.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !wallet.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := blockchain.UTXOSet{bc}
	defer bc.DB.Close()

	wallets, err := wallet.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := transaction.NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := transaction.NewCoinbaseTX(from, "")
		txs := []*transaction.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		network.SendTx(network.knownNodes[0], tx)
	}

	fmt.Println("Success!")
}
