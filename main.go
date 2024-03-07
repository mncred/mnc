package main

import (
	"fmt"
	"os"

	"github.com/mncred/mnc/cmd"
)

func main() {
	if err := cmd.Register().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
