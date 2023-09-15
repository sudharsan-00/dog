package dog

import (
	"strings"
)

func Whengrownup(s string) string {
	return "when the puppy grows up it says: " + strings.ToUpper(s)
}