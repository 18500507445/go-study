package main

import "fmt"

// 成员变量、全局变量
var (
	n6 = 100
	n7 = 99
)

func main() {
	//局部变量
	var age int = 19
	fmt.Println("age:", age)

	var score float32 = 90.02
	fmt.Println("score:", score)

	//省略 变量类型
	var score2 = 88.8888
	fmt.Println("score2:", score2)

	//省略 var
	sex := "男"
	fmt.Println("sex:", sex)

	//多变量声明
	var n1, n2, n3 = 1, "2", 3.03
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	n4, n5 := "我爱你", 5.05
	fmt.Println("n4:", n4)
	fmt.Println("n5:", n5)

	fmt.Println("n6:", n6)
	fmt.Println("n7:", n7)
}
