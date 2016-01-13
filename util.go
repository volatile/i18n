package i18n

import (
	"fmt"
	"sort"

	"golang.org/x/text/language"
)

func localeExists(l string) bool {
	_, ok := (*locales)[l]
	return ok
}

// SortedLocaleKeys returns the sorted keys of all available locales.
func SortedLocaleKeys() (kk []string) {
	for k := range *locales {
		kk = append(kk, k)
	}
	sort.Strings(kk)
	return
}

// MatchLocale returns the most appropriate and available locale for the client.
func MatchLocale(l string) string {
	if localeExists(l) {
		return l
	}

	tag, _, _ := language.ParseAcceptLanguage(l)

	for _, t := range tag {
		b, _ := t.Base()
		if _, ok := (*locales)[b.String()]; ok {
			return b.String()
		}
	}

	return defaultLocale
}

// CleanAcceptLanguage parses, cleans and returns the contents of a Accept-Language header.
func CleanAcceptLanguage(s string) (string, error) {
	tag, q, err := language.ParseAcceptLanguage(s)
	if err != nil {
		return "", err
	}

	s = ""
	for i := 0; i < len(tag); i++ {
		if i > 0 {
			s += ","
		}
		s += fmt.Sprintf("%s;q=%g", tag[i].String(), q[i])
	}
	return s, nil
}
