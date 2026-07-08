package sfid

import "strings"

// ID represents a Salesforce Identifier. IDs can be directly compared for equality using == and are safe
// for use as map keys.
type ID struct{ string }

// Parse accepts 15 or 18-rune Salesforce IDs and returns an ID. Ignores leading and trailing
// whitespace. If the ID is invalid in any way, false is returned with a "zero" ID.
func Parse(s string) (ID, bool) {
	s = strings.TrimSpace(s)
	if len(s) != 15 && len(s) != 18 {
		return ID{}, false
	}
	for i := range len(s) {
		if !base62(s[i]) {
			return ID{}, false
		}
	}
	if len(s) == 15 {
		return ID{s}, true
	}
	// len(s) == 18: the suffix is authoritative for the body's casing (see applyMask).
	body, ok := applyMask(s)
	if !ok {
		return ID{}, false
	}
	return ID{body}, true
}

// String returns the case-insensitive 18-rune form.
func (id ID) String() string {
	if id.string == "" {
		return ""
	}
	var out [18]byte
	copy(out[:15], id.string)
	writeMask(out[15:], id.string)
	return string(out[:])
}

// ID() returns the case-sensitive 15-rune form.
func (id ID) ID() string { return id.string }

// base62 reports whether c is in the Base-62 alphabet (0-9, A-Z, a-z).
func base62(c byte) bool {
	return '0' <= c && c <= '9' || 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z'
}

// suffixIndex returns the position of c in caseRunes, matching letters case-insensitively,
// or -1 if c is not a valid checksum character (A-Z, 0-5).
func suffixIndex(c byte) int {
	switch {
	case 'A' <= c && c <= 'Z':
		return int(c - 'A')
	case 'a' <= c && c <= 'z':
		return int(c - 'a')
	case '0' <= c && c <= '5':
		return int(c-'0') + 26
	default:
		return -1
	}
}

// writeMask writes the 3-rune capitalization checksum for the 15-rune body s into dst.
func writeMask(dst []byte, s string) {
	for i := range 3 {
		bits := 0
		for j := range 5 {
			if c := s[i*5+j]; 'A' <= c && c <= 'Z' {
				bits |= 1 << j
			}
		}
		dst[i] = caseRunes[bits]
	}
}

// applyMask decodes the case-sensitive 15-rune body of a case-insensitive 18-rune ID.
// Each suffix rune encodes the capitalization of one 5-rune chunk of the body.
// Returns false if a suffix rune is outside the checksum alphabet (A-Z, 0-5).
func applyMask(s string) (string, bool) {
	var body [15]byte
	for i := range 3 {
		bits := suffixIndex(s[15+i])
		if bits < 0 {
			return "", false
		}
		for j := range 5 {
			c := s[i*5+j]
			// only apply case shift for non-digits
			if c >= 'A' {
				if bits>>j&1 == 1 {
					c &^= 0x20
				} else {
					c |= 0x20
				}
			}
			body[i*5+j] = c
		}
	}
	return string(body[:]), true
}
