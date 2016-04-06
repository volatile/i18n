package i18n

import (
	"errors"
	"fmt"

	"github.com/volatile/core"

	"golang.org/x/text/language"
)

const contextDataKey = "i18nLocale"

// TnPlaceholder is the placeholder replaced by n in a translation, when using the TransN function.
const TnPlaceholder = "{{.n}}"

var (
	locales       Locales
	defaultLocale language.Tag
	matcher       language.Matcher
)

// Errors
var (
	ErrUnknownLocale = errors.New("i18n: unknown locale")
)

// TemplatesFuncs provides i18n functions that can be set for templates.
var TemplatesFuncs = map[string]interface{}{
	"clientLocale": ClientLocale,
	"fmtn":         Fmtn,
	"t":            T,
	"ht":           HT,
	"tn":           Tn,
	"htn":          HTn,
}

// Translations is a map of translations.
type Translations map[string]string

// Locales is a map of translations associated to language tags.
type Locales map[language.Tag]Translations

// Has checks if locale t exists in the locales map.
func (ll *Locales) Has(t language.Tag) bool {
	_, ok := (*ll)[t]
	return ok
}

// Init registers locales ll and default locale def for the entire app.
func Init(ll Locales, def language.Tag) {
	if locales != nil {
		panic(errors.New("i18n: Init called multiple times"))
	}

	locales = ll
	defaultLocale = def

	if !locales.Has(def) {
		panic(fmt.Errorf("i18n: default locale %q doesn't exist", def))
	}

	tt := []language.Tag{def}
	for t := range ll {
		if t != def {
			tt = append(tt, t)
		}
	}
	matcher = language.NewMatcher(tt)
}

// Use adds the handler to the default handlers stack.
// It matches locale for client, thanks to matchers.
// Multiple matchers can be used.
// The client locale is set as soon as a matcher is confident.
func Use(matchers ...Matcher) {
	core.Use(func(c *core.Context) {
		for _, m := range matchers {
			if t, conf := m(c); conf != language.No {
				if err := SetClientLocale(c, t); err == nil {
					break
				}
			}
		}
		c.Next()
	})
}

// CleanAcceptLanguage parses, cleans and returns the contents of a Accept-Language header.
// If an error is encountered, the returned string is the same as given.
func CleanAcceptLanguage(s string) (string, error) {
	tt, q, err := language.ParseAcceptLanguage(s)
	if err != nil {
		return s, err
	}

	s = ""
	for i := 0; i < len(tt); i++ {
		if i > 0 {
			s += ","
		}
		s += fmt.Sprintf("%s;q=%g", tt[i].String(), q[i])
	}
	return s, nil
}
