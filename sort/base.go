// 这是包注释
package sort

//	程序启动方式
//	1. go run base.go
//	2. go build base.go ./base

//	go.mod 标明一个根路径，代码可以引用当前路径下的方法 eg: init.EchoInitInfo()
import (
	"algorithme/base"
	"fmt"
	"sort"
)

// 这是方法注释
func BaseSort() {
	a := []int{20, 18, 30, 3, 90, 22}
	fmt.Println("未排序前的数组", a)

	sort.Ints(a)
	fmt.Println("排序后的数组", a)

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println("倒序排序数组", a)

	base.EchoInitInfo()
}
