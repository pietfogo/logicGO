package main

import (
	"fmt"
	"strings"
)

func countCharOccurrences(text string, char byte) int {
	count := 0

	for _, c := range text {
		if byte(c) == char {
			count++
		}
	}

	return count
}

func main() {
	var text string
	var char byte

	fmt.Println("Digite uma string:")
	fmt.Scanln(&text)

	fmt.Println("Digite o caractere a ser contado:")
	fmt.Scanf("%c", &char)

	count := countCharOccurrences(strings.ToLower(text), char)

	fmt.Printf("O caractere '%c' aparece %d vezes na string.\n", char, count)
}
