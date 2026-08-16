// stubber
// Original name: src/helpers/misc_test.go

package helpers

import "testing"

func TestExtractMajorMinorVersionString(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"1.20.3", "1.20"},
		{"1.33", "1.33"},
		{"1.2.3.4", "1.2"},
		{"1.26.6", "1.26"},
		{"1", "1"},
		{"", ""},
	}

	for _, c := range cases {
		if got := ExtractMajorMinorVersionString(c.in); got != c.want {
			t.Errorf("ExtractMajorMinorVersionString(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
