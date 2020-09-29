// 希尔排序 在指定步长的情况下将数据分组，进行组内插入排序
package sort

import "fmt"

func ShellSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	for step := len(arr) / 2; step > 0; step /= 2 {
		fmt.Println("step", step)
		for i := step; i < len(arr); i++ {
			fmt.Println("i = ", i)
			for j := i - step; j >= 0 && arr[j+step] < arr[j]; j -= step {
				fmt.Println("j = ", j)
				arr[j], arr[j+step] = arr[j+step], arr[j]
			}
		}
	}

}
