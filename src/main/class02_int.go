package main

import (
	"fmt"
	"unsafe"
)

/*
有符号整形   超出范围报错  编译直接报错
int8     byte    -128 - 127
int16    short     - 2^15 ~ 2^15-1
int32    int		- 2^31 ~ 2^31-1
int64    long		- 2^63 ~ 2^63-1

无符号整形：不能表示负数
uint8    0-255
uint16   0-2^16-1
uint32   0-2^32-1
uint64   0-2^64-1

其他整数类型
int   有符号   32位系统4字节，64位系统8字节
uint  无符号   32位系统4字节，64位系统8字节
rune  有符号   对应int32
byte  无符号   对应uint8 0-255

浮点数类型
float32 4字节
float64 8字节

*/

func main() {
	fmt.Println("------------整数类型--------------")
	var n1 int8 = -23
	fmt.Println("n1 = ", n1)
	var n2 uint8 = 23
	fmt.Println("n2 = ", n2)
	var n3 = 32
	//Printf 格式化打印 %T 类型
	fmt.Printf("n3的类型是：%T", n3)
	fmt.Println()
	fmt.Println("n3的字节是：", unsafe.Sizeof(n3))
	fmt.Println("------------浮点数类型--------------")
	//float32类型会损失精度，Go中默认的浮点数类型是float64
	var n4 float32 = 3.14e+2
	var n5 float64 = 3.14e-2
	var n6 float32 = 2.9346394623934
	var n7 float64 = 2.9346394623934

	fmt.Println("n4 = ", n4)
	fmt.Println("n5 = ", n5)
	fmt.Println("n6 = ", n6)
	fmt.Println("n7 = ", n7)
	fmt.Println("-----------字符类型---------------")

	//byte 只能表示0-255
	var c1 byte = 'a'
	var c2 byte = '6'
	var c3 = '<'
	fmt.Println("c1 = ", c1)
	fmt.Printf("c1的对应的字符是:%c", c1)
	fmt.Println()
	fmt.Println("c2 = ", c2)
	fmt.Println("c3 = ", c3)
	fmt.Printf("c3的类型是:%T", c3)

	fmt.Println()
	//byte 只能表示0-255
	var c4 int = '中'
	fmt.Println("c4 = ", c4)
	fmt.Println("------------转义字符--------------")
	//转义字符
	// /b 退格  /n 换行  /r 回车  /t 制表符  /"双引号  /' 单引号 // 反斜杠
	fmt.Println("111\n222")

	fmt.Println("333\r444") //会覆盖之前的

	fmt.Println("555\b666")
	//制表符 每8个字符为一个制表符
	fmt.Println("aaaaaaaabbbb")
	fmt.Println("aaaaa\tbbbb")
	fmt.Println("aaaa\tbbbb")

	fmt.Println("-----------bool类型---------------")

	var flag01 bool = true
	fmt.Println("faag01 = ", flag01)
	var flag02 bool = false
	fmt.Println("flag02 = ", flag02)
	var flag03 bool = 5 < 6
	fmt.Println("flag03 : 5<6 = ", flag03)

	fmt.Println("-----------字符串类型---------------")
	var str string = "你好"
	fmt.Println(str)
	s1 := "ahello go!"
	fmt.Println(s1)
	//字符串是不可变的 ，字符串一旦定义好，其中的值不能改变
	s1 = "let's go"
	//s1[0] = 'L'   //不能改变
	fmt.Println(s1)
	//字符串的表示形式 1.如果字符串中存在特殊字符，可以使用反引号``
	var ss = `
	fmt.Println("-----------字符串类型---------------")
	var str string = "你好"
	fmt.Println(str)
	s1 := "ahello go!"
	fmt.Println(s1)
	//字符串是不可变的 ，字符串一旦定义好，其中的值不能改变
	s1 = "let's go"
	//s1[0] = 'L'   //不能改变
	fmt.Println(s1)
	`
	fmt.Println(ss)
	//字符串的拼接
	var s5 string = "abc" + "efg"
	fmt.Println(s5)
	s5 += "hahha"
	fmt.Println(s5)
	//如果字符号拼接换行，则需要将加号留在上一行才可以
	var s6 string = "abc" + "efg" + "abc" + "efg" +
		"abc" + "efg" + "abc" + "efg" + "abc" +
		"efg" + "abc" + "efg" + "abc" + "efg" + "abc" + "efg"
	fmt.Println(s6)

	/*
		基本数据类型的默认值：
		整形 0
		浮点数 0
		布尔值 false
		字符串 ""
	*/

}
