package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	const H = "Hello"
	const W = "World"

	var buf = make([]byte, len(W)+len(H))
	copy(buf, H)
	copy(buf[len(H):], W)
	var vstr = string(buf)

	fmt.Println(vstr)
	{
		a := (*reflect.StringHeader)(unsafe.Pointer(&vstr))
		bts := unsafe.Slice((*byte)(unsafe.Pointer(a.Data)), 5)
		bts[0] = byte('A')
	}
	fmt.Println(vstr)
}

func help(a string, b string) {
}
