# Day1

## 项目目录结构

参考

* https://tonybai.com/2022/04/28/the-standard-layout-of-go-project/
* https://github.com/golang-standards/project-layout/blob/master/README_zh.md


## 命名

### 项目名（仓库名）

* 可使用字母/数字
* 多单词用`-`分隔
* 不建议驼峰、下划线`_`

GOOD
```
user
user-api
user-service
```

BAD
```
user_api
userApi
```
---------

### 文件名

* 使用小写和下划线
  * 如 `http_server.go`
* 测试文件，以 `_test.go` 结尾
  * 如 `http_server_test.go`
* 平台特定文件，应该在文件名后加入平台名
  * 如 `http_server_win.go`

---------

### 包名

* 简短
* **与目录名保持一致**
* 避免使用`util` `common` 之类的包名
  * 按提供的内容命名，而非包含的包


---------

### 常量

* 可用小驼峰命名   (lower-camelcase)

```
const tadayNews = "Hello"
```

---------

### 变量名

> 由于Golang的特殊性（用大小写来控制函数的可见性）

* 作用域越小，命名越短
  * for 循环内，用`i`表示`index`
* 驼峰式
* 特有名词根据可见性，使用全部大写或小写
  * `apiClient`
  * `URLString`
* 若变量为`bool`类型，名称可考虑以 `Has` `Is` `Can` `Allow` 开头
```
var isVIP bool
var canRead bool
```

---------

### 函数/方法命名

> 由于Golang的特殊性（用大小写来控制函数的可见性）

* 驼峰命名 (lower-camelcase / upper-camelcase)

```
func AddStr() string{
    // ....
    return ""
}

func removeCtx() bool {
	return true
}
```

#### Get Set 方法

* Go 不提供Get、Set方法的自动支持，需要自己提供
* 不要在方法名用Get
* 可在方法名前用Set

GOOD
```
obj.User()

obj.SetUser("xxxx")
```

BAD
```
obj.GetUser()
```

---------

### 结构体

* 简短
* 驼峰命名，根据访问清空控制大写/小写
* `struct` 的声明初始化采用多行书写

```
type CustomerOrder struct {
    Name string
    Address string
}
```

---------

### 接口

* 接口命名一般以 `er` 结尾
  * `Reader`
  * `Writer`

```
type Reader interface{
    // ......
}
```

---------

### 注释

* 一律使用双斜线 `//`


---------

### 错误处理

错误处理原则

* 不能丢弃任何有返回err的调用，不使用时， `_` 丢弃，必须全部处理
* 错误尽早处理

GOOD
```
if err != nil{
  // 处理错误 
  return
}
// 进入正常逻辑
```

BAD
```
if err != nil {
  // 处理错误 
  return
} else {
  // 进入正常逻辑
}

```



## 常量

常量使用关键字 const 定义，用于存储不会改变的数据

格式：`const identifier [type] = value`

```
// 显式类型
const URL string = "https://aaa.com"

// 隐式类型
const URL = "https://aaa.com"
```

可多行定义

```
const (
    a = "aaa"
    b = "bbb"
)
```

## 变量

* 关键字：`var`
* 格式：`var identifier type`

```
var a int
var b bool
var c string

// 也可以写作

var (
    a   int
    b   bool
    str string
)

// 初始化赋值
// 显式
var c int = 123
// 隐式
var e = false
```

### 全局变量

这种写法主要用于声明包级别的全局变量

```
var (
    HOME = os.Getenv("HOME")
    USER = os.Getenv("USER")
)
```

### 局部变量

在函数体内声明局部变量时，应使用简短语法： `:=`

```
a := 1
```

### init函数中初始化变量

* 除了可以在全局声明中初始化，也可以在init函数中初始化
* init函数不能人为调用，而是在每个包初始化后自动执行
* 优先级比 main 函数高

## 基本类型与运算符

### 布尔类型 bool

定义 `var b bool = true`

* 运算符
  * 相等 `==`
  * 不相等 `!=`
  * 非 `!`
  * 与 `&&`
  * 或 `||`

### 数字类型

#### 整型 int

> 整形的零值为： 0
> 
> `int` `uint` 在32位系统上，使用32位（4个字节）,在64位系统上，使用64位（8个字节）
> 
> int 是计算最快的一种类型


* int 整数
  * int8  (-128 -> 127)
  * int16 (-32768 -> 32767)
  * int32 (-2,147,483,648 -> 2,147,483,647)
  * int64 (-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807)
* uint 无符号整数
  * uint8  (0 -> 255)
  * uint16 (0 -> 65,535)
  * uint32 (0 -> 4,294,967,295)
  * uint64 (0 -> 18,446,744,073,709,551,615)
  * uintptr
    * 长度被设定为足够存放一个指针即可

