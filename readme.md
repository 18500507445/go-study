## go语言学习
- 简介：抽空每天学习一点，慢慢积累

## 学习地址
* [无闻老师：Go语言编程基础](https://github.com/unknwon/go-fundamental-programming  )  
* [go语言官方教程中文版](https://github.com/Go-zh/tour)  
* [学完前面的内容后进行项目练习](https://github.com/astaxie/build-web-application-with-golang)


## 1. 初识Golang
### 1.1 语法注意事项
* 源文件以"go"为拓展名
* 程序执行入口为main()函数
* 严格区分大小写
* 方法由一条条语句构成，每条语句不需要分号结尾（简洁）
* Go编译器一行行进性编译，所以多条语句不能同一行
* 定义的变量或者导入的包没有用到，代码编译不通过
* 大括号成对出现，缺一不可
* 注释与Java的语法、快捷键相同（行注释//，块注释/**/），官方推荐行注释
* 缩进、取消缩进与Java快捷键相同 
* 运算符两边加空白，var age = 10 + 18
* 统一语法格式，如下是错误编译不通过，大括号必须在上面，这点和Java、C++区分开（我感觉挺好，省着有强迫症哈哈）
    ~~~go
    func main()
    {
    }
    ~~~
* 行数长度约定，max80个字符，换行可以"，"拼接后
    ~~~go
    fmt.Println("123",
    "456")
    ~~~

### 1.2 API
* 官网提供了大量的标准库
* [中文网在线文档](https://studygolang.com/pkgdoc)

## 2. 基本变量与类型
### 2.1 变量（全局、局部变量）
* var name 类型 = value
* 如果没有写变量类型，可以根据value值进性自动推断
* var可以省略 name := value
* 支持多变量同时生命

### 2.2 数据类型
* 基本数据类型
  - 数值型（整数类型int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64、byte，浮点类型float32、float64）
  - 字符类型（没有单独的字符型，使用byte保存单个字符字符）
  - 字符串
  - 布尔类型
* 派生数据类型（复杂数据类型）
  - 指针
  - 数组
  - 结构体
  - 管道
  - 函数
  - 切片
  - 接口
  - Map

### 2.3 整数类型（默认int，尽量定义范围符合的）
* 有符号整数类型（int8、int16、int32、int64），相当于Java的byte、shot、int、long
 - 补充说明01111111转十进制就是127
* 无符号整数类型（uint8、uint16、uint32、uint64），只能表示正数
* 其它其它类型（int、uint、rune、byte）