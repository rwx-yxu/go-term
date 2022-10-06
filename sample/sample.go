package sample

import (
	"fmt"

	"github.com/rwx-yxu/go-term"
	"github.com/rwx-yxu/go-term/sequence"
)

//Run function will ran a sample command based on the argument being
//passed through.
//Supported sample outputs are: Basic text colors, bright colors and 256
//colors
func Run(out string) {

	switch out {
	case "Colors":
		fmt.Printf("%v", Colors())
	case "BrightColors":
		fmt.Printf("%v", BrightColors())
	case "256Colors":
		Colors256()
	default:
		fmt.Println("Please enter one of the following: Colors, BrightColors or 256Colors")
	}

}

func Colors() string {
	return fmt.Sprintf("%v A %v B %v C %v D %v E %v F %v G %v H %v \n", sequence.Black, sequence.Red, sequence.
		Green, sequence.Yellow, sequence.Blue, sequence.Megenta, sequence.Cyan, sequence.White, sequence.Reset)
}

func BrightColors() string {
	return fmt.Sprintf("%v A %v B %v C %v D %v E %v F %v G %v H %v \n", sequence.BBlack, sequence.BRed, sequence.BGreen, sequence.BYellow, sequence.BBlue, sequence.BMegenta, sequence.BCyan, sequence.BWhite, sequence.Reset)

}

func Colors256() {
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			code := i*16 + j
			color := term.Color256(code)
			fmt.Printf("%v %v", color, code)
		}
		fmt.Print("\n")
	}
	fmt.Println(sequence.Reset)

}
