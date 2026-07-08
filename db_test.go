package sfid_test

import (
	"testing"

	"github.com/hoodie-ninja/sfid"
	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	id, ok := sfid.Parse(testID15)
	assert.True(t, ok)

	v, err := id.Value()
	assert.NoError(t, err)
	assert.Exactly(t, testID18, v)

	// the zero ID cannot be stored.
	var zero sfid.ID
	v, err = zero.Value()
	assert.Error(t, err)
	assert.Nil(t, v)
}

func TestScan(t *testing.T) {
	cases := []struct {
		name string
		src  any
		ok   bool
	}{
		{"string 15-rune", testID15, true},
		{"string 18-rune", "001a0000006vm9uiac", true},
		{"bytes", []byte(testID15), true},
		{"invalid string bad length", "123", false},
		{"invalid string bad char", "12345678901234%", false},
		{"invalid bytes", []byte("nope"), false},
		{"null", nil, false},
		{"unsupported type", int64(1), false},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var id sfid.ID
			err := id.Scan(tt.src)
			if !tt.ok {
				assert.Error(t, err)
				assert.Exactly(t, "", id.String(), "id must be unchanged on error")
				return
			}
			assert.NoError(t, err)
			assert.Exactly(t, testID18, id.String())
		})
	}
}

func TestValueScanRoundTrip(t *testing.T) {
	original, ok := sfid.Parse(testID15)
	assert.True(t, ok)

	v, err := original.Value()
	assert.NoError(t, err)

	var scanned sfid.ID
	assert.NoError(t, scanned.Scan(v))

	// IDs are directly comparable and must round-trip losslessly.
	assert.True(t, original == scanned)
}
