package main

import (
	"fmt"
	"github.com/google/uuid"
	"go-lean-note/lib/model"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unsafe"
)

const (
	LINE = "-------------------------"
)

func main() {

	//ifElseFunc()

	//returnErrFunc1()

	//returnErrFunc2()

	//switchFunc()

	//forFunc1()
	//forFunc2()
	//forFunc3()
	//forFunc4()

	//arrFunc()
	//arrFunc2()

	//sliceFunc1()

	//sliceFun2()
	//sliceFunc3()

	//sliceFunc4()

	//mapFunc1()
	//mapFunc2()

	//fFunc1()

	//fFunc2()
	//structFunc1()

	//packageFunc1()
	tagFunc1()

}

func tagFunc1() {

	//type TagModel struct {
	//	Name string "这个是名称"
	//	Memo string "这里是Memo"
	//}

	t := new(model.TagModel)
	t.Name = "admin"
	t.Memo = "abxsasdf sdaadd啊打发"

	rfType := reflect.TypeOf(t)

	f, ok := rfType.Elem().FieldByName("Name")
	if ok {
		fmt.Println(f.Tag)
	}

}

func packageFunc1() {
	i := model.NewItem("物料名物料名物料名物料名物料名物料名物料名物料名", 123.2)
	(*i).Name = "a"
	fmt.Println(i, unsafe.Sizeof(i))

	fmt.Println(*i, unsafe.Sizeof(*i))
	fmt.Println(&i, unsafe.Sizeof(&i))
	fmt.Println(LINE)

	u := new(model.Unit)
	u.Id = uuid.NewString()
	u.Name = "个;asdfpasdfjopwjfepojas;dkljdsaf"

	fmt.Println(u, unsafe.Sizeof(u))
	fmt.Println(*u, unsafe.Sizeof(*u))
	fmt.Println(&u, unsafe.Sizeof(&u))
	fmt.Println(LINE)

	u2 := model.NewUnit("箱")
	fmt.Println(u2, unsafe.Sizeof(u2))
	fmt.Println(*u2, unsafe.Sizeof(*u2))
	fmt.Println(&u2, unsafe.Sizeof(&u2))

}

func structFunc1() {
	type User struct {
		Name string
		Age  int
		Meno string
	}

	u := User{Name: "admin", Age: 12, Meno: "memo_memo_123123"}

	fmt.Println(u)
	fmt.Println(u.Meno)
	fmt.Println(LINE)

	fmt.Println(LINE)

	upperFunc := func(u User) string {
		r := strings.ToUpper(u.Name)
		u.Name = r
		return r
	}

	rs := upperFunc(u)
	fmt.Println(rs)
	fmt.Println(u)

	fmt.Println(LINE)

	upperFunc2 := func(u *User) {
		u.Name = strings.ToUpper(u.Name)
	}
	upperFunc2(&u)
	fmt.Println(u)

	fmt.Println(LINE)

	u2 := &User{Name: "aaa", Age: 33, Meno: "sadfasdfsf"}
	upperFunc2(u2)
	fmt.Println(*u2)

}

func fFunc2() {

	s := "abcderfsadfa;kwqerqwer"
	rs := strings.ContainsFunc(s, customContain)
	fmt.Println(rs)

	fmt.Println("匿名函数")

	cusFunc := func(c rune) bool {
		return c == ';'
	}
	rs2 := strings.ContainsFunc(s, cusFunc)
	fmt.Println(rs2)

}

func customContain(c rune) bool {
	return c == ';'
}

func fFunc1() {

	rs1, rs2 := ff1(2, 3, 4, 5)
	fmt.Println(rs1)
	fmt.Println(rs2)

	input1 := 1
	input2 := []int{4, 5, 6}
	r1, r2 := ff1(input1, input2...)
	fmt.Println(r1)
	fmt.Println(r2)
}

func ff1(a int, b ...int) (rs1 int, rs2 []int) {
	rs1 = a * 10
	for _, v := range b {
		rs2 = append(rs2, v*100)
	}
	return
}

func mapFunc2() {
	m := map[string]string{}

	v, ok := m["a"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("no")
	}

	v2 := m["a"]
	fmt.Println(v2)

	m["a"] = "aaaa"
	m["b"] = "bbbb"
	m["c"] = "cccc"
	m["d"] = "ddddd"

	for key, _ := range m {
		fmt.Println(key)
	}
}

func mapFunc1() {
	m := map[string]string{}
	m["a"] = "123"
	m["a"] = "222"
	m["b"] = "333"

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, "->", v)
	}

	v1 := m["b"]

	fmt.Println(v1)
	v1 = "aaaaaaa"

	for k, v := range m {
		fmt.Println(k, "->", v)
	}
}

func sliceFunc4() {
	s := []string{"a", "bb", "ab", "dc", "d", "c"}
	fmt.Println(s)

	sort.Strings(s)
	fmt.Println(s)

	rs := sort.SearchStrings(s, "w")
	fmt.Println(rs)

}

func sliceFunc3() {
	s := make([]string, 0)
	s = append(s, "1", "2")
	fmt.Println(s)
	fmt.Println(LINE)

	sFrom := []string{"a", "b"}
	sTo := make([]string, len(sFrom))
	fmt.Println(sTo)
	copy(sTo, sFrom)
	fmt.Println(sTo)

	fmt.Println(LINE)
	bFrom := []string{"a", "b"}
	bTo := append([]string{}, bFrom...)
	fmt.Println(bTo)

}

func sliceFun2() {

	s := new([]int)
	rs := f4(s)
	fmt.Println(rs)
	fmt.Println(s)
	fmt.Println(LINE)

	s2 := new([]int)
	f5(s2)
	fmt.Println(s2)
	fmt.Println(*s2)

}

func f5(a *[]int) {
	*a = append(*a, 3, 4, 5)
}

func f4(a *[]int) []int {
	rs := append(*a, 1, 2, 3, 4)
	return rs

}

func sliceFunc1() {
	s := []int{1, 2, 3, 4, 5}

	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(LINE)

	s2 := []int{}
	s2 = append(s2, 1, 2, 3, 4, 5, 6)

	fmt.Println(s2[1:])         // [2 3 4 5 6]
	fmt.Println(s2[:2])         // [1 2]
	fmt.Println(s2[len(s2)-2:]) // [5 6]
	fmt.Println(s2[cap(s2)-2:]) // [5 6]
	fmt.Println(s2[:cap(s2)-2]) // [1 2 3 4]
}

func arrFunc2() {
	arr1 := [...]string{"a", "b", 5: "aaa"}

	fmt.Println(arr1)
	fmt.Println(len(arr1))
	fmt.Println(LINE)

	arr2 := [...]float64{12.34, 123.00, 12356.00001}

	rs := f3(arr2[:])
	fmt.Println(rs)

}

func f3(a []float64) (sum float64) {
	for _, f := range a {
		sum += f
	}
	return
}

func arrFunc() {
	var arr1 [5]string

	for i, _ := range arr1 {
		arr1[i] = strconv.Itoa(i) + "s"
	}

	fmt.Println(arr1)

	fmt.Println(LINE)

	rs1 := f1(arr1)
	fmt.Println(rs1)
	fmt.Println(arr1)
	fmt.Println(LINE)
	fmt.Println(arr1)
	f2(&arr1)
	fmt.Println(arr1)

}

func f2(a *[5]string) {
	for i, _ := range a {
		a[i] = strconv.Itoa(i) + "eeee"
	}
}

func f1(a [5]string) [5]string {
	for i, _ := range a {
		a[i] = strconv.Itoa(i) + "eeee"
	}
	return a
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
