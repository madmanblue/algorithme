package main

import "algorithme/sort"

func main() {
	sort.BaseSort()
	arr := []int{19, 28, 20, 49, 57, 10, 13}
	sort.InsertSort(arr)
	arr1 := []int{8, 7, 9}
	sort.BubbleSort(arr1, false)
}
