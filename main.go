package main

import (
	"fmt"
	"os"
)

func main() {
	output, err := Run(os.Args)
	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		return
	}
	fmt.Println(output)
}

func Run(args []string) (string, error) {
	if len(args) < 1 {
	}
	var out string
	for _, arg := range args {
		out += arg + "\n"
	}
	return out, nil
}
