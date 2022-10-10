package term

import "fmt"

//Color256 will return a formatted ansi escape code for a 256 color
//sequence
func Color256(i int) string {
	return fmt.Sprintf("\033[38;5;%vm", i)
}

//BkgColor256 will return a formatted ansi escape code for a 256
//background color sequence
func BkgColor256(i int) string {
	return fmt.Sprintf("\033[48;5;%vm", i)
}

//MvUp will return a formatted ansi escape code for moving the cursor up
//by i lines.
func MvUp(i int) string {
	return fmt.Sprintf("\033[{%v}A", i)
}

//MvDown will return a formatted ansi escape code for moving the cursor
//down by i lines
func MvDown(i int) string {
	return fmt.Sprintf("\033[{%v}B", i)
}

//MvRight will return a formatted ansi escape code for moving the cursor
//right by i steps.
func MvRight(i int) string {
	return fmt.Sprintf("\033[{%v}C", i)
}

//MvLeft will return a formatted ani escape code for moving the cursor
//to the left by i steps.
func MvLeft(i int) string {
	return fmt.Sprintf("\033[{%v}D", i)
}
