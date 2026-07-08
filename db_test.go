package sfid_test

import (
	"testing"

	"github.com/hoodie-ninja/sfid"
	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	v, err := sfid.MustParse("001A0000006Vm9u").Value()
	assert.NoError(t, err)
	assert.Exactly(t, "001A0000006Vm9uIAC", v)

	v, err = sfid.ID{}.Value()
	assert.NoError(t, err)
	assert.Exactly(t, "000000000000000AAA", v)
}

func TestScan(t *testing.T) {
	cases := []struct {
		name string
		src  any
		ok   bool
	}{
		{"string 15", "001A0000006Vm9u", true},
		{"string 18", "001a0000006vm9uiac", true},
		{"bytes", []byte("001A0000006Vm9u"), true},
		{"bad length", "123", false},
		{"bad rune", "12345678901234%", false},
		{"bad bytes", []byte("nope"), false},
		{"NULL", nil, false},
		{"unsupported type", int64(1), false},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var id sfid.ID
			err := id.Scan(tt.src)
			if !tt.ok {
				assert.Error(t, err)
				assert.Exactly(t, sfid.ID{}, id, "unchanged on error")
				return
			}
			assert.NoError(t, err)
			assert.Exactly(t, "001A0000006Vm9uIAC", id.String())
		})
	}
}

func TestValueScanRoundTrip(t *testing.T) {
	for _, id := range []sfid.ID{sfid.MustParse("001A0000006Vm9u"), {}} {
		v, err := id.Value()
		assert.NoError(t, err)

		var got sfid.ID
		assert.NoError(t, got.Scan(v))
		assert.Exactly(t, id, got)
	}
}
