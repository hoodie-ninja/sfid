# sfid

Go package for dealing with Salesforce IDs.

No external dependencies. 100% test coverage.

## Usage

```go
id, ok := sfid.Parse(input)
if !ok {
	slog.Error("invalid Salesforce ID!")
	return
}
id.ID()     // 15-rune case-sensitive Salesforce ID
id.String() // 18-rune case-insensitive Salesforce ID
```

IDs are comparable with `==` and safe for use as map keys. `ID` also implements:

- `driver.Valuer` / `sql.Scanner` — use IDs directly with `database/sql` and `pgx`/`pgxpool`
- `encoding.TextMarshaler` / `encoding.TextUnmarshaler` — use IDs directly with `encoding/json` and other text-based encoders

A zero `ID` cannot be written: `Value()` and `MarshalText()` return an error rather than
emitting an unreadable empty string. Use `*sfid.ID` for nullable columns and optional fields.
