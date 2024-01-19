
# method
a method is just a function with a receiver argument.
```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

# receiver
pointer as receiver
指针作为method的receiver，如果在method中修改了指针对应的值，原来的值也会修改。（传引用）
value作为method的receiver，会复制一份出来，原来的值并不会修改。（传值）

指针作为receiver的好处：
1. 可以直接修改receiver的值
2. 减少不必要的拷贝