package algorithms

import (
	"cmp"
	"testing"
)

type sortTest[T cmp.Ordered] struct {
	Name     string
	Vals     []T
	Expected []T
}

var testElements = []sortTest{
	{
		Name:     "Simple",
		Vals:     []int{1, 5, 3},
		Expected: []any{1, 3, 5},
	},
	{
		Name:     "Hard",
		Vals:     []any{1, 5, 3, 4623, 20, 0, 1, 1, 1, 32, 0, 4623, 2623, 4624},
		Expected: []any{0, 0, 1, 1, 1, 1, 3, 5, 20, 32, 2623, 4623, 4623, 4624},
	},
}

func TestQuicksort(t *testing.T) {

	for _, v := range testElements {
		t.Run(v.Name, func(t *testing.T) {
			toSort := v.Vals

			Quicksort(toSort)
		})
	}
}
