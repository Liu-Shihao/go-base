package main

import (
	"fmt"
	"strconv"
)

/*
Author:LiuShihao
Date: 2022/7/25 14:43
Desc: 基本数据类型的转换，在go中所有的类型转换都是显示的


%v	值的默认格式表示
%+v	类似%v，但输出结构体时会添加字段名
%#v	值的Go语法表示
%T	值的类型的Go语法表示
%%	百分号
%t	单词true或false
%b	表示为二进制
%c	该值对应的unicode码值
%d	表示为十进制
%o	表示为八进制
%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
%x	表示为十六进制，使用a-f
%X	表示为十六进制，使用A-F
%U	表示为Unicode格式：U+1234，等价于"U+%04X"
%b	无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat
%e	科学计数法，如-1234.456e+78
%E	科学计数法，如-1234.456E+78
%f	有小数部分但无指数部分，如123.456
%F	等价于%f
%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
%s	直接输出字符串或者[]byte
%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
%x	每个字节用两字符十六进制数表示（使用a-f）
%X	每个字节用两字符十六进制数表示（使用A-F）
%p	表示为十六进制，并加上前导的0x
%f:    默认宽度，默认精度
%9f    宽度9，默认精度
%.2f   默认宽度，精度2
%9.2f  宽度9，精度2
%9.f   宽度9，精度0
*/

func main() {
	var i1 int = 23
	//var f1 float32  = n1  //无法进行转化
	fmt.Println(i1)
	var f1 float32 = float32(i1)
	fmt.Println(f1)

	// int8 的范围是从-128 ~ 127 从大范围转换为小范围 会损失精度 ，编译不会报错
	var i2 int32 = 88888
	var i3 int8 = int8(i2)
	fmt.Println(i3)

	var i4 int32 = 888888
	var i5 int64 = int64(i4) + 50
	fmt.Println(i5)

	fmt.Println("------方式1：fmt.Sprintf（）-----------")

	var f2 float32 = 2.3333

	//1.使用fmt.Sprintf()函数
	// 整数 类型
	var s1 string = fmt.Sprintf("%d", i1)
	fmt.Printf("s1对应的类型是%T,对应的值是%v", s1, s1)
	fmt.Println()

	// 浮点数 类型
	var s2 string = fmt.Sprintf("%f", f2)
	fmt.Printf("s2对应的类型是%T,对应的值是%v", s2, s2)
	fmt.Println()

	// 布尔 类型
	var b1 bool = false
	var s3 string = fmt.Sprintf("%t", b1)
	fmt.Printf("s3对应的类型是%T,对应的值是%v", s3, s3)
	fmt.Println()

	fmt.Println("-------方式2：strconv----------")

	var n1 int = 19
	var s4 string = strconv.FormatInt(int64(n1), 10) //第一个参数必须转为int64类型，第二个参数指定字面值的进制形式为10进制
	fmt.Printf("s4对应的类型是%T,对应的值是%v", s4, s4)
	fmt.Println()

	var f3 float64 = 3.1415
	//第三个参数为小数点后保留位数，第四个参数为float64类型
	var s5 string = strconv.FormatFloat(f3, 'f', 9, 64)
	fmt.Printf("s5对应的类型是%T,对应的值是%v", s5, s5)
	fmt.Println()

	var flag bool = true
	//第三个参数为小数点后保留位数，第四个参数为float64类型
	var s6 string = strconv.FormatBool(flag)
	fmt.Printf("s6对应的类型是%T,对应的值是%v", s6, s6)
	fmt.Println()

}
