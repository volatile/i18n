package i18n

import (
	"errors"
	"fmt"

	"golang.org/x/text/language"
)

const (
	contextDataKey = "i18nLocale"
	// TnPlaceholder is the placeholder replaced by n in a translation, when using the TransN function.
	TnPlaceholder = "{{.n}}"
)

// Errors
var (
	ErrUnknownLocale = errors.New("i18n: unknown locale")
)

var (
	locales    Locales
	matcher    language.Matcher
	useCookie  bool
	cookieName = "locale"
)

// TemplatesFuncs provides i18n functions that can be set for templates.
var TemplatesFuncs = map[string]interface{}{
	"clientLocale": ClientLocale,
	"fmtn":         Fmtn,
	"t":            T,
	"tn":           Tn,
}

// Translations is a map of translations for a language tag.
type Translations map[string]string

// Locales contains language tags and their translations.
type Locales map[language.Tag]Translations

// Has checks if the locale tag t exists in the locales map.
func (ll *Locales) Has(t language.Tag) bool {
	_, ok := (*ll)[t]
	return ok
}

// Use registers locales ll.
// On request, if none matches the client accepted languages, the locale def will be used.
// If cookie is true, a cookie will be used to save the most appropriate and available locale tag for the client.
func Use(ll Locales, def language.Tag, cookie bool) {
	locales = ll
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

	useCookie = cookie
}

// CleanAcceptLanguage parses, cleans and returns the contents of a Accept-Language header.
// If an error is encountered, the result is an empty string.
func CleanAcceptLanguage(s string) (string, error) {
	tt, q, err := language.ParseAcceptLanguage(s)
	if err != nil {
		return "", err
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
