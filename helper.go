package i18n

import "regexp"

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
	locales         *Locales // The map of available locales. It is ensured to not be empty.
	localeKeyRegexp = regexp.MustCompile("^[a-z]{2}$")
)

// ViewsFuncs provides i18n functions that can be set for templates.
var ViewsFuncs = map[string]interface{}{
	"locale":           ClientLocale,
	"num":              Num,
	"sortedLocaleKeys": SortedLocaleKeys,
	"trans":            Trans,
	"transn":           TransN,
}

// Use registers locales l.
// On request, if none match to the client accepted languages, the locale def will be used.
// And if useCookie is true, a cookie will be used to save the most appropriate and available locale key for the client.
func Use(l *Locales, def string, cookie bool) {
	if l == nil {
		panic("i18n: locales map can't be empty")
	}

	for i := range *l {
		if !localeKeyRegexp.MatchString(i) {
			panic(`i18n: locale key must be an ISO 639-1 code (2 letters lowercase) so "` + i + `" is invalid`)
		}
	}
	locales = l

	if !localeExists(def) {
		panic("i18n: default locale " + def + " doesn't exist")
	}
	defaultLocale = def

	useCookie = cookie
}
