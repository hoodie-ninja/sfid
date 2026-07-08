package sfid

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

var (
	_ driver.Valuer = ID{}
	_ sql.Scanner   = (*ID)(nil)
)

// Value implements driver.Valuer. Returns the 18-rune form.
func (id ID) Value() (driver.Value, error) { return id.String(), nil }

// Scan implements sql.Scanner. Accepts a string or []byte holding a 15- or 18-rune Salesforce ID;
// NULL is rejected (there is no nil ID). Scan into *sfid.ID to accept a NULL column.
func (id *ID) Scan(src any) error {
	switch v := src.(type) {
	case string:
		return scan(id, v)
	case []byte:
		return scan(id, string(v))
	case nil:
		return fmt.Errorf("sfid: cannot scan NULL into ID")
	default:
		return fmt.Errorf("sfid: cannot scan %T into ID", src)
	}
}

func scan(id *ID, s string) error {
	parsed, ok := Parse(s)
	if !ok {
		return fmt.Errorf("sfid: cannot scan %q into ID: invalid Salesforce ID", s)
	}
	*id = parsed
	return nil
}
