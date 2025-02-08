package binarysearch

func BinarySearch(arr []int, target int) int {
	high := len(arr) - 1
	low := 0

	for low <= high {
		midle := (high + low) / 2
		guess := arr[midle]
		switch {
		case target > guess:
			low = midle + 1
		case target < guess:
			high = midle - 1
		case target == guess:
			return midle
		}
	}
	return -1
}
