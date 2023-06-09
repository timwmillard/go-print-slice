package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s1 []int
	var s2 = []int{}       // s2 is not nil, it's a slice with zero length
	s3 := make([]int, 0)   // s3 is not nil, it's a slice with zero length
	fmt.Println(s1 == nil) // true
	fmt.Println(s2 == nil) // false
	fmt.Println(s3 == nil) // false
	fmt.Println(len(s1))   // 0
	fmt.Println(len(s2))   // 0
	fmt.Println(len(s3))   // 0
	printSlice(s1, "[]int")
	printSlice(s2, "[]int{}")
	printSlice(s3, "make([]int, 0)")

	s4 := []int{99, 98, 97, 96, 95}

	printSlice(s4, "[]int{99, 98, 97, 96, 95}")

	// s5 := []int{100, 200, 300}
	// printSlicePtr(&s5, "&[]int{100, 200, 300}")

	// s5 := make([]int, 0, 3)
	// printSlice(s5, "make([]int, 0, 3)")

	empty := struct{}{}
	fmt.Printf("Address of struct{}{} = %p\n", unsafe.Pointer(&empty))
}
