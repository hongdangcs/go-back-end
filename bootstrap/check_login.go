package bootstrap

import (
	"strings"
)

func ContainsAtSymbol(input string) bool {
	return strings.Contains(input, "@")
}
