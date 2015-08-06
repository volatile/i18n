package i18n

import (
	"errors"
	"net/http"
	"time"

	"github.com/volatile/core"
	"github.com/volatile/core/httputil"
	"github.com/volatile/core/log"
)

// GetLocale returns the current locale for the client and sets it if it's not already done.
func GetLocale(c *core.Context) string {
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
	SetLocale(c, l)
	return l
}

// SetLocale changes the locale for the actual client, but only if the locale exists.
func SetLocale(c *core.Context, l string) {
	if localeExists(l) {
		if useCookie {
			http.SetCookie(c.ResponseWriter, &http.Cookie{
				Name:    cookieName,
				Value:   l,
				Expires: time.Now().Add(3 * 365 * 24 * time.Hour), // 3 years cookie
			})
		}
		c.Data[contextDataKey] = l
	} else {
		log.Stack(errors.New("locale " + l + " doesn't exists"))
	}
}

// matchLocale returns the most appropriate and available locale key for the client.
// Content Language Headers: https://tools.ietf.org/html/rfc3282
func matchLocale(r *http.Request) string {
	acceptedLangs := httputil.AcceptedLanguages(r)
	if acceptedLangs == nil {
		return defaultLocale
	}

	for _, lang := range acceptedLangs {
		if len(lang) >= 2 {
			prefix := lang[:2]
			if _, ok := locales[prefix]; ok {
				return string(prefix)
			}
		}
	}

	return defaultLocale
}

func localeExists(l string) bool {
	_, ok := locales[l]
	return ok
}
