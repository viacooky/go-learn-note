package main

import (
	"fmt"
	"os"
	"time"
	"unicode"
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

	//integerFunc()

	//charFunc()

	//str()

	timeFunc()

}

func timeFunc() {
	t := time.Now()
	fmt.Println(t.Weekday())
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println("yyyy-MM-dd hh:mm:ss ->" + t.Format("2006-01-02 15:04:05"))
	fmt.Println("yyyyMMdd ->", t.Format("20060102"))
	fmt.Println("yyMMdd ->", t.Format("060102"))
	fmt.Println("hh:mm:ss ->", t.Format("15:04:05"))
	fmt.Println("hh:mm:ss.ffffff ->", t.Format("15:04:05.000000"))

}

func str() {
	a := "aaaa\nbbbb"
	fmt.Println(a)

	fmt.Println(LINE)

	b := `asdfsdf
sadfsadf;klja;sdf \n
casdfasdfe`
	fmt.Println(b)

	fmt.Println(LINE)

	s := "abcdefg"

	fmt.Println(s[0])
	fmt.Println(s[len(s)-1])

	fmt.Println(LINE)

	s2 := "aaa" + "bbb"
	s2 += "dddd"
	fmt.Println(s2)

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

	fmt.Println(LINE)

	var aChar3 rune = '\x41'
	flag := unicode.IsNumber(aChar3)
	fmt.Println(flag)
	flag = unicode.IsLetter(aChar3)
	fmt.Println(flag)

}

func integerFunc() {
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
