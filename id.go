package sfid

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

// ID represents a Salesforce Identifier. IDs can be directly compared for equality using == and are safe
// for use as map keys.
type ID struct{ string }

// Parse accepts a 15 or 18-rune Salesforce ID and returns an ID. Ignores leading and trailing whitespace.
// If the ID is invalid in any way, false is returned with a "zero" ID.
func Parse(s string) (ID, bool) {
	s = strings.TrimSpace(s)
	if len(s) != 15 && len(s) != 18 {
		return ID{}, false
	}
	if strings.ContainsFunc(s, func(r rune) bool { return !strings.ContainsRune(idRunes, r) }) {
		return ID{}, false
	}
	if len(s) == 15 {
		fmt.Println("case mask:", caseMask(s))
		return ID{s + caseMask(s)}, true
	}
	// len(s) == 18
	return ID{applyMask(s)}, true
}

// CaseSafe returns the case-insensitive 18-rune "default" format ID.
func (id ID) CaseSafe() string { return id.string }

// String implements fmt.Stringer. Returns CaseSafe().
func (id ID) String() string { return id.string }

// ID() returns the case-sensitive 15-rune "internal" format ID.
func (id ID) ID() string {
	if id.string == "" {
		return ""
	}
	return id.string[:15]
}

// helper to generate case mask.
// only called after clean() on 15-rune input
func caseMask(s string) string {
	// s is now 15 runes
	buf := new(bytes.Buffer)
	for i := range 3 {
		mask := 0b00000
		for j := 4; j >= 0; j-- {
			r := rune(s[i*5+j]) // +1 starts at the end, -j walks backwards, -1 adjusts position to index
			if unicode.IsUpper(r) {
				mask |= 0b1 << uint(j)
			}
		}
		buf.WriteByte(caseRunes[mask])
	}
	return buf.String()
}

// applyMask applies capitalization rules in mask to prior string.
// only called after clean() on 18-rune input
func applyMask(s string) string {
	mask := s[15:]
	idbuf := new(strings.Builder)
	for i := range 3 {
		chunk := []rune(s[i*5 : (i+1)*5])
		maskRune := unicode.ToUpper(rune(mask[i]))
		chunkMask := strings.IndexRune(caseRunes, maskRune)
		for j := 4; j >= 0; j-- {
			if chunkMask>>j&0b1 == 0b1 {
				chunk[j] = unicode.ToUpper(chunk[j])
			}
		}
		idbuf.WriteString(string(chunk))
	}
	return idbuf.String() + strings.ToUpper(string(mask))
}
