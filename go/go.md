# go

## 核心问题

变量：作用域，生命周期

上下文：context

闭包：
function是闭包
if是闭包吗？for循环算闭包吗？

对象体系：
常量/变量
function/method
传值/传引用

## 参考链接
https://go.dev/tour/welcome/1

https://gobyexample.com/

https://go.dev/blog/declaration-syntax

## 目录

- if else
- for (while)
- array / slice
- map
- function
- method
- interface
- generic type
- concurrency

## 目录2

- 基础知识
	- identifier 标识符
	- pointer 指针
	- array 数组
	- slice 切片
	- map
	- struct
	- 控制结构
		- if
		- switch
		- for
		- 标签和跳转？
- 函数
	- 基础
		- 多值返回
		- 实参 / 形参
		- 不定参数
	- 函数签名 / 匿名函数
	- defer
	- 闭包
	- panic / recover
	- 错误处理
	- 底层实现
		- 函数调用规约
		- 汇编基础
		- 多值返回分析
		- 闭包底层实现
- 类型系统
	- 类型简介
	  - 命名类型 / 未命名类型
		- 底层类型
		- 类型相同和类型赋值
		- 类型强制转换
	- 类型方法
		- 自定义类型
		- 方法
	- 方法调用
		- 一般调用
		- 方法值
		- 方法表达式
		- 方法集
		- 值调用和表达式调用的方法集
	- 组合和方法集
		- 组合
		- 组合的方法集
	- 函数类型
- 接口
	- 基本
		- 接口声明
		- 接口初始化
		- 接口方法调用
		- 接口的动态类型和静态类型
	- 接口运算
		- 类型断言(type assertion)
		- 类型查询(type switch)
		- 接口优点和使用形式
	- 空接口
		- 空接口的用途
		- 空接口和nil
	- 接口内部实现
		- 数据结构
		- 接口调用过程
		- 接口调用代价
		- 空接口数据结构
- 并发
- 反射
- 语言陷阱
  - 多值赋值
	- 短变量声明和赋值
	- range复用临时变量
	- defer陷阱
	- 切片困惑
	- 值，指针，引用
		- 传值 / 传引用
		- 函数名的意义
		- 引用语法
	- 习惯用法
		- 干净与强迫症
		- comma，ok表达式
		- 简写模式
		- 包中的函数或方法设计
		- 多值返回函数
- 编程哲学

pointer

struct

range
可以返回数组对应的index以及value

# type assertions
```go
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) // panic
	fmt.Println(f)
}
```

# type switches
```go
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```

# select and default select
