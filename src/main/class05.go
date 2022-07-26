package main

import "fmt"

/*
Author:LiuShihao
Date: 2022/7/25 16:58
Desc:
*/

func main() {
	var num int = 32
	//变量前加上&符合，可以打印对象的辞内存地址： 0x14000122008
	fmt.Println(&num)
	//定义一个指针变量，类型为*int ，&num是一个内存地址：0x14000122008
	var ptr *int = &num
	fmt.Println(ptr)
}
