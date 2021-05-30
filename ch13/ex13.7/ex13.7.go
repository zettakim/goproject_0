// Memory Padding Exam

package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	a int8 // 1 byte
	b int
	c int8
	d int
	e int8
}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}
