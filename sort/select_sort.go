// 选择排序 遍历数组，找到最小值或者最大值，与当前值交换，继续遍历下一个
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
