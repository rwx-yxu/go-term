package sample_test

import (
	"fmt"

	"github.com/rwx-yxu/term/sample"
)

func ExampleTextColors() {
	out := sample.Colors()
	fmt.Printf("%q", out)
	//Output:
	//"\x1b[30m A \x1b[31m B \x1b[32m C \x1b[33m D \x1b[34m E \x1b[35m F \x1b[36m G \x1b[37m H \x1b[0m \n"
}

func ExampleBrightColors() {
	out := sample.BrightColors()
	fmt.Printf("%q", out)
	//Output:
	//"\x1b[30;1m A \x1b[31;1m B \x1b[32;1m C \x1b[33;1m D \x1b[34;1m E \x1b[35;1m F \x1b[36;1m G \x1b[37;1m H \x1b[0m \n"

}

func ExampleBkgColors() {
	out := sample.BkgColors()
	fmt.Printf("%q", out)
	//Output:
	//"\x1b[40m  \x1b[41m  \x1b[42m  \x1b[43m  \x1b[44m  \x1b[45m  \x1b[46m  \x1b[47m  \x1b[0m  \n"

}

func ExampleBBkgColors() {
	out := sample.BBkgColors()
	fmt.Printf("%q", out)
	//Output:
	//"\x1b[40;1m  \x1b[41;1m  \x1b[42;1m  \x1b[43;1m  \x1b[44;1m  \x1b[45;1m  \x1b[46;1m  \x1b[47;1m  \x1b[0m  \n"

}

func ExampleRun_fail() {
	sample.Run("Nothing")
	//Output:
	//Please enter one of the following:
	//  - Colors
	//  - BrightColors
	//  - 256Colors
	//  - BkgColors
	//  - BrightBkgColors
	//  - BkgColors256
	//  - Decoration
	//  - Loading
}

func ExampleDecoration() {
	out := sample.Decoration()
	fmt.Printf("%q", out)
	//Output:
	//"\x1b[1m BOLD \x1b[0m \x1b[4m Underline \x1b[0m \x1b[7m Reversed \x1b[0m\n"

}
