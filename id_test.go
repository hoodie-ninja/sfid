package sfid_test

import (
	"testing"

	"github.com/hoodie-ninja/sfid"
	"github.com/stretchr/testify/assert"
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
		{"valid 15-character ID", "001A0000006Vm9u", "001A0000006Vm9u", "001A0000006Vm9uIAC", true},
		{"valid 18-character ID", "001a0000006vm9uiac", "001A0000006Vm9u", "001A0000006Vm9uIAC", true},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual, ok := sfid.Parse(tt.input)
			assert.Exactly(t, tt.ok, ok)
			assert.Exactly(t, tt.expected18, actual.String())
			assert.Exactly(t, tt.expected18, actual.CaseSafe())
			assert.Exactly(t, tt.expected15, actual.ID())
		})
	}
}
