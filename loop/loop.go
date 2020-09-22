// 代码结构相关 ， 循环，if-else， switch， select - channel
package main

import "fmt"

func main() {
	arr := []int{68, 99, 3, -1, 18, 77, 80}
	loop(arr)
	break_continue_loop(arr)
}

func loop(arr []int) {
	for ix, val := range arr {
		fmt.Printf("当前学生 %d, 分数 %d， 成绩 %s\n", ix+1, val, rank(val))
	}
}

func rank(i int) string {
	r := ""
	if i < 0 {
		r = "wdnmd, 这也是人能考出来的成绩！"
	} else if i > 80 {
		r = "成绩优秀"
	} else if i <= 80 && i > 60 {
		r = "成绩及格"
	} else {
		r = "不及格"
	}
	return r
}

func break_continue_loop(arr []int) {
	for ix, val := range arr {

		if val == -1 {
			break
		}
		if val%2 == 0 {

			fmt.Println(ix, "-", val)
			continue
		}

	}
}
