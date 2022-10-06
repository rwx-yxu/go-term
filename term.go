package term

import "fmt"

func Color256(i int) string {

	return fmt.Sprintf("\033[38;5;%vm", i)
}
