package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var  a,b = "hello world",0
	const W,H = 12,4
	//fmt.Println("谢谢你")
	var area = W*H
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(area)
	fmt.Println(unsafe.Sizeof(a))
}
