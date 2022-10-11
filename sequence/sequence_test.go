package sequence_test

import (
	"fmt"

	"github.com/rwx-yxu/term/sequence"
)

func ExampleColors_basic() {
	fmt.Printf("%q\n", sequence.Black)
	fmt.Printf("%q\n", sequence.Red)
	fmt.Printf("%q\n", sequence.Green)
	fmt.Printf("%q\n", sequence.Yellow)
	fmt.Printf("%q\n", sequence.Blue)
	fmt.Printf("%q\n", sequence.Megenta)
	fmt.Printf("%q\n", sequence.Cyan)
	fmt.Printf("%q\n", sequence.White)
	fmt.Printf("%q\n", sequence.Reset)

	//Output:
	//"\x1b[30m"
	//"\x1b[31m"
	//"\x1b[32m"
	//"\x1b[33m"
	//"\x1b[34m"
	//"\x1b[35m"
	//"\x1b[36m"
	//"\x1b[37m"
	//"\x1b[0m"
}

func ExampleColors_bright() {
	fmt.Printf("%q\n", sequence.BBlack)
	fmt.Printf("%q\n", sequence.BRed)
	fmt.Printf("%q\n", sequence.BGreen)
	fmt.Printf("%q\n", sequence.BYellow)
	fmt.Printf("%q\n", sequence.BBlue)
	fmt.Printf("%q\n", sequence.BMegenta)
	fmt.Printf("%q\n", sequence.BCyan)
	fmt.Printf("%q\n", sequence.BWhite)
	fmt.Printf("%q\n", sequence.Reset)

	//Output:
	//"\x1b[30;1m"
	//"\x1b[31;1m"
	//"\x1b[32;1m"
	//"\x1b[33;1m"
	//"\x1b[34;1m"
	//"\x1b[35;1m"
	//"\x1b[36;1m"
	//"\x1b[37;1m"
	//"\x1b[0m"
}

func ExampleBkgColors() {
	fmt.Printf("%q\n", sequence.BkgBlack)
	fmt.Printf("%q\n", sequence.BkgRed)
	fmt.Printf("%q\n", sequence.BkgGreen)
	fmt.Printf("%q\n", sequence.BkgYellow)
	fmt.Printf("%q\n", sequence.BkgBlue)
	fmt.Printf("%q\n", sequence.BkgMegenta)
	fmt.Printf("%q\n", sequence.BkgCyan)
	fmt.Printf("%q\n", sequence.BkgWhite)
	fmt.Printf("%q\n", sequence.Reset)

	//Output:
	//"\x1b[40m"
	//"\x1b[41m"
	//"\x1b[42m"
	//"\x1b[43m"
	//"\x1b[44m"
	//"\x1b[45m"
	//"\x1b[46m"
	//"\x1b[47m"
	//"\x1b[0m"
}

func ExampleBkgColors_bright() {
	fmt.Printf("%q\n", sequence.BBkgBlack)
	fmt.Printf("%q\n", sequence.BBkgRed)
	fmt.Printf("%q\n", sequence.BBkgGreen)
	fmt.Printf("%q\n", sequence.BBkgYellow)
	fmt.Printf("%q\n", sequence.BBkgBlue)
	fmt.Printf("%q\n", sequence.BBkgMegenta)
	fmt.Printf("%q\n", sequence.BBkgCyan)
	fmt.Printf("%q\n", sequence.BBkgWhite)
	fmt.Printf("%q\n", sequence.Reset)

	//Output:
	//"\x1b[40;1m"
	//"\x1b[41;1m"
	//"\x1b[42;1m"
	//"\x1b[43;1m"
	//"\x1b[44;1m"
	//"\x1b[45;1m"
	//"\x1b[46;1m"
	//"\x1b[47;1m"
	//"\x1b[0m"
}

func ExampleDecorations() {
	fmt.Printf("%q\n", sequence.ULine)
	fmt.Printf("%q\n", sequence.Reverse)
	fmt.Printf("%q\n", sequence.Bold)

	//Output:
	//"\x1b[4m"
	//"\x1b[7m"
	//"\x1b[1m"

}

func ExampleClearLine() {
	fmt.Printf("%q\n", sequence.CLEntire)
	fmt.Printf("%q\n", sequence.CLStart)
	fmt.Printf("%q\n", sequence.CLEnd)
	//Output:
	//"\x1b[2K"
	//"\x1b[1K"
	//"\x1b[0K"

}
