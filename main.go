package main

import (
	"fmt"
	cli "github.com/dtan4/imagemv/cli"
	"os"
)

func main() {
	c := cli.New(os.Stdout, os.Stderr)

	if err := c.Run(os.Args); err != nil {
		if os.Getenv("DEBUG") == "1" {
			fmt.Fprintf(os.Stderr, "%+v\n", err)
		} else {
			fmt.Fprintln(os.Stderr, err)
		}

		os.Exit(1)
	}
}
