// 定义常量和枚举
package constant

import "fmt"

const commonParam = "common"
const startNo int = 0

type Color int

const (
	RED Color = iota
	YELLOW
	GREEN
	BLUE
	BLACK
)

func Col() {

	fmt.Println(RED)
	fmt.Println(BLACK)
}
