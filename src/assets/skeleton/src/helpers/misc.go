// {{ SOFTWARE NAME }}
// src/misc/misc.go

package helpers

import (
	"fmt"
	"github.com/jwalton/gchalk"
)

func Changelog() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")

	fmt.Print(`
VERSION			DATE			COMMENT
-------			----			-------
{{ PACKAGE VERSION  }}	2023.xx.yy		stub
`)
}

func Red(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightRed().Bold(sentence))
}

func Green(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightGreen().Bold(sentence))
}

func White(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightWhite().Bold(sentence))
}

func Yellow(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightYellow().Bold(sentence))
}

// FIXME : Normal() is the same as White()
func Normal(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithWhite().Bold(sentence))
}
