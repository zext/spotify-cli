package stringsx

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

func Join(s []string) string {
	return strings.Join(s, ", ")
}

func JoinAnd(s []string) string {
	l := len(s)
	if l < 2 {
		return Join(s)
	}

	return fmt.Sprintf("%s%s and %s", Join(s[:l-1]), lo.Ternary(l == 2, "", ","), s[l-1])
}
