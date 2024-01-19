// go run go_pointer.go

package main

import "fmt"

func printMap(ma map[string]int) {
	fmt.Println(len(ma), ma)
	for k, v := range ma {
		fmt.Println("key = ", k, "value = ", v)
	}
}

func main() {
	ma := make(map[string]int)
	ma["a"] = 1
	ma["b"] = 2
	ma["c"] = 3
	printMap(ma)

	fmt.Println(ma["x"])
	v, ok := ma["x"]
	if ok {
		v += 1
	} else {
		v = 1
	}
	ma["x"] = v
	printMap(ma)
}
