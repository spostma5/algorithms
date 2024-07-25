package algorithms

import (
	"cmp"
	"log/slog"
)

func Quicksort[V cmp.Ordered](items []V) {
	slog.Info("Sorting with quicksort", "unsorted", items)

	slog.SetLogLoggerLevel(slog.LevelDebug)
	quicksort(items)

	slog.Info("Sorting with quicksort", "sorted", items)
}

func quicksort[V cmp.Ordered](items []V) {
	slog.Debug("quicksort", "items", items)
	if len(items) <= 1 {
		return
	}

	ind := len(items) / 2
	pivot := items[ind]

	left := 0
	right := len(items) - 1
	for left < right {
		if items[left] >= pivot && pivot >= items[right] {
			tmp := items[left]
			items[left] = items[right]
			items[right] = tmp

			right--
		}

		left++
	}

	slog.Debug("quicksort", "pivot", pivot, "left", items[:ind], "right", items[ind:])

	quicksort(items[:ind])
	quicksort(items[ind:])
}
