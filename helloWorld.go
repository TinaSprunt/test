package main

import "fmt"

func main() {

	var stockcode = 123
	var msg = "2021-3-4"
	var url = "code=%d&enDate=%s"
	var target = fmt.Sprintf(url, stockcode, msg)

	fmt.Println(target)

	//可以只定义变量不赋值，这样会输出对应变量类型的0值，比如string就是 ""
	var i int
	var f float64
	var b bool
	var s string

	fmt.Println("%v %v %v %q\n", i, f, b, s)

	//全局变量是允许声明但不使用的,其他变量声明后必须使用

	testFun("2020")

}

func testFun(year string) (int string) {
	fmt.Println(year + "-03-04")
	var res string = year + "-03-04"
	return 0 res
}