#### 浮点型 float

> 浮点型的零值为： 0.0
> 
> IEEE-754 标准
> 
> `float32` `float64` 精度不同，使用 `==` `!=` 比较时，需要注意

* float32 (± 1e-45 -> ± 3.4 * 1e38)
  * 精确到小数点后7位
* float64 (± 5 * 1e-324 -> 107 * 1e308)
  * 精确到小数点后15位

### 运算符与优先级

| 优先级 | 运算符                            |
|-----|--------------------------------|
| 7   | `^` `!`                        |
| 6   | `*` `/` `%` `<<` `>>` `&` `&^` |
| 5   | `+` `-` `\|` `^`               |
| 4   | `==` `!=` `<` `>` `<=`  `>=`   |
| 3   | `<-`                           |
| 2   | `&&`                           |
| 1   | `\|\|`                         |

### 字符类型

https://learnku.com/docs/the-way-to-go/basic-types-and-operators/3586#5cad4b

> 严格来说，这并不是 Go 语言的一个类型，字符只是整数的特殊用例

* byte
  * `uint8` 的别名
  * 用于表示 `ascii` 字符
  * 能够表示 0~255 范围内的 `ascii` 字符
  * 声明使用单引号 `'`
  * 写法
    * `\x` 紧跟长度为2的16进制数 如：`\x41`
    * `\ ` 紧跟长度为3的8进制数  如：`\377`
* rune
  * `int32` 的别名
  * 大小 4 个字节
  * 表示所有 `unicode` 字符
  * 声明使用单引号 `'`
  * 写法
    * `\u` `\U` 紧跟16进制数
      * `\u0041`


> `fmt.Printf()`格式化说明
> 
> * `%c` 表示字符
> * `%v` `%d` 输出表示字符的整数
> * `%U` 输出格式为 `U+hhhh` 的字符串


> unicode 包有一些测试字符串的函数
> * `unicode.IsLetter(ch)`  是否字母
> * `unicode.IsDigit(ch)`  是否数字
> * `unicode.IsSpace(ch)`  是否空白


## 字符串 string

> string的零值为： `""`

字符串表示
* 解释字符串
  * 用双引号包围 `""`
  * 转义字符
    * 换行符 `\n`
    * 回车符 `\r`
    * tab `\t`
    * Unicode 字符 `\u`  `\U`
    * 反斜杠自身  `\\`
* 非解释字符串
  * 该类字符串使用反引号 `` ` `` 括起来，支持换行

### 字符串比较

一般的比较运算符（`==`、`!=`、`<`、`<=`、`>=`、`>`），通过在内存中按字节进行比较


### 字符串截取

> **需要注意的是，这种转换方案只对纯 ASCII 码的字符串有效。**
> 
> 字符串内容可以通过索引进行获取，索引从 0 开始计数

```
s[0] // 第一个字节
s[len(s) -1] // 最后一个字节
```


### 字符串拼接

用 `+` 进行字符串拼接

加号拼接并不是最高效的做法，可选用 `strings` 包

```
s := "aaa" + "bbb"
s += "dddd"
```

### strings 包

用于字符串处理相关

> strings.HasPrefix(s, prefix string) bool
> 
> 是否以字符串开头
---
> strings.HasSuffix(s, suffix string) bool
>
> 是否以字符串结尾
---
> strings.Contains(s, substr string) bool
> 
> 是否包含字符串
---
> strings.Index(s, str string) int
> 获取字符串索引， -1 表示不包含
> 
> strings.LastIndex(s, str string) int
> 获取字符串最后的索引， -1 表示不包含
---
> strings.Replace(str, old, new string, n int) string
> 
> 替换字符串，返回新字符串， n=-1 表示替换所有字符串

### strconv 包

用于字符串相关的类型转换

文档：https://pkg.go.dev/strconv

## 日期和时间

`time` 包

`t := time.Now()`

* t.Day()
* t.Year()

### 格式化

这里很大不同，不是以常见的 `yyyy-MM-dd`

而是用 golang 诞生的时间 `2006-01-02 15:04:05 -0700 MST` 作为格式化

记忆技巧：

`月/日/时/分/秒/年/时区`

`1/2/3/4/5/6/7`

`01/02/03:04:05PM/06/-0700`

```
// 示例
fmt.Println("yyyy-MM-dd hh:mm:ss ->" + t.Format("2006-01-02 15:04:05"))
fmt.Println("yyyyMMdd ->", t.Format("20060102"))
fmt.Println("yyMMdd ->", t.Format("060102"))
fmt.Println("hh:mm:ss ->", t.Format("15:04:05"))
fmt.Println("hh:mm:ss.ffffff ->", t.Format("15:04:05.000000"))
```

## 指针

参考：https://go.timpaik.top/04.9.html



