package main

import (
	"fmt"
	"strconv"
)

func main() {

	//ifElseFunc()

	//returnErrFunc1()

	//returnErrFunc2()

	//switchFunc()

	//forFunc1()
	//forFunc2()
	//forFunc3()
	forFunc4()
}

func forFunc4() {
	str := "abc 123"
	for i, s := range str {
		fmt.Println(i, string(s))
	}
}

func forFunc3() {
	i := 5
	for {
		println(i)
		i = i - 1
		if i <= 0 {
			return
		}
	}
}

func forFunc2() {
	i := 5
	for i > 0 {
		i = i - 1
		println(i)
	}
}

func forFunc1() {
	for i := 0; i < 10; i++ {
		println(i)
	}
}

func switchFunc() {
	k := 6
	switch k {
	case 4:
		fmt.Println("1111")
		fallthrough
	case 5:
		fmt.Println("2222")
		fallthrough
	case 6:
		fmt.Println("3333")
		fallthrough
	case 7:
		fmt.Println("4444")
		fallthrough
	case 8:
		fmt.Println("5555")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

func returnErrFunc2() {
	orinStr := "abc"
	rs, _ := strconv.Atoi(orinStr)
	fmt.Println(rs) // rs是int，将打印零值：0
}

func returnErrFunc1() {
	orinStr := "abc"
	rs, err := strconv.Atoi(orinStr)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(rs)
}

func ifElseFunc() {
	if "aa" == "bb" {
		//
	} else {
		//
	}

	if "aa" == "bb" {
		//
	} else if 3 == 2 {
		//
	} else {
		//
	}

}
