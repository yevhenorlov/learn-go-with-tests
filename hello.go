package main

import (
	"fmt"
)

func GetPrefix(locale string) (prefix string) {
	const enHelloPrefix = "Hello, "
	const uaHelloPrefix = "Привіт, "
	const nlHelloPrefix = "Hoi, "
	switch locale {
	case "ua":
		prefix = uaHelloPrefix
	case "nl":
		prefix = nlHelloPrefix
	case "en":
	default:
		prefix = enHelloPrefix
	}
	return
}

func Hello(name, locale string) string {
	if name == "" {
		name = "World"
	}
	return GetPrefix(locale) + name
}

func main() {
	fmt.Println(Hello("World", "en"))
}
