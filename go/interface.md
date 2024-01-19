

# interface

struct 通过实现接口中定义的方法，来实现接口的继承。

receiver is nil for interface method

# Stringer
类似ruby.to_s方法，自定义struct可以通过实现 String()
```go
type Stringer interface {
    String() string
}
```

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
```
