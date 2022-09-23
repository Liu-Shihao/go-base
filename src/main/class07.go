package main

import "fmt"

/*
Author:LiuShihao
Date: 2022/9/23 13:51
Desc: 错误处理 defer + recover
*/

func main() {
	res := test()
	fmt.Println("res:", res)
	fmt.Println("test()执行完成，执行下面代码....")
}

func test() bool {
	var isShown bool
	defer func() {
		err := recover()
		if err != nil {
			isShown = false
			fmt.Println("发生异常：", err)
		} else {
			isShown = true
			fmt.Println("正常执行")
		}
	}()

	num1 := 10
	num2 := 0
	result := num1 / num2 // panic: runtime error: integer divide by zero

	fmt.Println("result:", result)

	return isShown

}
