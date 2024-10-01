package main

import (
	"fmt"
	"os"

	"github.com/lutefd/portgen/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
