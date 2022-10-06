package main

import (
	"os"

	"github.com/rwx-yxu/go-term/sample"
)

func main() {
	if os.Args[1] != "" {

		sample.Run(os.Args[1])
	}
}
