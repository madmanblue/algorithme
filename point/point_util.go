package main

import "fmt"

func main() {

	i := 5
	fmt.Printf("i 值 is : %d \n i 指针 val is : %p\n", i, &i)
	var ptr *int
	ptr = &i
	fmt.Printf("ptr 指针 is : %p\n ptr 值 is : %d\n", ptr, *ptr)
}
