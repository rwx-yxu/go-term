package term_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/rwx-yxu/term"
)

func ExampleColor256() {
	color := term.Color256(50)
	fmt.Printf("%q \n", color)
	//Output:
	//"\x1b[38;5;50m"
}

func ExampleBkgColor256() {
	bkg := term.BkgColor256(50)
	fmt.Printf("%q \n", bkg)

	//Output:
	//"\x1b[48;5;50m"
}

func ExampleMvUp() {
	out := term.MvUp(2)
	fmt.Printf("%q\n", out)
	//Output:
	//"\x1b[2A"
}

func ExampleMvDown() {
	out := term.MvDown(2)
	fmt.Printf("%q\n", out)
	//Output:
	//"\x1b[2B"
}

func ExampleMvRight() {
	out := term.MvRight(2)
	fmt.Printf("%q\n", out)
	//Output:
	//"\x1b[2C"
}

func ExampleMvLeft() {
	out := term.MvLeft(2)
	fmt.Printf("%q\n", out)
	//Output:
	//"\x1b[2D"
}

func ExampleReadLine() {
	//Turn string into standard input
	//As if it is standard input
	sr := strings.NewReader("Sample\r\n")
	line, err := term.ReadLine(sr)
	if err != nil {
		log.Println(err)
		return
	}
	//Using %s will pass
	fmt.Printf("%q", line)
	//Output:
	//"Sample"

}

func ExamplePrompt_default() {
	//Turn string into standard input
	//As if it is standard input
	sr := strings.NewReader("Sample\r\n")
	line, err := term.Prompt(sr, "")
	if err != nil {
		log.Println(err)
		return
	}
	//Using %s will pass
	fmt.Printf("%q", line)
	//Output:
	//> "Sample"
}

func ExamplePrompt_explicit() {
	//Turn string into standard input
	//As if it is standard input
	sr := strings.NewReader("Sample\r\n")
	line, err := term.Prompt(sr, "--> ")
	if err != nil {
		log.Println(err)
		return
	}
	//Using %s will pass
	fmt.Printf("%q", line)
	//Output:
	//--> "Sample"
}
