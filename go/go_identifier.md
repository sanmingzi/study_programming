# identifier 标识符

## keywords

- 8个程序整体结构的关键字

package
import
const
var
func
defer
go
return

- 4个复合数据类型的关键字

struct
interface
map // map类型
chan // 通道

- 控制程序的13个关键字

if else
for range break continue
switch select type case default fallthrough
goto

## 内置数据类型

强类型静态编译语言，所以在定义常量和变量的时候需要指定数据类型，可以显示指定，也可以通过go的自动类型推导来确定。

整型（12个）
  byte int int8 int16 int32 int64
  uint unint8 uint16 uint32 uint64 uintprt
浮点型（2个）
  float32 float64
复数型（2个）
  complex64 complex128
字符串 (2个)
  string rune
接口（1个）
  error
布尔型（1个）
  bool

## 内置函数(15个)

make new len cap append copy delete panic recover close complex real image Print Printin

## 常量标识符

true false
iota // 枚举
nil

## 空白标识符

_
