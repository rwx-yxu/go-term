package sample

import (
	"fmt"
	"time"

	"github.com/rwx-yxu/term"
	"github.com/rwx-yxu/term/sequence"
)

//Run function will ran a sample command based on the argument being
//passed through.
//Supported sample outputs are:
//Basic text colors
//Bright colors
//256colors
//Background colors
//Bright background colors
//256 background colors
//Decorations
//Loading
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
	case "Loading":
		Loading()
	default:
		fmt.Println("Please enter one of the following:")
		fmt.Println(" - Colors")
		fmt.Println(" - BrightColors")
		fmt.Println(" - 256Colors")
		fmt.Println(" - BkgColors")
		fmt.Println(" - BrightBkgColors")
		fmt.Println(" - BkgColors256")
		fmt.Println(" - Decoration")
		fmt.Println(" - Loading")
	}
}

//Colors returns a formatted colored text for the eight basic terminal
//colors:black, red, green, yellow, megenta,
//cyan and white
func Colors() string {
	return fmt.Sprintf("%v A %v B %v C %v D %v E %v F %v G %v H %v \n", sequence.Black, sequence.Red, sequence.
		Green, sequence.Yellow, sequence.Blue, sequence.Megenta, sequence.Cyan, sequence.White, sequence.Reset)
}

//BrightColors returns a formatted string for the eight basic terminal
//bright colors: black, red, green, yellow, megenta,
//cyan and white
func BrightColors() string {
	return fmt.Sprintf("%v A %v B %v C %v D %v E %v F %v G %v H %v \n", sequence.BBlack, sequence.BRed, sequence.BGreen, sequence.BYellow, sequence.BBlue, sequence.BMegenta, sequence.BCyan, sequence.BWhite, sequence.Reset)

}

//Color256 outputs the 256 colors in a 15x15 grid.
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

//BkgColors will return an ansi escape string that displays the basic
//background colors:black, red, green, yellow, megenta,
//cyan and white

func BkgColors() string {
	return fmt.Sprintf("%v  %v  %v  %v  %v  %v  %v  %v  %v  \n", sequence.BkgBlack, sequence.BkgRed, sequence.BkgGreen, sequence.BkgYellow, sequence.BkgBlue, sequence.BkgMegenta, sequence.BkgCyan, sequence.BkgWhite, sequence.Reset)
}

//BBkgColors will return a formatted ansi escape string that displays
//the bright basic background colors: black, red, green, yellow, megenta,
//cyan and white
func BBkgColors() string {
	return fmt.Sprintf("%v  %v  %v  %v  %v  %v  %v  %v  %v  \n", sequence.BBkgBlack, sequence.BBkgRed, sequence.BBkgGreen, sequence.BBkgYellow, sequence.BBkgBlue, sequence.BBkgMegenta, sequence.BBkgCyan, sequence.BBkgWhite, sequence.Reset)
}

//BkgColors256 will output all 256 colors as background.
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

//Decoration will return a formatted string for bold, underline and
//reversed
func Decoration() string {
	return fmt.Sprintf("%v BOLD %v %v Underline %v %v Reversed %v\n", sequence.Bold, sequence.Reset, sequence.ULine, sequence.Reset, sequence.Reverse, sequence.Reset)

}

func Loading() {
	fmt.Println("Loading...")
	for i := 0; i < 101; i++ {
		fmt.Print(term.MvLeft(1000))
		fmt.Printf("%v%% ", i)
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Print("\n")
}
