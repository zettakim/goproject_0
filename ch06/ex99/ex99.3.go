package main

import "fmt"

func main() {
	var x int8 = 1
	fmt.Printf("a -> %4d : %+09b\n", x, x)

	x <<= 7
	fmt.Printf("a -> %4d : %09b\n", x, x)

	x >>= 7
	fmt.Printf("a -> %4d : %09b\n", x, x)
}
