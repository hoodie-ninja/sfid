/*
Package sfid provides a model of a Salesforce Identifier (ID) that enables the following:
- easy parsing from API requests, databases (sql.Scanner), and JSON
- easy serialization into API responses, databases (driver.Valuer), and JSON
- easy conversion between the 15 and 18 rune formats of Salesforce IDs
- easy case-insensitive comparison of IDs
- easy embedding into other structs
*/
package sfid

// caseRunes is the indexed set of capitalization checksums for the 18-rune format.
const caseRunes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ012345"

const (
	zeroBody   = "000000000000000"    // the zero ID's 15-rune form
	zeroString = "000000000000000AAA" // the zero ID's canonical 18-rune form
)
