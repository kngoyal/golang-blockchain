package main

import (
	"os"

	"github.com/golang-blockchain/part-7/cli"
	"github.com/golang-blockchain/part-7/utils"
)

func main() {
	utils.SetLogLevel()
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
