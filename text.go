package sfid

import (
	"encoding"
	"fmt"
)

var (
	_ encoding.TextMarshaler   = ID{}
	_ encoding.TextUnmarshaler = (*ID)(nil)
)

// MarshalText implements encoding.TextMarshaler. Returns the 18-rune form.
func (id ID) MarshalText() ([]byte, error) { return []byte(id.String()), nil }

// UnmarshalText implements encoding.TextUnmarshaler. Accepts a 15- or 18-rune forms.
func (id *ID) UnmarshalText(text []byte) error {
	parsed, ok := ParseBytes(text)
	if !ok {
		return fmt.Errorf("sfid: cannot unmarshal %q into ID: invalid Salesforce ID", text)
	}
	*id = parsed
	return nil
}
