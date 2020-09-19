// 变量声明
package typeall

import "os"

var A int = 9
var B = 10

func TypeAll(x int, y int) (int, string) {
	return x + y, os.Getenv("PATH")
}
