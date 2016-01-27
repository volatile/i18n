<p align="center"><img src="http://volatile.whitedevops.com/images/repositories/i18n/logo.png" alt="Volatile I18n" title="Volatile I18n"><br><br></p>

Volatile I18n is a helper for the [Core](https://github.com/volatile/core).  
It provides internationalization functions following the client preferences.

## Installation

```Shell
$ go get github.com/volatile/i18n
```

## Usage [![GoDoc](https://godoc.org/github.com/volatile/i18n?status.svg)](https://godoc.org/github.com/volatile/i18n)

```Go
package main

import (
	"github.com/volatile/core"
	"github.com/volatile/i18n"
	"github.com/volatile/response"
	"golang.org/x/text/language"
)

func main() {
	i18n.Init(locales, language.English)         // Default locale is language.English and client locale will be saved in a cookie.
	response.TemplatesFuncs(i18n.TemplatesFuncs) // Set functions for templates.

	i18n.Use(i18n.MatcherFormValue, i18n.MatcherAcceptLanguageHeader) // Try to match the client locale with the "locale" form value, or with his Accept-Language header, in this order.

	core.Use(func(c *core.Context) {
		response.Template(c, "hello", response.DataMap{
			"name":        "John Doe",
			"coinsNumber": 500,
		})
	})

	core.Run()
}

var locales = i18n.Locales{
	language.English: {
		"decimalMark":   ".",
		"thousandsMark": ",",

		"hello":      "Hello %s,",
		"how":        "How are you?",
		"coinsZero":  "Your wallet is empty.",
		"coinsOne":   "You have a single and precious coin.",
		"coinsOther": "You have " + i18n.TnPlaceholder + " coins.",
	},
	language.French: {
		"decimalMark":   ",",
		"thousandsMark": " ",

		"hello":      "Bonjour %s,",
		"how":        "Comment allez-vous?",
		"coinsZero":  "Vous êtes fauché.",
		"coinsOne":   "Vous avez une seule et précieuse pièce.",
		"coinsOther": "Vous possédez " + i18n.TnPlaceholder + " pièces.",
	},
}
```

In `templates/hello.gohtml`:

```HTML
{{define "hello"}}
	<!DOCTYPE html>
	<html>
		<head>
			<title>Hello</title>
		</head>
		<body>
			{{trans .c "hello" .name}}<br>     <!-- Hello John Doe,          -->
			{{trans .c "how"}}<br>             <!-- How are you?             -->
			{{transn .c "coins" .coinsNumber}} <!-- You have a 50,000 coins. -->
		</body>
	</html>
{{end}}
```

### Match locale

To match the client preferences, you need to set a handler with [`Use`](https://godoc.org/github.com/volatile/i18n#Use) and provide at least one matching function.

These ones are actually available:

- [`MatcherAcceptLanguageHeader`](https://godoc.org/github.com/volatile/i18n#MatcherAcceptLanguageHeader) to match the `Accept-Language` header.
- [`MatcherFormValue`](https://godoc.org/github.com/volatile/i18n#MatcherFormValue) to match the `locale` form value.

#### Get locale

Use [`ClientLocale`](https://godoc.org/github.com/volatile/i18n#ClientLocale) to get the locale used for the client.

#### Set locale

After parsing a language tag, use [`SetClientLocale`](https://godoc.org/github.com/volatile/i18n#SetClientLocale) to manually set the locale used for the client.

### Translations

Use [`T`](https://godoc.org/github.com/volatile/i18n#T) to get the translation for the client matched locale.
If the translation value contains format verbs (like `%s` or `%d`), the variadic receives the content for them.

When the translation associated to key doesn't exist, an empty string is returned in production mode (otherwise, the key).

#### Pluralization

[`Tn`](https://godoc.org/github.com/volatile/i18n#Tn) works like [`T`](https://godoc.org/github.com/volatile/i18n#T) but it tries to find the best translation form, following a number of elements.

A pluralized translation has 3 forms: zero, one, other.  
They are defined at the end of the key: `myTranslationKeyZero`, `myTranslationKeyOne` and `myTranslationKeyOther`.  
If [`TnPlaceholder`](https://godoc.org/github.com/volatile/i18n#pkg-constants) is used in the translation, the number of elements will take this place.

Translation example:

```Go
"appleZero"   = "There are no apples."
"appleOne"   = "There is a single apple."
"appleOther" = "There are " + i18n.TnPlaceholder + " apples."
```

Function example:

```Go
i18n.Tn(c, "apple", 7)
```
…results in `There are 7 apples.`.

### Numbers

Use [`Fmtn`](https://godoc.org/github.com/volatile/i18n#Fmtn) to get a formatted number with decimal and thousands marks.
If set, the special `decimalMark` and `thousandsMark` keys will be used from the matched locale.

### Templates functions

[`TemplatesFuncs`](https://godoc.org/github.com/volatile/i18n#TemplatesFuncs) provides a map of all functions usable in templates.

Example with the [Response](https://github.com/volatile/response) package:

```Go
response.TemplatesFuncs(i18n.TemplatesFuncs)
```
