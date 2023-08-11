// stubber : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/misc/misc.go
// 2023.06.25 8:58:03

package helpers

import (
	"fmt"
	"github.com/jwalton/gchalk"
	"strings"
)

func Changelog() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")

	fmt.Print(`
VERSION		DATE			COMMENT
-------		----			-------
0.500		2023.08.11		completed apk, deb, rpm, skeleton
0.100		2023.06.25		stub
`)
}

// This is a quick-and-dirty way to extract the Major + Minor number of a version string
// Thus, 1.20.3 would return 1.20, 1.33 would return 1.33, 1.2.3.4 would return 1.2
func ExtractMajorMinorVersionString(versionNum string) string {
	var extractedStr string
	p := strings.SplitN(versionNum, ".", 3)
	if len(p) >= 2 {
		extractedStr = p[0] + "." + p[1]
	} else {
		extractedStr = versionNum
	}
	return extractedStr
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
