package main

import (
	"os"

	"github.com/golang-blockchain/part-5/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.run()
}
