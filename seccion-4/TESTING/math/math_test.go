package math

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	} else {
		t.Logf("Test Passed got %d, wanted %d", got, want)
	}
}

func TestAddV2(t *testing.T) {
	c := require.New(t)
	got := Add(4, 6)
	want := 10

	c.Equal(got, want)
}

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	{2, 3, 5},
	{4, 8, 12},
	{4, 1, 5},
	{3, 10, 13},
}

func TestAddV3(t *testing.T) {
	for index, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Prueba %v, Output %d not equal to expected %d", index, output, test.expected)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}

func ExampleAdd() {
	fmt.Println(Add(4, 6))
	// Output: 10
}
