package chetiTLS

import "strings"

// CanonicalDomain returns a lower case domain with trim space.
func CanonicalDomain(domain string) string {
	return strings.ToLower(strings.TrimSpace(domain))
}
