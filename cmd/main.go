package main

import (
	"algo/algorithms"
	"log/slog"
)

func main() {
	slog.Info("Starting algorithms...")

	items := []int{1,3,2,8,3,1,0,99,30,1234,7}
	algorithms.Quicksort(items)
}
