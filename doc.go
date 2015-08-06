/*
Package i18n is a helper for the Core.
It provides internationalization functions following the client preferences.

Usage

Full example:

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

In "views/hello.gohtml":

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

Translations

Use Trans to get the translation for the client matched locale.
If the translation value contains format verbs (like %s or %d), the last variadic receives the content for them.

When the translation associated to key doesn't exist, an empty string is returned in production mode (otherwise, the key).

Pluralization

TransN works like Trans but it tries to find the best translation form, following a number of elements.

A pluralized translation has 3 forms: zero, one, other.
They are defined at the end of the key: "myTranslationKey.zero", "myTranslationKey.one" and "myTranslationKey.other".
If TransNPlaceholder is used in the translation, the number of elements will take this place.

Translation example:

	"apple.zero" = "There are no apples."
	"apple.one" = "There is a single apple."
	"apple.other" = "There are " + i18n.TransNPlaceholder + " apples."

Function example:

	i18n.TransN(c, "apple", 7)

…results in "There are 7 apples".

Numbers

Use #Num to get a formatted number with decimal and thousands marks.
If set, the special "decimalMark" and "thousandsMark" keys will be used from the matched locale.

Views functions

ViewsFuncs provides a map of all functions usable in templates.

Example with the [Response](https://github.com/volatile/response) package:

	response.ViewsFuncs(i18n.ViewsFuncs)

Get current locale

Use GetLocale to the current locale for the client.

Set current locale

Use SetLocale to manually set the locale for the client.
*/
package i18n
