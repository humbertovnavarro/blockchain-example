package main

import (
	"fmt"

	"github.com/humbertovnavarro/blockchain-example/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block After Genesis")
	chain.AddBlock("Second Block After Genesis")
	chain.AddBlock("third Block After Genesis")
	chain.AddBlock("Fourth Block After Genesis")
	for _, block := range chain.Blocks {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Print("\n")
	}
}
