package main

import (
	"io"
	"os"
)

func main() {
	fn := os.Args[1]

	file, err := os.Open(fn)
	if err != nil {
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
