# day2

## 控制结构

### if-else

```
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
```

> 注意
> 
> 关键字 if 和 else 之后的左大括号 { 必须和关键字在同一行
> 
> 如果你使用了 else-if 结构，则前段代码块的右大括号 } 必须和 else-if 关键字在同一行

```
// 非法的
if x{
}
else {	// 无效的
}
```

if可以包含初始化语句

格式：
```
if initialization; condition {
  // do something
}
```
```
if var1 := GetValue(); var1 == 1 {
  // do something
}
```

### 多返回值函数的错误

* 处理错误
```
orinStr := "abc"
rs, err := strconv.Atoi(orinStr)

// 习惯用法
if err != nil {
    fmt.Println("err:", err)          // 处理错误或日志
    return err                        // return err
}
fmt.Println(rs)
```

* 忽略错误

```
orinStr := "abc"
rs, _ := strconv.Atoi(orinStr)
fmt.Println(rs) // rs是int，将打印零值：0
```

输出
```
0
```
### switch 结构

参考：https://go.timpaik.top/05.3.html

```
switch var1 {
	case val1:
		...
	case val2:
		...
	default:
		...
}
```

* 与其他语言类似
* 从上到下逐一执行比对，知道进入某个 case 或 default
* 一旦成功地匹配到某个分支，在执行完相应代码后就会退出整个 switch 代码块
  * 不用特别的 `break`
* 如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，可以使用 `fallthrough` 关键字来达到目的


```
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
```

输出

```
3333
4444
5555
default case
```


### for 结构

#### 基于计数器

```
for 初始化语句; 条件语句; 修饰语句 {}
```
示例
```
for i := 0; i < 10; i++ {
    println(i)
}
```

#### 基于条件

类似 while

```
for 条件语句 {}
```
示例
```
i := 5
for i > 0 {
    i = i - 1
    println(i)
}

```

输出
```
4
3
2
1
0
```

#### 无限循环

*  如果 for 循环的头部没有条件语句，那么就会认为条件永远为 true，因此循环体内必须有相关的条件判断以确保会在某个时刻退出循环。
  * `break`
  * `return`
* 一般省略条件，可直接写作 `for { }`

示例

```
i := 5
for {
    println(i)
    i = i - 1
    if i <= 0 {
        return
    }
}
```
比较经典的，不断等待和接受新的请求
```
for t, err = p.Token(); err == nil; t, err = p.Token() {
	// ...
}
```

#### for-range 结构

go特有的迭代结构，可以迭代任何一个集合（包括数组和 map）

一般写法
```
for ix, val := range coll { }
```

* 要注意的是，val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值
* 如果 val 为指针，则会产生指针的拷贝，依旧可以修改集合中的原值

```
str := "abc 123"
for i, s := range str {
    fmt.Println(i, string(s))
}
```
输出
```
0 a
1 b
2 c
3  
4 1
5 2
6 3
```

#### break 和 continue

* break
  * 用于跳出循环、switch、select
* continue
  * 只能用于 for 循环
  * 忽略剩余语句，直接进入下一次循环

#### LABEL 和 goto

用的比较少

https://go.timpaik.top/05.6.html

## 数组和切片

## map
