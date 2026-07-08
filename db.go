package sfid

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
)

var (
	_ driver.Valuer = ID{}
	_ sql.Scanner   = (*ID)(nil)
)

// Value implements driver.Valuer. Returns the 18-rune form. The zero ID errors; for nullable columns use *ID.
func (id ID) Value() (driver.Value, error) {
	if id.string == "" {
		return nil, errors.New("sfid: cannot store zero ID; use *ID for nullable columns")
	}
	return id.String(), nil
}

// Scan implements sql.Scanner. Accepts string or []byte holding a 15- or 18-rune Salesforce ID.
func (id *ID) Scan(src any) error {
	switch v := src.(type) {
	case string:
		return id.scan(v)
	case []byte:
		return id.scan(string(v))
	case nil:
		return errors.New("sfid: cannot scan NULL into ID; use *ID for nullable columns")
	default:
		return fmt.Errorf("sfid: cannot scan %T into ID", src)
	}
}

func (id *ID) scan(s string) error {
	parsed, ok := Parse(s)
	if !ok {
		return fmt.Errorf("sfid: cannot scan %q into ID: invalid Salesforce ID", s)
	}
	*id = parsed
	return nil
}
