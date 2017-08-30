package isIPv4Address

import (
	"fmt"
	"regexp"
  "strconv"
)

func isIPv4Address(inputString string) bool {
	fmt.Println(inputString)
	var validatingIPv4regex string
	validatingIPv4regex = strconv.Quote("\b((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4}\b")
	r, _ := regexp.Compile(validatingIPv4regex)
	isIpV4 := r.MatchString(inputString)
	return isIpV4
}
