package slices

import "testing"

func TestRotaresWorks(t *testing.T) {
	type Test struct {
		input                string
		expectedRotatedRight string
		expectedRotatedLeft  string
	}

	var tests []Test = []Test{
		{
			input:                "",
			expectedRotatedRight: "",
			expectedRotatedLeft:  "",
		},
		{
			input:                "AA",
			expectedRotatedRight: "AA",
			expectedRotatedLeft:  "AA",
		},
		{
			input:                "ABC",
			expectedRotatedRight: "CAB",
			expectedRotatedLeft:  "BCA",
		},
		{
			input:                "MNO",
			expectedRotatedRight: "OMN",
			expectedRotatedLeft:  "NOM",
		},
		{
			input:                "xyza",
			expectedRotatedRight: "axyz",
			expectedRotatedLeft:  "yzax",
		},
	}

	for _, test := range tests {
		generatedRotatedRight := RotateRightMyVersion(test.input)
		if generatedRotatedRight != test.expectedRotatedRight {
			t.Errorf("RotateChainRight: No son iguales, generated: %s, expected: %s", generatedRotatedRight, test.expectedRotatedRight)
		}
		generatedRotatedLeft := RotateLeftMyVersion(test.input)
		if generatedRotatedLeft != test.expectedRotatedLeft {
			t.Errorf("RotateChainLeft: No son iguales, generated: %s, expected: %s", generatedRotatedRight, test.expectedRotatedRight)
		}
	}
}

func BenchmarkRotateRightMyVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RotateRightMyVersion("ABCDEFGH")
	}
}

func BenchmarkRotateRightVituVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RotateRightVituVersion("ABCDEFGH")
	}
}
