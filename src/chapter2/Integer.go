package main

import (
	"fmt"
	"unsafe"
)

func main() {

	fmt.Println("-------------有符号整数类型-------------")

	//1字节 范围：-2^7 ~ 2^7-1 (-128 ~ 127)
	var num1 int8 = 7
	fmt.Println("num1:", num1)

	//2字节 范围：-2^15 ~ 2^15-1
	var num2 int16 = 15
	fmt.Println("num2:", num2)

	//4字节 范围：-2^31 ~ 2^31-1
	var num3 int16 = 31
	fmt.Println("num3:", num3)

	//8字节 范围：-2^63 ~ 2^63-1
	var num4 int16 = 63
	fmt.Println("num4:", num4)

	fmt.Println("-------------无符号整数类型-------------")

	//1字节 范围：0 ~ 2^7-1 (0 ~ 255)
	var num5 uint8 = 7
	fmt.Println("num5:", num5)

	//2字节 范围：0 ~ 2^15-1 (0 ~ 255)
	var num6 uint16 = 15
	fmt.Println("num6:", num6)

	//4字节 范围：0 ~ 2^31-1 (0 ~ 255)
	var num7 uint32 = 31
	fmt.Println("num7:", num7)

	//8字节 范围：0 ~ 2^63-1 (0 ~ 255)
	var num8 uint64 = 63
	fmt.Println("num8:", num8)

	fmt.Println("-------------其它整数类型-------------")

	//32位系统 4字节 -2^31 ~ 2^31-1
	//64位系统 8字节 -2^63 ~ 2^63-1
	//观察int可以省略，所以是默认数据类型
	var num9 int = 63
	fmt.Printf("num9的类型是: %T", num9)
	fmt.Println()
	fmt.Println("num9占用字节数：", unsafe.Sizeof(num9))

	//32位系统 4字节 0 ~ 2^31-1
	//64位系统 8字节 0 ~ 2^63-1
	var num10 uint = 63
	fmt.Println("num10:", num10)
}
