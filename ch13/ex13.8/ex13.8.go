// Memory Padding Exam, 메모리 절약

package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	a int8 // 1 byte
	c int8
	e int8
	b int
	d int
}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}
