package main

import (
	"algorithme/sort"
	"fmt"
)

func main() {
	//sort.BaseSort()
	arr := []int{19, 28, 20, 49, 57, 10, 13}
	//sort.InsertSort(arr)
	//arr1 := []int{8, 7, 9}
	//sort.BubbleSort(arr1, false)
	//
	//constant.Col()
	//
	//all, os := typeall.TypeAll(typeall.A, typeall.B)
	//fmt.Println("typeall return ", all)
	//fmt.Println("typeall return ", os)
	//
	//sort.SelectSort(arr1)
	//fmt.Println(arr1)
	//
	//typeall.StrUtil("fmt.print")

	sort.ShellSort(arr)
	fmt.Println(arr)
}
