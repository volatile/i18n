package i18n

import (
	"errors"
	"net/http"

	"github.com/volatile/core"
)

// ErrUnknownLocale is returned when the wanted locale doesn't exist.
var ErrUnknownLocale = errors.New("i18n: unknown locale")

// ClientLocale returns the current locale used by the client.
// If the locale has not been matched already, it will be done before returning.
func ClientLocale(c *core.Context) string {
	// Use context data to match locale a single time per request.
	if v, ok := c.Data[contextDataKey]; ok && localeExists(v.(string)) {
		return v.(string)
	}

	// Use cookie if exists and valid.
	if useCookie {
		if cookie, err := c.Request.Cookie(cookieName); err == nil && localeExists(cookie.Value) {
			return cookie.Value
		}
	}

	// Match, save and return locale key.
	return SetClientLocale(c, c.Request.Header.Get("Accept-Language"))
}

// SetClientLocale changes the locale for the actual client.
// l can either be an available locale or a Accept-Language header as defined in RFC 2616 and RFC 3282.
// If the locale can't be matched, the default locale will be set.
// The matched locale is finally returned.
func SetClientLocale(c *core.Context, l string) string {
	l = MatchLocale(l)

	if useCookie {
		http.SetCookie(c.ResponseWriter, &http.Cookie{
			Name:   cookieName,
			Value:  l,
			Path:   "/",
			MaxAge: 315569260, // 10 years cookie
		})
	}

	c.Data[contextDataKey] = l
	return l
}
