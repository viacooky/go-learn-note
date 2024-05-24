package main

import (
	"fmt"
	"os"
)

const (
	AA   = "aaaaa"
	BB   = "bbbbb"
	CC   = "ccccc"
	LINE = "-------------------------"
)

// 全局变量
var (
	HOME = os.Getenv("HOME")
	USER = os.Getenv("USER")
	TEST = "global var"
)

func init() {
	TEST = "初始化函数赋值"
}

func main() {
	// Hello World
	fmt.Println("Hello World")
	// 常量
	//constLean()

	// 变量
	//varLean()

	// bool
	//boolFunc()

	//intergerFunc()

	charFunc()

}

func charFunc() {
	var aChar byte = 'a'
	fmt.Println("var aChar byte = 'a'")
	fmt.Printf("Character: %c", aChar)
	fmt.Println()
	fmt.Printf("interger: %d", aChar)
	fmt.Println()
	fmt.Printf("UTF-8 bytes: %x", aChar)
	fmt.Println()
	fmt.Printf("UTF-8 code point: %U", aChar)
	fmt.Println()

	fmt.Println(LINE)

	var aChar2 byte = '\x41'
	fmt.Println("var aChar2 byte = '\\x41'")
	fmt.Printf("Character: %c", aChar2)
	fmt.Println()
	fmt.Printf("interger: %d", aChar2)
	fmt.Println()
	fmt.Printf("UTF-8 bytes: %x", aChar2)
	fmt.Println()
	fmt.Printf("UTF-8 code point: %U", aChar2)
	fmt.Println()

}

func intergerFunc() {
	//var a32 int32 = 32

	var uintVar int = -32
	fmt.Println(uintVar)

}

func boolFunc() {
	fmt.Println(LINE)
	fmt.Println("bool")
	aVar := 10
	fmt.Println(aVar == 5)
	fmt.Println(aVar == 10)

	fmt.Println(aVar != 5)
	fmt.Println(aVar != 10)

}

func varLean() {
	fmt.Println(LINE)
	fmt.Println("变量")
	var (
		a   int
		b   bool
		str string
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(str)

	var c int = 123
	var d = 333
	var e = false

	fmt.Println(c, d, e)

	// 全局变量
	fmt.Println("全局变量")
	fmt.Println(HOME, USER)

	// 局部变量
	aaa := "这里是局部变量"
	fmt.Println("局部变量")
	fmt.Println(aaa)

	// 变量在初始化赋值
	fmt.Println(TEST)

}

func constLean() {
	fmt.Println(LINE)
	fmt.Println("常量")
	fmt.Println(AA)
}
