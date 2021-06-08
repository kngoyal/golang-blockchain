package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/golang-blockchain/part-3/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	// chain.AddBlock("1st Block after Genesis")
	// chain.AddBlock("2nd Block after Genesis")
	// chain.AddBlock("3rd Block after Genesis")

	iter := chain.Iterator()

	for block := iter.Previous(); bytes.Compare(block.PrevHash, []byte{}) >= 0; block = iter.Previous() {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("POW: %s\n\n", strconv.FormatBool(pow.Validate()))
	}
}
