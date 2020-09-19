// 函数的格式
package function

// 无返回值，首字母小写私有方法
func has_no_return_value() {

}

// 首字母大写是公共方法
func Pub_fuc(receiver int) {

}

// 返回一个参数
func has_return_val() int {
	return 1
}

// 返回多个参数
func has_multi_return_val() (i int, b byte) {
	return 7, 'd'
}

// 传入指针参数
func param_is_point(point *int, i iv) {

}

type (
	iv int
	fv float32
	sv string
)

// 自定义别名
func self_params(i iv, fl fv, s sv) {

}

// ???
func (t sv) Method1() {
	//...
}
