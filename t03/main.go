package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	const H = "Hello"
	const W = "World"
	//
	help(H, W)
	//
	var vstr = H + ", " + W
	fmt.Println(vstr)
	{
		a := (*reflect.StringHeader)(unsafe.Pointer(&vstr))
		bts := unsafe.Slice((*byte)(unsafe.Pointer(a.Data)), 5)
		bts[0] = byte('A')
	}
	fmt.Println(vstr)
}
func help(a string, b string) {
	var vstr = a + ", " + b
	fmt.Println(vstr)
	{
		a := (*reflect.StringHeader)(unsafe.Pointer(&vstr))
		bts := unsafe.Slice((*byte)(unsafe.Pointer(a.Data)), 5)
		bts[0] = byte('A')
	}
	fmt.Println(vstr)
}
