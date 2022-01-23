package main

import (
	"fmt"
	"strings"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	lowerCasedLang := strings.ToLower(language)
	prefix := getLangPrefix(lowerCasedLang)
	return prefix + name
}

func getLangPrefix(language string) (prefix string) {
	switch language {
	case "english":
		prefix = englishHelloPrefix
		break
	case "spanish":
		prefix = spanishHelloPrefix
		break
	case "french":
		prefix = frenchHelloPrefix
		break
	default:
		prefix = englishHelloPrefix
		break
	}

	return
}

func main() {
	fmt.Println(Hello("", "english"))
}
