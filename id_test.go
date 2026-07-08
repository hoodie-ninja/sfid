package sfid_test

import (
	"testing"

	"github.com/hoodie-ninja/sfid"
	"github.com/stretchr/testify/assert"
)

// canonical fixture ID shared across the package's test files
const (
	testID15 = "001A0000006Vm9u"
	testID18 = "001A0000006Vm9uIAC"
)

func TestParse(t *testing.T) {
	cases := []struct {
		name       string
		input      string
		expected15 string
		expected18 string
		ok         bool
	}{
		{"empty input", "", "", "", false},
		{"bad length short", "123", "", "", false},
		{"bad length mid", "1234567890123456", "", "", false},
		{"bad length long", "1234567890123456789", "", "", false},
		{"whitespace15 only", "\t    \n    \t    ", "", "", false},
		{"whitespace18 only", "\t    \n    \t    \t\t\t", "", "", false},
		{"invalid characters (15)", "12345678901234%", "", "", false},
		{"invalid characters (18)", "12345678901234567*", "", "", false},
		{"valid 15-character ID", testID15, testID15, testID18, true},
		{"valid 18-character ID", "001a0000006vm9uiac", testID15, testID18, true},
		{"18-character ID uppercase body canonicalizes", "001A0000006VM9UIAC", testID15, testID18, true},
		{"18-character ID mixed-case body canonicalizes", "001a0000006VM9Uiac", testID15, testID18, true},
		{"18-character ID with digit checksum", "abaBBaaaaaaaaaa0aa", "aBaBBaaaaaaaaaa", "aBaBBaaaaaaaaaa0AA", true},
		{"invalid suffix alphabet (6-9)", "001A0000006Vm9uIA6", "", "", false},
		{"invalid suffix alphabet all digits", "001A0000006Vm9u999", "", "", false},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual, ok := sfid.Parse(tt.input)
			assert.Exactly(t, tt.ok, ok)
			assert.Exactly(t, tt.expected18, actual.String())
			assert.Exactly(t, tt.expected15, actual.ID())
		})
	}
}
