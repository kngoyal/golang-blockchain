package main

import (
	"os"

	"github.com/golang-blockchain/part-5/logger"
	"github.com/golang-blockchain/part-5/cli"
)

func main() {
	logger.SetLogLevel()
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
