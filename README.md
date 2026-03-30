# sfid

Go package for dealing with Salesforce IDs.

No external dependencies. 100% test coverage.

## Usage

```go

id,ok := sfid.Parse(input)
if !ok {
	slog.Log("invalid Salesforce ID!")
	return
}
id.String() // 15 rune Salesforce ID
id.CaseSafe() // 18 rune Salesforce ID
