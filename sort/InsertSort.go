// 插入排序 参考扑克牌 在一个已经排序好的数组里插入一个值，适合数量少的排序
package sort

import "fmt"

func InsertSort(arr []int) {
	fmt.Println("插入排序前 : ", arr)
	n := len(arr)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	fmt.Println("插入排序后 : ", arr)
}
