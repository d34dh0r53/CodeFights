package isIPv4Address

import (
	"regexp"
)

func isIPv4Address(inputString string) bool {
	validatingIPv4regex := `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
	r := regexp.MustCompile(validatingIPv4regex)
	return r.MatchString(inputString)
}
