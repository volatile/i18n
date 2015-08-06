package i18n

import (
	"regexp"
	"sort"
)

// Locale contains translations and special i18n values associated to unique keys.
type Locale map[string]string

// Locales contains locales associated to their names.
type Locales map[string]Locale

const (
	// TransNPlaceholder is the placeholder replaced by n in a translation, when using the TransN function.
	TransNPlaceholder = "{{.n}}"

	contextDataKey = "i18nLocale"
)

var (
	useCookie  bool       // Defines if the matched locale must be saved in a cookie after matching.
	cookieName = "locale" // The name of the cookie used to save the client matched locale.

	defaultLocale   string
	locales         = make(Locales)
	localeKeyRegexp = regexp.MustCompile("^[a-z]{2}$")
)

// ViewsFuncs adds internationalization functions to views.
var ViewsFuncs = map[string]interface{}{
	"getLocale":        GetLocale,
	"num":              Num,
	"sortedLocaleKeys": SortedLocaleKeys,
	"trans":            Trans,
	"transn":           TransN,
}

// Use registers locales l.
// If none match to the client accepted languages, the def locale will be used.
// If cookie is true, a cookie will be used to save the most appropriate and available locale key for the client.
func Use(l Locales, def string, cookie bool) {
	if _, ok := l[def]; !ok {
		panic("i18n: default locale " + def + " doesn't exist")
	}

	for i, v := range l {
		if !localeKeyRegexp.MatchString(i) {
			panic(`i18n: locale key must be an ISO 639-1 code (2 letters lowercase), so "` + i + `" is invalid`)
		}
		locales[i] = v
	}
	defaultLocale = def

	useCookie = cookie
}

// SortedLocaleKeys return the sorted keys of all available locales.
func SortedLocaleKeys() (kk []string) {
	for k := range locales {
		kk = append(kk, k)
	}
	sort.Strings(kk)
	return
}
