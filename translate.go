package i18n

import (
	"fmt"
	"strings"

	"github.com/volatile/core"
)

// T returns the translation associated to key, for the client locale.
func T(c *core.Context, key string, a ...interface{}) string {
	return tn(c, key, -1, a)
}

// Tn returns the translation associated to key, for the client locale.
// If the translation defines plural forms (zero, one, other), it uses the most apropriate.
// All TnPlaceholder in the translation are replaced with number n.
func Tn(c *core.Context, key string, n interface{}, a ...interface{}) string {
	return strings.Replace(tn(c, key, n, a...), TnPlaceholder, Fmtn(c, n), -1)
}

func tn(c *core.Context, key string, n interface{}, a ...interface{}) string {
	if trs, ok := locales[ClientLocale(c)]; ok {
		if n == 0 {
			if v, ok := trs[key+"Zero"]; ok {
				return fmt.Sprintf(v, a...)
			}
		}

		if n == 1 {
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
