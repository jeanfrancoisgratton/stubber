// stubber : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/misc/misc.go
// 2023.06.25 8:58:03

package helpers

import (
	"strings"
)

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
