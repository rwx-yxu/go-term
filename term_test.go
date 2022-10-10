package term_test

import (
	"fmt"

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
	//"\x1b[{2}A"
}

func ExampleMvDown() {
	out := term.MvDown(2)
	fmt.Printf("%q\n", out)
	//Output:
	//"\x1b[{2}B"
}

func ExampleMvRight() {
	out := term.MvRight(2)
	fmt.Printf("%q\n", out)
	//Output:
	//"\x1b[{2}C"
}

func ExampleMvLeft() {
	out := term.MvLeft(2)
	fmt.Printf("%q\n", out)
	//Output:
	//"\x1b[{2}D"
}
