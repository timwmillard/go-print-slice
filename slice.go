package main

import (
	"fmt"
	"os"
	"reflect"
	"unsafe"

	"github.com/olekukonko/tablewriter"
)

func printSlicePtr(s *[]int, description string) {
	if s != nil {
		printSliceInfo(*s, description)
	}
	printPtr(s)
	if s != nil {
		printSliceHeader(*s)
		printSliceData(*s)
	}
	printSliceEnd()
}

func printSlice(s []int, description string) {
	printSliceInfo(s, description)
	printSliceHeader(s)
	printSliceData(s)
	printSliceEnd()
}

func printSliceInfo(s []int, description string) {
	fmt.Printf("** SLICE ****************************************\n")
	fmt.Printf("  %s\n", description)
	fmt.Printf("  len = %d\n", len(s))
	fmt.Printf("  cap = %d\n", cap(s))
}

func printSliceEnd() {
	fmt.Printf("*************************************************\n")
	fmt.Println()
}

func printPtr(s *[]int) {

	printSliceHeader(*s)

	empty := struct{}{}
	emptyPtr := (*uintptr)(unsafe.Pointer(&empty))
	fmt.Printf("Pointer\n")
	fmt.Printf("Pointer size = %v\n", unsafe.Sizeof(s))
	fmt.Printf("struct{}{} address = %p\n", emptyPtr)

	ptr := (*uintptr)(unsafe.Pointer(s))

	data := [][]string{
		{
			"ptr",
			fmt.Sprintf("0x%x", &s),
			fmt.Sprintf("0x%x", ptr),
			fmt.Sprintf("0x%x", *ptr),
			fmt.Sprintf("0x%x", &ptr),
			fmt.Sprintf("0x%x", unsafe.Sizeof(s)),
		},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Address", "Value"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

func printSliceHeader(s []int) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("Header\n")

	data := [][]string{
		{
			"ptr",
			fmt.Sprintf("0x%x", &sh.Data),
			fmt.Sprintf("0x%x", sh.Data),
		},
		{
			"len",
			fmt.Sprintf("0x%x", &sh.Len),
			fmt.Sprintf("%d", sh.Len),
		},
		{
			"cap",
			fmt.Sprintf("0x%x", &sh.Cap),
			fmt.Sprintf("%d", sh.Cap),
		},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Address", "Value"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

func printSliceData(s []int) {
	fmt.Printf("Data\n")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Address", "Value"})

	for i, v := range s {
		row := []string{
			fmt.Sprintf("[%d]", i),
			fmt.Sprintf("0x%x", &s[i]),
			fmt.Sprintf("%d", v),
		}
		table.Append(row)
	}
	table.Render() // Send output
}
