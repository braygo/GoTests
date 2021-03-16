package main

import (
	"fmt"
)

func Hello(name string) string {
	phrase := "Hello, " + name
	return phrase
}

func main() {
	fmt.Println(Hello("Chris"))
}
