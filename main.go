package main

import (
	"os"

	"countbeat/cmd"

	_ "countbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
