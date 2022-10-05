package term

import (
	"fmt"

	"github.com/rwx-yxu/go-term/sequence"
)

func TextColors() {
	fmt.Printf("%v A %v B %v C %v D %v E %v F %v G %v H %v \n", sequence.Black, sequence.Red, sequence.Green, sequence.Yellow, sequence.Blue, sequence.Megenta, sequence.Cyan, sequence.White, sequence.Reset)
}

func TextBrightColors() {
	fmt.Printf("%v A %v B %v C %v D %v E %v F %v G %v H %v \n", sequence.BBlack, sequence.BRed, sequence.BGreen, sequence.BYellow, sequence.BBlue, sequence.BMegenta, sequence.BCyan, sequence.BWhite, sequence.Reset)

}
