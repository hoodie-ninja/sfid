package sfid_test

import (
	"encoding/json"
	"testing"

	"github.com/hoodie-ninja/sfid"
	"github.com/stretchr/testify/assert"
)

func TestMarshalText(t *testing.T) {
	b, err := sfid.MustParse("001A0000006Vm9u").MarshalText()
	assert.NoError(t, err)
	assert.Exactly(t, "001A0000006Vm9uIAC", string(b))

	b, err = sfid.ID{}.MarshalText()
	assert.NoError(t, err)
	assert.Exactly(t, "000000000000000AAA", string(b))
}

func TestUnmarshalText(t *testing.T) {
	cases := []struct {
		name, input, want string
		ok                bool
	}{
		{"15-rune", "001A0000006Vm9u", "001A0000006Vm9uIAC", true},
		{"18-rune", "001a0000006vm9uiac", "001A0000006Vm9uIAC", true},
		{"zero", "000000000000000AAA", "000000000000000AAA", true},
		{"invalid", "nope", "", false},
		{"empty", "", "", false},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var id sfid.ID
			err := id.UnmarshalText([]byte(tt.input))
			if !tt.ok {
				assert.Error(t, err)
				assert.Exactly(t, sfid.ID{}, id, "unchanged on error")
				return
			}
			assert.NoError(t, err)
			assert.Exactly(t, tt.want, id.String())
		})
	}
}

func TestJSONRoundTrip(t *testing.T) {
	type record struct {
		ID sfid.ID `json:"id"`
	}
	orig := sfid.MustParse("001A0000006Vm9u")

	out, err := json.Marshal(record{ID: orig})
	assert.NoError(t, err)
	assert.Exactly(t, `{"id":"001A0000006Vm9uIAC"}`, string(out))

	// the 18-, 15-, lowercased, and absent forms all unmarshal to the expected ID
	for _, tt := range []struct {
		in   string
		want sfid.ID
	}{
		{`{"id":"001A0000006Vm9uIAC"}`, orig},
		{`{"id":"001A0000006Vm9u"}`, orig},
		{`{"id":"001a0000006vm9uiac"}`, orig},
		{`{}`, sfid.ID{}},
	} {
		var r record
		assert.NoError(t, json.Unmarshal([]byte(tt.in), &r))
		assert.Exactly(t, tt.want, r.ID)
	}

	// the zero ID marshals to the canonical zero string
	out, err = json.Marshal(record{})
	assert.NoError(t, err)
	assert.Exactly(t, `{"id":"000000000000000AAA"}`, string(out))
}
