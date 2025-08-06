package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	formattedText := strings.Fields(strings.ToLower(text))
	return formattedText
}
