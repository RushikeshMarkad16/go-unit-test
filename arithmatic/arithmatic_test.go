package arithmatic

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {

	//Individual case handled
	out := div(8, 4)

	if out != 2 {
		fmt.Errorf("Failed to divide 8 by 4")
	}

	out = div(0, 4)

	if out != 2 {
		fmt.Errorf("Failed to divide 0 by 4")
	}

	out = div(8, 0)

	if out != 2 {
		fmt.Errorf("Failed to divide 8 by 0")
	}
}

// To run similar test cases with the for loop
func TestTDD(t *testing.T) {
	//anonymous struct
	tests := []struct {
		name     string
		x        float64
		y        float64
		expected float64
	}{
		{name: "Divide 8 by 4", x: 8, y: 4, expected: 2},
		{name: "Divide 0 by 4", x: 0, y: 4, expected: 0},
		{name: "Divide 8 by 0", x: 8, y: 0, expected: 0},
	}

	for _, tt := range tests {
		//To run the subtests
		t.Run(tt.name, func(t *testing.T) {
			res := div(tt.x, tt.y)
			if res != tt.expected {
				t.Errorf("Failed to divide %f by %f, Got : %f, Expected: %f", tt.x, tt.y, res, tt.expected)
			}
		})
	}
}
