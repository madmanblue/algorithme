//指针使用
package main

import "fmt"

func main() {

	i := 5
	fmt.Printf("i 值 is : %d \n i 指针 val is : %p\n", i, &i)
	var ptr *int
	ptr = &i
	fmt.Printf("ptr 指针 is : %p\n ptr 值 is : %d\n", ptr, *ptr)

	*ptr = 6
	fmt.Printf("i 值 为 %d\n", i)

	const j = 8
	fmt.Printf("j 值 %d\n", j)

}
