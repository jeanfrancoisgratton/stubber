// stubber
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/createAssets/updateversion.go
// Original time: 2023/08/04 16:11

package createAssets

import cerr "github.com/jeanfrancoisgratton/customError/v3"

func UpdateVersion() *cerr.CustomError {
	return &cerr.CustomError{Fatality: cerr.Undefined, Title: "Not implemented", Message: "UpdateVersion() is not yet implemented.", Code: -2}
}
