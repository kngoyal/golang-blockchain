package main

import (
	"fmt"

	"github.com/golang-blockchain/part-2/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("1st Block after Genesis")
	chain.AddBlock("2nd Block after Genesis")
	chain.AddBlock("3rd Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
}
