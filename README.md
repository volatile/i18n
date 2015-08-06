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
)

func main() {
	i18n.Use(locales, "en", true)        // Default locale local is "en" and client local will be saved in a cookie on first match.
	response.ViewsFuncs(i18n.ViewsFuncs) // Functions for views templates

	core.Use(func(c *core.Context) {
		response.View(c, "hello", map[string]interface{}{
			"name":        "John Doe",
			"coinsNumber": 500,
		})
	})

	core.Run()
}

var locales = i18n.Locales{
	"en": i18n.Locale{
		"decimalMark":   ".",
		"thousandsMark": ",",

		"hello":       "Hello %s,",
		"how":         "How are you?",
		"coins.zero":  "Your wallet is empty.",
		"coins.one":   "You have a single and precious coin.",
		"coins.other": "You have " + i18n.TransNPlaceholder + " coins.",
	},
	"fr": i18n.Locale{
		"decimalMark":   ",",
		"thousandsMark": " ",

		"hello":       "Bonjour %s,",
		"how":         "Comment allez-vous?",
		"coins.zero":  "Vous êtes fauché.",
		"coins.one":   "Vous avez une seule et précieuse pièce.",
		"coins.other": "Vous possédez " + i18n.TransNPlaceholder + " pièces.",
	},
}

```

In `views/hello.gohtml`:

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

### Translations

Use `i18n.Trans(*core.Context, string, ...interface{})` to get the translation for the client matched locale.  
The string parameter is the translation key.  
If the translation value contains format verbs (like `%s` or `%d`), the last variadic receives the content for them.

When the translation associated to key doesn't exist, an empty string is returned en production mode. Otherwise, the key is returned.

#### Pluralization
`i18n.TransN(*core.Context, string, int, ...interface{})` works like `Trans` but it tries to find the best form following a number of elements.  
The third parameter is the number of elements.

The translation can have 3 forms: zero, one, other. They are defined at the end of the key: `myTranslationKey.zero`.  
If the TransNPlaceholder is used in the translation, the number of elements will take this place.

Translation example:
```Go
"apple.zero" = "There are no apples."
"apple.one" = "There is a single apple."
"apple.other" = "There are " + i18n.TransNPlaceholder + " apples."
```

Function example:
```Go
i18n.TransN(c, "apple", 7)
```
Results in `There are 7 apples.`

### Numbers

Use `Num(*core.Context, interface{})` the get a formatted number with decimal and thousands marks.
The special `decimalMark` and `thousandsMark` fields will be used from the matched locale.

### Views functions

To use internationalization in templates, `i18n.ViewsFuncs` provides a map of all functions.

Example with the [Response](https://github.com/volatile/response) package:

```Go
response.ViewsFuncs(i18n.ViewsFuncs)
```

In your templates, you can now use these functions:

- `trans` for `Trans`
- `transn` for `TransN`
- `num` for `Num`
