package main

func SelectionSort(ints []int) {
	for i := len(ints) - 1; i > 0; i-- {
		maxIdx := 0
		for j := 1; j <= i; j++ {
			if ints[j] > ints[maxIdx] {
				maxIdx = j
			}
		}

		if i != maxIdx {
			Swap(ints, i, maxIdx)
		}
	}
}
