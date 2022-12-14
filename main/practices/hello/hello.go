package main

import "fmt"

const englishGreeting = "Hello, "
const frenchGreeting = "Bonjour, "
const spanishGreeting = "Hola, "
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name;
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = frenchGreeting
	case "spanish":
		prefix = spanishGreeting
	default:
		prefix = englishGreeting
	}
	return
}

func main() {
	fmt.Println(Hello("Alan",""))
}
