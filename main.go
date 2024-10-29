package main

import (
	cli "github.com/sg-milad/stupid-blockchain/cmd/blockchain"
)

func main() {
	cli := cli.CLI{}
	cli.Run()
}
