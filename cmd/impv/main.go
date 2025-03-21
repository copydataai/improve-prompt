package main

import (
	"fmt"
	"github.com/copydataai/improve-prompt/pkg/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
