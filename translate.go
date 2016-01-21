package i18n

import (
	"fmt"
	"strings"

	"github.com/volatile/core"
)

// T returns the translation associated to key, for the client matched locale.
func T(c *core.Context, key string, a ...interface{}) string {
	return t(c, -1, key, a)
}

// Tn returns the translation associated to key, for the client matched locale.
// If the translation defines plural forms (zero, one, other), it uses the most apropriate.
// All TNPlaceholder in the translation are replaced with n.
func Tn(c *core.Context, key string, n int, a ...interface{}) string {
	return strings.Replace(t(c, n, key, a), TnPlaceholder, Fmtn(c, n), -1)
}

func t(c *core.Context, count int, key string, a []interface{}) string {
	if trs, ok := locales[ClientLocale(c)]; ok {
		if count == 0 {
			if v, ok := trs[key+"Zero"]; ok {
				return fmt.Sprintf(v, a...)
			}
		}

		if count == 1 {
			if v, ok := trs[key+"One"]; ok {
				return fmt.Sprintf(v, a...)
			}
		}

		if v, ok := trs[key+"Other"]; ok {
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
