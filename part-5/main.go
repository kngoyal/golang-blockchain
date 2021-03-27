package main

import (
	"os"

	"github.com/golang-blockchain/part-5/wallet"
)

func main() {
	defer os.Exit(0)
	// cmd := cli.CommandLine{}
	// cmd.Run()

	w := wallet.MakeWallet()
	w.Address()
}
