package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{20, 18, 30, 3, 90, 22}
	fmt.Println("未排序前的数组", a)

	sort.Ints(a)
	fmt.Println("排序后的数组", a)

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println("倒序排序数组", a)
}
