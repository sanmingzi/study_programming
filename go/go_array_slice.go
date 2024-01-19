// go run go_array_slice.go

package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	printArray(a)

	for i, v := range a {
		fmt.Println(i, v)
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	b := [...]int{1: 1, 2: 2}
	printArray(b)

	// slice
	var array = [...]int{0, 1, 2, 3, 4, 5, 6}
	s1 := array[0:4]
	s2 := array[:4]
	s3 := array[2:]
	s4 := array[:]
	array[2] = -2
	printSlice(s1)
	printSlice(s2)
	printSlice(s3)
	printSlice(s4)

	s4 = append(s4, -1) // 此时的 s4 和 array 已经没有关系了
	array[2] = 2
	printSlice(s4)
	fmt.Println("------")

	testSlice()
}

func printArray(a [3]int) {
	fmt.Println(len(a), cap(a), a)
}

func printSlice(s []int) {
	fmt.Println(len(s), cap(s), s)
}

func testSlice() {
	a := make([]int, 10)
	printSlice(a)
	a = append(a, 1)
	printSlice(a)

	b := make([]int, 10, 15)
	printSlice(b)
	b = append(b, 1)
	printSlice(b)
}
