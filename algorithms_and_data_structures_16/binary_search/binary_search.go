package binarysearch

func binSearch(arr []int, val int) int {
	lb := 0
	rb := len(arr) - 1

	for lb <= rb {
		mid := (lb + rb) / 2
		if arr[mid] == val {
			return mid
		}

		if arr[mid] < val {
			lb = mid + 1
		} else {
			rb = mid - 1
		}
	}

	return -1
}
