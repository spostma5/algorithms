package algorithms

import (
	"cmp"
	"log/slog"
)

func Quicksort[V cmp.Ordered](items []V) {
	slog.Info("Sorting with quicksort", "unsorted", items)

	slog.SetLogLoggerLevel(slog.LevelDebug)
	quicksort(items, 0, len(items)-1)

	slog.Info("Sorting with quicksort", "sorted", items)
}

func quicksort[V cmp.Ordered](items []V, low, high int) {
	if len(items) <= 1 || high-low <= 1 {
		return
	}

	ind := (high - low) / 2 + low
	pivot := items[ind]

	countLower := low
	for i := low; i <= high; i++ {
		if items[i] < pivot {
			countLower++
		}
	}

	swap(items, ind, countLower)
	ind = countLower

	for i, j := low, high; i < ind && j > ind; {
		for items[i] < pivot {
			i++
		}

		for items[j] > pivot {
			j--
		}

		swap(items, i, j)
		i++
		j--
	}

	quicksort(items, low, ind)
	quicksort(items, ind+1, high)
}

func swap[V cmp.Ordered](items []V, a, b int) {
	tmp := items[a]
	items[a] = items[b]
	items[b] = tmp
}
