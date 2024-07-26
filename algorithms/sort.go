package algorithms

import (
	"cmp"
)

func Quicksort[V cmp.Ordered](items []V) {
	quicksort(items, 0, len(items)-1)
}

func quicksort[V cmp.Ordered](items []V, low, high int) {
	if high < low {
		return
	}

	if len(items) <= 1 || high-low <= 1 {
		if items[low] > items[high] {
			swap(items, low, high)
		}
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

		for items[j] >= pivot {
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
	tmp := items[a]
	items[a] = items[b]
	items[b] = tmp
}
