package main

import "fmt"

func SumDivisores(number int) int {
	sum := 0
	for i := 1; i < number; i++ {
		if number%i == 0 {
			sum += i
		}
	}
	return sum
}

func PerfectNumber(number int) bool {
	sumDivisores := SumDivisores(number)
	return number == sumDivisores && number != 0
}

func FriendNumbers(numberOne, numberTwo int) bool {
	sumDivisoresNumberOne := SumDivisores(numberOne)
	sumDivisoresNumberTwo := SumDivisores(numberTwo)
	return sumDivisoresNumberOne == numberTwo && sumDivisoresNumberTwo == numberOne
}

func main() {
	var number = 220
	isPerfectNumber := PerfectNumber(number)
	if isPerfectNumber {
		fmt.Printf("%d es un número perfecto\n", number)
	} else {
		fmt.Printf("%d no es un número perfecto\n", number)
	}
	var numberTwo = 284
	areFriendNumbers := FriendNumbers(number, numberTwo)
	if areFriendNumbers {
		fmt.Printf("%d y %d son números amigos\n", number, numberTwo)
	} else {
		fmt.Printf("%d y %d no son números amigos\n", number, numberTwo)
	}

}
