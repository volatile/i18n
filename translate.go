package i18n

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/volatile/core"
)

// T returns the translation associated to key, for the client locale.
func T(c *core.Context, key string, a ...interface{}) string {
	return Tn(c, key, -1, a...)
}

// HT works like T but returns an HTML unescaped translation.
func HT(c *core.Context, key string, a ...interface{}) template.HTML {
	return HTn(c, key, -1, a...)
}

// Tn returns the translation associated to key, for the client locale.
// If the translation defines plural forms (zero, one, other), it uses the most apropriate.
// All TnPlaceholder in the translation are replaced with number n.
func Tn(c *core.Context, key string, n interface{}, a ...interface{}) (s string) {
	if trs, ok := locales[ClientLocale(c)]; ok {
		if n == 0 {
			if v, ok := trs[key+"Zero"]; ok {
				s = fmt.Sprintf(v, a...)
			}
		} else if n == 1 {
			if v, ok := trs[key+"One"]; ok {
				s = fmt.Sprintf(v, a...)
			}
		} else if v, ok := trs[key+"Other"]; ok {
			s = fmt.Sprintf(v, a...)
		} else if v, ok := trs[key]; ok {
			s = fmt.Sprintf(v, a...)
		}
	}

	s = strings.Replace(s, TnPlaceholder, Fmtn(c, n), -1)

	if !core.Production && s == "" {
		s = key
	}
	return
}

// HTn works like Tn but returns an HTML unescaped translation.
func HTn(c *core.Context, key string, n interface{}, a ...interface{}) template.HTML {
	return template.HTML(strings.Replace(Tn(c, key, n, a...), "\n", "<br>", -1))
}
