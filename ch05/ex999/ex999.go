package main

import "fmt"

func main() {
	var a = 345
	var b = 3.1415

	var c int
	var d int

	fmt.Printf("%05d\n", a)
	fmt.Printf("%5.2f\n", b)

	fmt.Scanln(&c, &d)
	fmt.Println(c, d)

	c = 123
	d = 4567
	f := 3.14159269
	fmt.Printf("%6d\n", c)
	fmt.Printf("%06d\n", d)
	fmt.Printf("%6.2f\n", f)
}
