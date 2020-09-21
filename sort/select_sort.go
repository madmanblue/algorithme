package sort

func SelectSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	for i := 0; i < len(arr); i++ {
		k := i
		for j := 0; j < len(arr)-1; j++ {
			if arr[i] > arr[j] {
				k = j
			}
		}
		if k != i {
			arr[i], arr[k] = arr[k], arr[i]
		}
	}
}
