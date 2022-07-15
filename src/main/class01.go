package main

import "fmt"

//定义全局变量 方式一：
var num = 18
var gender = "男"

//定义全局变量 方式二：
var (
	n6 = 99
	n7 = 99.9
	n8 = "哈哈"
	n9 = true
)

func main() {
	// 1. 定义变量
	var age1 int
	// 2.赋值  如果不赋值，变量为默认值，int类型默认值为0
	age1 = 22
	// 3.使用
	fmt.Println("age1 = ", age1)

	var age2 int = 23
	fmt.Println("age2 = ", age2)

	//如果没有定义变量类型，程序会自动根据=号后的值，自动判断变量类型
	var age3 = 23
	fmt.Println("age3 = ", age3)

	// 直接省略 var关键字和类型 ，通过:=定义变量和赋值
	age4 := 24
	fmt.Println("age4 = ", age4)

	//声明多个变量
	var n1, n2, n3 int
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	var n4, name, n5 = 66, "zhangsan", 7.8
	fmt.Println(n4)
	fmt.Println(name)
	fmt.Println(n5)

	weight, hight := 60.1, 180
	fmt.Println(weight)
	fmt.Println(hight)

	fmt.Println(gender)
	fmt.Println(num)
}
