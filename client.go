package i18n

import (
	"net/http"

	"golang.org/x/text/language"

	"github.com/volatile/core"
)

// ClientLocale returns the current locale used by the client.
// If the locale has not been matched already, it will be done before returning.
func ClientLocale(c *core.Context) language.Tag {
	// Use context data to match locale a single time per request.
	if v, ok := c.Data[contextDataKey]; ok && locales.Has(v.(language.Tag)) {
		return v.(language.Tag)
	}

	if useCookie {
		if cookie, err := c.Request.Cookie(cookieName); err == nil {
			if t, err := language.Parse(cookie.Value); err == nil && locales.Has(t) {
				return t
			}
		}
	}

	pref, _, _ := language.ParseAcceptLanguage(c.Request.Header.Get("Accept-Language"))
	t, _, _ := matcher.Match(pref...)
	SetClientLocale(c, t)
	return t
}

// SetClientLocale changes the locale for the actual client.
// If the language tag t doesn't match any available locale, error ErrUnknownLocale is returned.
func SetClientLocale(c *core.Context, t language.Tag) error {
	if !locales.Has(t) {
		return ErrUnknownLocale
	}

	if useCookie {
		http.SetCookie(c.ResponseWriter, &http.Cookie{
			Name:   cookieName,
			Value:  t.String(),
			Path:   "/",
			MaxAge: 315569260, // 10 years
		})
	}

	c.Data[contextDataKey] = t
	return nil
}
