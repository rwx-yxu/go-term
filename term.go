package term

import "fmt"

func Color256(i int) string {
	return fmt.Sprintf("\033[38;5;%vm", i)
}

func BkgColor256(i int) string {
	return fmt.Sprintf("\033[48;5;%vm", i)
}

//Returns Ansi escape string to move cursor up by i lines
func MvUp(i int) string {
	return fmt.Sprintf("\033[{%v}A", i)
}

func MvDown(i int) string {
	return fmt.Sprintf("\033[{%v}B", i)
}

func MvRight(i int) string {
	return fmt.Sprintf("\033[{%v}C", i)
}

func MvLeft(i int) string {
	return fmt.Sprintf("\033[{%v}D", i)
}
