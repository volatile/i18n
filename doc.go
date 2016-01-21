/*
Package i18n is a helper for the Core (https://github.com/volatile/core).
It provides internationalization functions following the client preferences.

Installation

In the terminal:

	$ go get github.com/volatile/i18n

Usage

Example:

	package main

	import (
		"github.com/volatile/core"
		"github.com/volatile/i18n"
		"github.com/volatile/response"
		"golang.org/x/text/language"
	)

	func main() {
		i18n.Use(&locales, language.English, true) // Default locale is language.English and client locale will be saved in a cookie.
		response.TemplatesFuncs(i18n.TemplatesFuncs)       // Functions for templates

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

In "templates/hello.gohtml":

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

Locale detection

Only the standard Accept-Language header is used to detect the most appropriate locale for the client.

Get locale

Use ClientLocale to get the locale used for the client.

Set locale

If you don't want to use the Accept-Language header, you are free to use any client side strategy to get the a language tag:

● Form parameter: http://example.com/?locale=en

● Subdomain: http://en.example.com/

● URL path element: http://example.com/en/

After parsing a language tag, use SetClientLocale to manually set the locale used for the client.

Translations

Use T to get the translation for the client matched locale.
If the translation value contains format verbs (like %s or %d), the last variadic receives the content for them.

When the translation associated to key doesn't exist, an empty string is returned in production mode (otherwise, the key).

Pluralization

Tn works like T but it tries to find the best translation form, following a number of elements.

A pluralized translation has 3 forms: zero, one, other.
They are defined at the end of the key: "myTranslationKeyZero", "myTranslationKeyOne" and "myTranslationKeyOther".
If TnPlaceholder is used in the translation, the number of elements will take this place.

Translation example:

	"appleZero" = "There are no apples."
	"appleOne" = "There is a single apple."
	"appleOther" = "There are " + i18n.TnPlaceholder + " apples."

Function example:

	i18n.Tn(c, "apple", 7)

…results in "There are 7 apples".

Numbers

Use Fmtn to get a formatted number with decimal and thousands marks.
If set, the special "decimalMark" and "thousandsMark" keys will be used from the matched locale.

Templates functions

TemplatesFuncs provides a map of all functions usable in templates.

Example with the Response (https://github.com/volatile/response) package:

	response.TemplatesFuncs(i18n.TemplatesFuncs)
*/
package i18n
