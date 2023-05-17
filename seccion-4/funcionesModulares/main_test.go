package main

import (
	"testing"
)

type testNumbers struct {
	numberOne int
	numberTwo int
	expected  bool
}

func TestPerfectNumber(t *testing.T) {
	var numbersTest = []testNumbers{
		{numberOne: 6, expected: true},
		{numberOne: 13, expected: false},
		{numberOne: 28, expected: true},
		{numberOne: 0, expected: false},
		{numberOne: 1, expected: false},
	}
	for index, number := range numbersTest {
		result := PerfectNumber(number.numberOne)
		if result != number.expected {
			t.Errorf("Test %d: Value of result(%v) not match with expected(%v)", index, result, number.expected)
		}
	}
}

func TestFriendNumbers(t *testing.T) {
	var numbersTest = []testNumbers{
		{numberOne: 220, numberTwo: 284, expected: true},
		{numberOne: 6, numberTwo: 18, expected: false},
		{numberOne: 1184, numberTwo: 1210, expected: true},
	}
	for index, number := range numbersTest {
		result := FriendNumbers(number.numberOne, number.numberTwo)
		if result != number.expected {
			t.Errorf("Test %d: Value of result(%v) not match with expected(%v)", index, result, number.expected)
		}
	}
}
