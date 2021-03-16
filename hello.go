package main

import (
	"fmt"
)

func Hello() string {
	phrase := "Hello World"
	return phrase
}

func main() {
	fmt.Println(Hello())
}
