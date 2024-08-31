package main

import (
	"fmt"
)

const enHelloPrefix = "Hello, "
const uaHelloPrefix = "Привіт, "

func Hello(name, locale string) string {
	if name == "" {
		name = "World"
	}
	if locale == "ua" {
		return uaHelloPrefix + name
	}
	return enHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", "en"))
}
