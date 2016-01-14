package i18n

import (
	"fmt"
	"strings"

	"github.com/volatile/core"
)

// Trans returns the translation associated to key, for the client matched locale.
func Trans(c *core.Context, key string, a ...interface{}) string {
	return trans(c, -1, key, a)
}

// TransN returns the translation associated to key, for the client matched locale.
// If the translation defines plural forms (zero, one, other), it uses the most apropriate.
// all TransNPlaceholder in the translation are replaced with n.
func TransN(c *core.Context, key string, n int, a ...interface{}) string {
	return strings.Replace(trans(c, n, key, a), TransNPlaceholder, Num(c, n), -1)
}

func trans(c *core.Context, count int, key string, a []interface{}) string {
	if trs, ok := locales[ClientLocale(c)]; ok {
		if count == 0 {
			if v, ok := trs[key+".zero"]; ok {
				return fmt.Sprintf(v, a...)
			}
		}

		if count == 1 {
			if v, ok := trs[key+".one"]; ok {
				return fmt.Sprintf(v, a...)
			}
		}

		if v, ok := trs[key+".other"]; ok {
			return fmt.Sprintf(v, a...)
		}

		if v, ok := trs[key]; ok {
			return fmt.Sprintf(v, a...)
		}
	}

	if core.Production {
		return ""
	}
	return key
}
