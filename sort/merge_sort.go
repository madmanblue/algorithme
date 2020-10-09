// 归并排序
package sort

func MergeSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	merge(arr, 0, len(arr)-1)

}

// 将数组拆分
func merge(arr []int, li int, ri int) {
	if li == ri {
		return
	}
	midi := li + ((ri - li) >> 1)
	merge(arr, li, midi)
	merge(arr, midi+1, ri)

	sortArr(arr, li, midi, ri)
}

// 数组排序
func sortArr(arr []int, li int, midi int, ri int) {
	tmp := make([]int, ri-li+1)

	p1 := li
	p2 := midi + 1

	index := 0

	for p1 <= midi && p2 <= ri {
		if arr[p1] > arr[p2] {
			tmp[index] = arr[p1]
			p1++
		} else {
			tmp[index] = arr[p2]
			p2++
		}
		index++
	}

	for p1 <= midi {
		tmp[index] = arr[p1]
		index++
		p1++
	}

	for p2 <= ri {
		tmp[index] = arr[p2]
		index++
		p2++
	}

	for i := 0; i < len(tmp); i++ {
		arr[li+i] = tmp[i]
	}
}
