// 变量声明
package typeall

import (
	"fmt"
	"os"
)

var A int = 9
var B = 10

func TypeAll(x int, y int) (int, string) {
	return x + y, os.Getenv("PATH")
}

func StrUtil(str string) {
	fmt.Println("当前字符:", str)
	fmt.Println("字符长度：", len(str))
	fmt.Println("首字母:", str[0:1])
	fmt.Println("剩余:", str[1:])
	fmt.Println("other:", str[1:len(str)-1])
}
