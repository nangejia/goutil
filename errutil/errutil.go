//这是一个处理错误的口
package errutil

//如果有错误，则直接抛出错误
func Error(err error) {
	if err != nil {
		panic(err)
	}
}
