package sort

import "fmt"

func BubbleSort(arr []int, asc bool) {

	if len(arr) < 2 {
		fmt.Println("无需排序")
		return
	}

	if asc {
		fmt.Println("正序排序")
		for i := 0; i < len(arr)-1; i++ {
			for j := i + 1; j < len(arr); j++ {
				if arr[i] > arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}

	} else {
		fmt.Println("倒序排序")
		for i := 0; i < len(arr)-1; i++ {
			for j := i + 1; j < len(arr); j++ {
				if arr[i] < arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
	}

	fmt.Println(arr)
}
