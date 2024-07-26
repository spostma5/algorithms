package algorithms

import (
	"cmp"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sortTest[T cmp.Ordered] struct {
	Name     string
	Vals     []T
	Expected []T
}

var testElements = []sortTest[int]{
	// {
	// 	Name:     "Case #1",
	// 	Vals:     []int{1, 5, 3},
	// 	Expected: []int{1, 3, 5},
	// },
	// {
	// 	Name:     "Case #2",
	// 	Vals:     []int{1, 5, 3, 4623, 20, 0, 1, 1, 1, 32, 0, 4623, 2623, 4624},
	// 	Expected: []int{0, 0, 1, 1, 1, 1, 3, 5, 20, 32, 2623, 4623, 4623, 4624},
	// },
	{
		Name:     "Case #3",
		Vals:     []int{0, 0, 0, 0, 0, 0},
		Expected: []int{0, 0, 0, 0, 0, 0},
	},
	{
		Name:     "Case #4",
		Vals:     []int{1, 2, 3, 4, 3, 2, 1, 0, -1, -2},
		Expected: []int{-2, -1, 0, 1, 1, 2, 2, 3, 3, 4},
	},
	{
		Name:     "Case #5",
		Vals:     []int{},
		Expected: []int{},
	},
	{
		Name:     "Case #6",
		Vals:     []int{2, 1},
		Expected: []int{1, 2},
	},
	{
		Name:     "Case #7",
		Vals:     []int{1},
		Expected: []int{1},
	},
}

func TestQuicksort(t *testing.T) {
	for _, v := range testElements {
		t.Run(v.Name, func(t *testing.T) {
			toSort := v.Vals

			Quicksort(toSort)
			assert.Equal(t, v.Expected, toSort)
		})
	}
}
