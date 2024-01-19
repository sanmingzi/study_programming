
# type parameters

泛型函数
参数可以是任意类型，只需要实现某些方法即可。有点类似c里面的sort。
```go
func Index[T comparable](s []T, x T) int
```

# Generic types

泛型类型

```go
type List[T any] struct {
	next *List[T]
	val  T
}
```