package main

import (
	"fmt"
	"strings"
)

// a: 1 moneda e: 1 moneda i: 2 monedas o: 3 monedas u: 4 monedas

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaronu", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	for _, user := range users {
		coin := vowelCount(user, &coins)
		distribution[user] = coin
	}
	fmt.Println(distribution)

	fmt.Println("Coins left:", coins)
}

func vowelCount(userName string, coins *int) int {
	count := 0
	mapVowelValue := map[string]int{"a": 1, "e": 1, "i": 2, "o": 3, "u": 4}
	for _, character := range userName {
		char := strings.ToLower(string(character))
		if valor, ok := mapVowelValue[char]; ok {
			count += valor
		}
		if count > 10 {
			count = 10
		}
		if count > *coins {
			count -= *coins
		}
	}
	*coins -= count

	return count
}
