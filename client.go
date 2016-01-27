package i18n

import (
	"golang.org/x/text/language"

	"github.com/volatile/core"
)

// ClientLocale returns the current locale used by the client.
func ClientLocale(c *core.Context) language.Tag {
	if v, ok := c.Data[contextDataKey]; ok {
		return v.(language.Tag)
	}
	return defaultLocale
}

// SetClientLocale changes the locale for the actual client.
// If the language tag t doesn't match any available locale, error ErrUnknownLocale is returned.
func SetClientLocale(c *core.Context, t language.Tag) error {
	if !locales.Has(t) {
		return ErrUnknownLocale
	}
	c.Data[contextDataKey] = t
	return nil
}
