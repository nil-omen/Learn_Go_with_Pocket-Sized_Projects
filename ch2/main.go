package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang,
		"lang",
		"fr",
		"The required language, e.g. en, ur, ar ...")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// language represents the language's code
type language string

var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"ar": "مرحبًا يا عالم",    // Arabic
	"ur": "ہیلو دنیا",         // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

// greet says hello to the world in various languages
func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
