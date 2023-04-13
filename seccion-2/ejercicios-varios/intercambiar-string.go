package main

import "fmt"

func main() {
	word := "carlos"
	fmt.Println(intercambio(word))
}

func intercambio(word string) string {
	wordReverse := ""
	for _, character := range word {
		wordReverse = string(character) + wordReverse
	}
	return wordReverse
}
