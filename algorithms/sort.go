package algorithms

import (
	"cmp"
)

func InsertionSort[V cmp.Ordered](items []V) {
	if len(items) <= 1 {
		return
	}

	for j := 1; j < len(items); j++ {
		for i := j - 1; i >= 0; i-- {
			if items[i] > items[i+1] {
				swap(items, i, i+1)
			} else {
				break
			}
		}
	}
}

func Quicksort[V cmp.Ordered](items []V) {
	quicksort(items, 0, len(items)-1)
}

func quicksort[V cmp.Ordered](items []V, low, high int) {
	if high < low {
		return
	}

	ind := (high-low)/2 + low
	pivot := items[ind]

	countLower := low
	for i := low; i <= high; i++ {
		if items[i] < pivot {
			countLower++
		}
	}

	swap(items, ind, countLower)
	ind = countLower

	for i, j := low, high; i < j; {
		for items[i] < pivot {
			i++
		}

		for j >= low && items[j] >= pivot {
			j--
		}

		if i >= j {
			break
		}

		swap(items, i, j)
		i++
		j--
	}

	quicksort(items, low, ind-1)
	quicksort(items, ind+1, high)
}

func swap[V cmp.Ordered](items []V, a, b int) {
	items[a], items[b] = items[b], items[a]
}
