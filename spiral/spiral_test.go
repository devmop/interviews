package spiral_test

import (
	"testing"

	"github.com/devmop/interviews/spiral"
)

func TestUnravelsSpirals(t *testing.T) {
	res := spiral.Unravel([][]uint{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})

	expected := []uint{1, 2, 3, 6, 9, 8, 7, 4, 5}

	if len(res) != len(expected) {
		t.Fatalf("The unravelled result '%v' should have as many values as the input", res)
	}

	matching := true
	for i := 0; i < len(res); i++ {
		if res[i] != expected[i] {
			matching = false
		}
	}

	if !matching {
		t.Errorf("the unravelled matrix doesn't match expected '%v' but got '%v'", expected, res)
	}
}
