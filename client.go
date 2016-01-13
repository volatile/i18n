package i18n

import (
	"errors"
	"net/http"

	"golang.org/x/text/language"

	"github.com/volatile/core"
)

// ClientLocale returns the current locale used by the client.
// If the locale has not been matched already, it will be done before returning.
func ClientLocale(c *core.Context) string {
	// Use context data to match locale a single time per request.
	if v, ok := c.Data[contextDataKey]; ok {
		return v.(string)
	}

	// Use cookie if exists and valid.
	if useCookie {
		if cookie, err := c.Request.Cookie(cookieName); err == nil && localeExists(cookie.Value) {
			return cookie.Value
		}
	}

	// Match, save and return locale key.
	l := matchLocale(c.Request)
	SetClientLocale(c, l)
	return l
}

// SetClientLocale changes the locale for the actual client, but only if the locale exists.
func SetClientLocale(c *core.Context, l string) {
	if localeExists(l) {
		if useCookie {
			http.SetCookie(c.ResponseWriter, &http.Cookie{
				Name:   cookieName,
				Value:  l,
				Path:   "/",
				MaxAge: 315569260, // 10 years cookie
			})
		}
		c.Data[contextDataKey] = l
	} else {
		panic(errors.New("i18n: locale " + l + " doesn't exists"))
	}
}

// matchLocale returns the most appropriate and available locale key for the client.
// Content Language Headers: https://tools.ietf.org/html/rfc3282
func matchLocale(r *http.Request) string {
	tags, _, err := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	if err != nil {
		panic(err)
	}

	for _, tag := range tags {
		base, _ := tag.Base()
		if _, ok := (*locales)[base.String()]; ok {
			return base.String()
		}
	}

	return defaultLocale
}

func localeExists(l string) bool {
	_, ok := (*locales)[l]
	return ok
}
