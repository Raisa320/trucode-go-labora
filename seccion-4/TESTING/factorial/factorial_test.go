package factorial

import "testing"

type factorialTest struct {
	arg1, expected int
}

var factorialTests = []factorialTest{
	{0, 1},
	{3, 6},
	{6, 720},
	{7, 5040},
}

func TestFactorial(t *testing.T) {
	for index, factorial := range factorialTests {
		if output := Factorial(factorial.arg1); output != factorial.expected {
			t.Errorf("Prueba %v, Output %d not equal to expected %d", index, output, factorial.expected)
		}
	}
}

func BenchmarkFactorial(b *testing.B) {
	for _, factorial := range factorialTests {
		Factorial(factorial.arg1)
	}
}
