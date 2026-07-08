package sfid

import (
	"encoding"
	"errors"
	"fmt"
)

var (
	_ encoding.TextMarshaler   = ID{}
	_ encoding.TextUnmarshaler = (*ID)(nil)
)

// MarshalText implements encoding.TextMarshaler. Returns the 18-rune form. The zero ID
// errors rather than emitting an unreadable empty string; for optional fields use *ID.
func (id ID) MarshalText() ([]byte, error) {
	if id.string == "" {
		return nil, errors.New("sfid: cannot marshal zero ID; use *ID for optional fields")
	}
	return []byte(id.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler. Accepts a 15- or 18-rune forms.
func (id *ID) UnmarshalText(text []byte) error {
	parsed, ok := Parse(string(text))
	if !ok {
		return fmt.Errorf("sfid: cannot unmarshal %q into ID: invalid Salesforce ID", text)
	}
	*id = parsed
	return nil
}
