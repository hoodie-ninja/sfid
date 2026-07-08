package sfid_test

import (
	"encoding/json"
	"testing"

	"github.com/hoodie-ninja/sfid"
	"github.com/stretchr/testify/assert"
)

func TestMarshalText(t *testing.T) {
	id, ok := sfid.Parse(testID15)
	assert.True(t, ok)

	b, err := id.MarshalText()
	assert.NoError(t, err)
	assert.Exactly(t, testID18, string(b))

	// the zero ID cannot be marshaled.
	var zero sfid.ID
	b, err = zero.MarshalText()
	assert.Error(t, err)
	assert.Nil(t, b)
}

func TestUnmarshalText(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
		ok    bool
	}{
		{"15-rune", testID15, testID18, true},
		{"18-rune", "001a0000006vm9uiac", testID18, true},
		{"invalid", "nope", "", false},
		{"empty", "", "", false},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var id sfid.ID
			err := id.UnmarshalText([]byte(tt.input))
			if !tt.ok {
				assert.Error(t, err)
				assert.Exactly(t, "", id.String(), "id must be unchanged on error")
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

	original, ok := sfid.Parse(testID15)
	assert.True(t, ok)

	out, err := json.Marshal(record{ID: original})
	assert.NoError(t, err)
	assert.Exactly(t, `{"id":"`+testID18+`"}`, string(out))

	var in record
	assert.NoError(t, json.Unmarshal(out, &in))
	assert.True(t, original == in.ID)

	// 15-rune input on the wire parses to the same ID.
	assert.NoError(t, json.Unmarshal([]byte(`{"id":"`+testID15+`"}`), &in))
	assert.True(t, original == in.ID)

	// a struct with a zero ID cannot be marshaled.
	_, err = json.Marshal(record{})
	assert.Error(t, err)
}
