package main

import (
	cli "github.com/sg-milad/stupid-blockchain/cmd/blockchain"
	"github.com/sg-milad/stupid-blockchain/internal/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.Db.Close()

	cli := cli.CLI{bc}
	cli.Run()
}
