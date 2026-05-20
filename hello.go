package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"

	prefix := englishHelloPrefix
	switch language {
		case spanish:
		  prefix = spanishHelloPrefixz
		case frecnh:
		  prefix = frenchHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
