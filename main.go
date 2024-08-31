package main

import (
	"fmt"
	"os"
)

func main() {
	output, err := Run(os.Args[1:])
	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
	print(output)
}
