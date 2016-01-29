<p align="center"><img src="http://volatile.whitedevops.com/images/repositories/i18n/logo.png" alt="Volatile I18n" title="Volatile I18n"><br><br></p>

[![GoDoc](https://godoc.org/github.com/volatile/i18n?status.svg)](https://godoc.org/github.com/volatile/i18n)

Volatile I18n is a helper for the [Core](https://github.com/volatile/core).  
It provides internationalization functions following the client preferences.

## Set translations

A translation is associated to a key, which is associated to a language tag, which is part of [Locales](https://godoc.org/github.com/volatile/i18n#Locales) map.

All translations can be stored like this:

```Go
var locales = i18n.Locales{
	language.English: {
		"decimalMark":   ".",
		"thousandsMark": ",",

		"hello":         "Hello %s,",
		"how":           "How are you?",
		"basementZero":  "All the money hidden in your basement has been spent.",
		"basementOne":   "A single dollar remains in your basement.",
		"basementOther": "You have " + i18n.TnPlaceholder + " bucks in your basement.",
	},
	language.French: {
		"decimalMark":   ",",
		"thousandsMark": " ",

		"hello":         "Bonjour %s,",
		"how":           "Comment allez-vous?",
		"basementZero":  "Tout l'argent caché dans votre sous-sol a été dépensé.",
		"basementOne":   "Un seul dollar se trouve dans votre sous-sol.",
		"basementOther": "Vous avez " + i18n.TnPlaceholder + " briques dans votre sous-sol.",
	},
}
```

`decimalMark` and `thousandsMark` are special keys that defines the number decimal and thousands separators when using [Tn](https://godoc.org/github.com/volatile/i18n#Tn) or [Fmtn](https://godoc.org/github.com/volatile/i18n#Fmtn).

With these translations, you need to [Init](https://godoc.org/github.com/volatile/i18n#Init) this package (the second argument is the default locale):

```Go
i18n.Init(locales, language.English)
```

## Detect client locale

When a client makes a request, the best locale must be matched to his preferences.  
To achieve this, you need to [Use](https://godoc.org/github.com/volatile/i18n#Use) a handler with one or more matchers:

```Go
i18n.Use(i18n.MatcherFormValue, i18n.MatcherAcceptLanguageHeader)
```

The client locale is set as soon as a matcher is confident.

A matcher is a function that returns the locale parsed from [core.Context](https://godoc.org/github.com/volatile/core#Context) with its level of confidence.  
These ones are actually available: [MatcherAcceptLanguageHeader](https://godoc.org/github.com/volatile/i18n#MatcherAcceptLanguageHeader) and [MatcherFormValue](https://godoc.org/github.com/volatile/i18n#MatcherFormValue).

## Use translations

A translation can be accessed with [T](https://godoc.org/github.com/volatile/i18n#T), receiving the core.Context (which contains the matched locale), the translation key, and optional arguments (if the translation contains formatting verbs):

```Go
i18n.T(c, "hello", "Walter White")
i18n.T(c, "how")
```

If a translation has pluralized forms, you can use [Tn](https://godoc.org/github.com/volatile/i18n#Tn) and the most appropriate form will be used according to the quantity:

```Go
i18n.Tn(c, "basement", 333000.333)
```

will result in `You have 333,000.333 bucks in your basement.`.

If you use templates, [TemplatesFuncs](https://godoc.org/github.com/volatile/i18n#TemplatesFuncs) provides a map of all usable functions.  
Example for the [Response](https://github.com/volatile/response) package:

```Go
response.TemplatesFuncs(i18n.TemplatesFuncs)
```
