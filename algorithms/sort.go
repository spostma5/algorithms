package algorithms

import (
	"cmp"
)

func InsertionSort[V cmp.Ordered](items []V) {
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

func BubbleSort[V cmp.Ordered](items []V) {
	for i := range items {
		for j := 0; j < len(items)-i-1; j++ {
			if items[j] > items[j+1] {
				swap(items, j, j+1)
			}
		}
	}
}

func QuickSort[V cmp.Ordered](items []V) {
	quicksort(items, 0, len(items)-1)
}

func MergeSort[V cmp.Ordered](items []V) {
	mergesort(&items, 0, len(items)-1)
}

func mergesort[V cmp.Ordered](items *[]V, low, high int) {
	if high <= low {
		return
	}

	mid := (high - low) / 2

	mergesort(items, low, low+mid)
	mergesort(items, low+mid+1, high)

	merge(items, low, mid, high)
}

func merge[V cmp.Ordered](items *[]V, low, mid, high int) {
	sorted := make([]V, high-low+1)

	for i, j, cnt := low, low+mid+1, 0; cnt <= high-low; cnt++ {
		if i > low+mid {
			sorted[cnt] = (*items)[j]
			j++
		} else if j > high {
			sorted[cnt] = (*items)[i]
			i++
		} else if (*items)[i] < (*items)[j] {
			sorted[cnt] = (*items)[i]
			i++
		} else {
			sorted[cnt] = (*items)[j]
			j++
		}
	}

	//TODO: cleanup
	for i := low; i <= high; i++ {
		(*items)[i] = sorted[i-low]
	}
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
