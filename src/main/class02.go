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
	var n1 int8 = -23
	fmt.Println("n1 = ", n1)
	var n2 uint8 = 23
	fmt.Println("n2 = ", n2)
	var n3 = 32
	//Printf 格式化打印 %T 类型
	fmt.Printf("n3的类型是：%T", n3)
	fmt.Println()
	fmt.Println("n3的字节是：", unsafe.Sizeof(n3))
	fmt.Println("--------------------------")
	//float32类型会损失精度，Go中默认的浮点数类型是float64
	var n4 float32 = 3.14e+2
	var n5 float64 = 3.14e-2
	var n6 float32 = 2.9346394623934
	var n7 float64 = 2.9346394623934

	fmt.Println("n4 = ", n4)
	fmt.Println("n5 = ", n5)
	fmt.Println("n6 = ", n6)
	fmt.Println("n7 = ", n7)
	fmt.Println("--------------------------")

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

}
