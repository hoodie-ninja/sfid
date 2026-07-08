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
		short, str string
		ok         bool
	}{
		{"empty", "", "", "", false},
		{"too short", "123", "", "", false},
		{"length 16", "1234567890123456", "", "", false},
		{"length 19", "1234567890123456789", "", "", false},
		{"whitespace 15", "\t    \n    \t    ", "", "", false},
		{"whitespace 18", "\t    \n    \t    \t\t\t", "", "", false},
		{"invalid rune 15", "12345678901234%", "", "", false},
		{"invalid rune 18", "12345678901234567*", "", "", false},
		{"valid 15", "001A0000006Vm9u", "001A0000006Vm9u", "001A0000006Vm9uIAC", true},
		{"valid 18", "001a0000006vm9uiac", "001A0000006Vm9u", "001A0000006Vm9uIAC", true},
		{"18 uppercase body canonicalizes", "001A0000006VM9UIAC", "001A0000006Vm9u", "001A0000006Vm9uIAC", true},
		{"18 mixed-case body canonicalizes", "001a0000006VM9Uiac", "001A0000006Vm9u", "001A0000006Vm9uIAC", true},
		{"18 with digit checksum", "abaBBaaaaaaaaaa0aa", "aBaBBaaaaaaaaaa", "aBaBBaaaaaaaaaa0AA", true},
		{"zero 15", "000000000000000", "000000000000000", "000000000000000AAA", true},
		{"zero 18", "000000000000000AAA", "000000000000000", "000000000000000AAA", true},
		{"suffix rune 6-9", "001A0000006Vm9uIA6", "", "", false},
		{"suffix all digits", "001A0000006Vm9u999", "", "", false},
		{"suffix flags a digit", "000000000000000BAA", "", "", false},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := sfid.Parse(tt.input)
			assert.Exactly(t, tt.ok, ok)
			if !ok {
				assert.Exactly(t, sfid.ID{}, got, "failed parse returns the zero ID")
				return
			}
			assert.Exactly(t, tt.str, got.String())
			assert.Exactly(t, tt.short, got.Short())
		})
	}
}

func TestParseBytes(t *testing.T) {
	id, ok := sfid.ParseBytes([]byte("001A0000006Vm9u"))
	assert.True(t, ok)
	assert.Exactly(t, "001A0000006Vm9uIAC", id.String())

	_, ok = sfid.ParseBytes([]byte("nope"))
	assert.False(t, ok)
}

func TestMustParse(t *testing.T) {
	assert.Exactly(t, "001A0000006Vm9uIAC", sfid.MustParse("001A0000006Vm9u").String())
	assert.Panics(t, func() { sfid.MustParse("nope") })
}

func TestZeroValue(t *testing.T) {
	var zero sfid.ID
	assert.Exactly(t, "000000000000000AAA", zero.String())
	assert.Exactly(t, "000000000000000", zero.Short())

	// every spelling of the all-zeros ID normalizes to the zero value
	for _, in := range []string{"000000000000000", "000000000000000AAA"} {
		got, ok := sfid.Parse(in)
		assert.True(t, ok)
		assert.Exactly(t, zero, got)
	}
}
