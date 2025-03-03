package selectionsort

func lower(arr []int) int {
	lower := arr[0]
	lowerI := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < lower {
			lower = arr[i]
			lowerI = i
		}
	}
	return lowerI
}

func Appoff(arr []int, i int) []int {
	if i == len(arr) {
		return arr[:i]
	}
	if i == 0 {
		return arr[1:]
	}
	return append(arr[:i], arr[i+1:]...)
}

func SelectionSort(arr []int) []int {
	orderedArr := []int{}
	temp := arr

	for i := 0; i < len(arr); i++ {
		tLower := lower(temp)
		orderedArr = append(orderedArr, temp[tLower])
		temp = Appoff(temp, tLower)
	}
	return orderedArr
}
