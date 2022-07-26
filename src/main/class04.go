package main

import (
	"fmt"
	"strconv"
)

/*
Author:LiuShihao
Date: 2022/7/25 16:04
Desc: 字符串转换为其他类型
*/

func main() {
	fmt.Println("-- -----字符串转换为其他类型----------")
	var s1 string = "true"
	var s2 string = "34"
	var s3 string = "3.1415926"

	//func ParseBool(str string) (bool, error)  一个入参，两个返回值
	//第一个返回值为转换后的bool类型，第二个参数为可能出现测错误，可以使用_忽略
	var bb bool
	bb, _ = strconv.ParseBool(s1)
	fmt.Println("string -> bool:", bb)
	// 第二个参数为10进制，第三个参数为int64类型
	var num int64
	num, _ = strconv.ParseInt(s2, 10, 64)
	fmt.Println("string -> int64:", num)
	//第二个参数表示float64类型
	var f1 float64
	f1, _ = strconv.ParseFloat(s3, 64)
	fmt.Println("string -> float64:", f1)

}
