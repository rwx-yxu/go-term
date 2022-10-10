package main

import (
	"log"
	"os"

	"github.com/rwx-yxu/term/sample"
)

func main() {
	if len(os.Args) == 1 {

		log.Println("Please rerun program with a sample command argument.")
		return
	}

	sample.Run(os.Args[1])

}
