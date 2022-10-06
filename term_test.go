package term_test

import (
	"fmt"

	"github.com/rwx-yxu/go-term"
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
