package main

import "fmt"

func main() {
	var a int8

	fmt.Printf("a      -> %4d : %08b\n", a, a)
	a |= 2
	fmt.Printf("a |= 2 -> %4d : %08b\n", a, a)
	a |= 4
	fmt.Printf("a |= 4 -> %4d : %08b\n", a, a)
	a |= 8
	fmt.Printf("a |= 8 -> %4d : %08b\n\n", a, a)

	var b int8
	b = 4
	fmt.Printf("a      -> %4d : %08b\n", a, a)
	fmt.Printf("b      -> %4d : %08b\n", b, b)

	a &^= b
	fmt.Printf("a &^= b > %4d : %08b\n", a, a)
}
