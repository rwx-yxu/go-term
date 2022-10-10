package sample

import (
	"fmt"

	"github.com/rwx-yxu/term"
	"github.com/rwx-yxu/term/sequence"
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
	case "BkgColors":
		fmt.Printf("%v", BkgColors())
	case "BrightBkgColors":
		fmt.Printf("%v", BBkgColors())
	case "BkgColors256":
		BkgColors256()
	case "Decoration":
		fmt.Printf("%v", Decoration())
	default:
		fmt.Println("Please enter one of the following: Colors, BrightColors,  256Colors, BkgColors, BkgColors256 or Decoration")
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

func BkgColors() string {
	return fmt.Sprintf("%v  %v  %v  %v  %v  %v  %v  %v  %v  \n", sequence.BkgBlack, sequence.BkgRed, sequence.BkgGreen, sequence.BkgYellow, sequence.BkgBlue, sequence.BkgMegenta, sequence.BkgCyan, sequence.BkgWhite, sequence.Reset)
}

func BBkgColors() string {
	return fmt.Sprintf("%v  %v  %v  %v  %v  %v  %v  %v  %v  \n", sequence.BBkgBlack, sequence.BBkgRed, sequence.BBkgGreen, sequence.BBkgYellow, sequence.BBkgBlue, sequence.BBkgMegenta, sequence.BBkgCyan, sequence.BBkgWhite, sequence.Reset)
}

func BkgColors256() {
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			code := i*16 + j
			color := term.BkgColor256(code)
			fmt.Printf("%v   ", color)
		}
		fmt.Printf(" %v \n", sequence.Reset)

	}
	fmt.Println(sequence.Reset)
}

func Decoration() string {
	return fmt.Sprintf("%v BOLD %v %v Underline %v %v Reversed %v\n", sequence.Bold, sequence.Reset, sequence.ULine, sequence.Reset, sequence.Reverse, sequence.Reset)

}
