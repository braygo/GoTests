package main

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
	case "Spanish":
		return spanishHelloPrefix + name
	case "French":
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	//fmt.Println(Hello("Chris"))
}
