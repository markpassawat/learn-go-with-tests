package main

import "strings"

/*
Write a test
Make the compiler pass
Run the test, see that it fails and check the error message is meaningful
Write enough code to make the test pass
Refactor
*/
const (
	spanish = "spanish"
	thai    = "thai"

	spanishHelloPrefix = "Hola, "
	englishHelloPrefix = "Hello, "
	thaiHelloPrefix    = "สวัสดี, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch strings.ToLower(language) {
	case spanish:
		prefix = spanishHelloPrefix
	case thai:
		prefix = thaiHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
