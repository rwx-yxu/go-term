package sequence_test

import (
	"fmt"

	"github.com/rwx-yxu/go-term/sequence"
)

func ExampleColors_eight() {
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
