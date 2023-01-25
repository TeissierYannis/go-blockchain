package main

import (
	"fr.YT.blockchain/cli"
	"os"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
