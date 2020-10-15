// 快速排序 修炼一下给出
package sort

func QuickSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

}

//
func quickSort(arr []int, L int, R int) {

	if L < R {
		pivotNum := pivot(arr, L, R)
		quickSort(arr, L, pivotNum-1)
		quickSort(arr, pivotNum+1, R)
	}

}

func pivot(arr []int, L int, R int) int {
	return 0
}
