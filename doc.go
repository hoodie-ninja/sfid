/*
Package sfid provides a model of a Salesforce Identifier (ID) that enables...
- easy parsing from API requests
- easy serialization into API responses
- easy conversion between the 15 and 18 rune formats of Salesforce IDs
- easy case-insensitive comparison of IDs
*/
package sfid

// caseRunes is the indexed set of capitalization checksums for the 18-rune format.
const caseRunes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ012345"
