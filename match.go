package i18n

import (
	"github.com/volatile/core"
	"golang.org/x/text/language"
)

// Matcher is a matching function used by the handler.
type Matcher func(*core.Context) (language.Tag, language.Confidence)

// Match matches the first of the given tags to reach a certain confidence threshold with an available locale.
// The tags should therefore be specified in order of preference.
// Extensions are ignored for matching.
func Match(tt ...language.Tag) (t language.Tag, c language.Confidence) {
	t, _, c = matcher.Match(tt...)
	return
}

// MatchString parses string s and matches the first of the given tags to reach a certain confidence threshold with an available locale.
// The string can be a single language tag or a list of language tags with preference values (from the Accept-Language header, for example).
func MatchString(s string) (language.Tag, language.Confidence) {
	pref, _, _ := language.ParseAcceptLanguage(s)
	return Match(pref...)
}

// MatcherAcceptLanguageHeader matches the Accept-Language header.
func MatcherAcceptLanguageHeader(c *core.Context) (language.Tag, language.Confidence) {
	return MatchString(c.Request.Header.Get("Accept-Language"))
}

// MatcherFormValue matches the "locale" form value.
func MatcherFormValue(c *core.Context) (language.Tag, language.Confidence) {
	return MatchString(c.Request.FormValue("locale"))
}

// TODO: Matchers for location, GeoIP, subdomain, TLD, request path.
