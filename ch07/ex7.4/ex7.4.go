package main

import "fmt"

func Divide(a int, b int) (int, bool) {
	if b == 0 {
		return 0, false
	} else {
		return a / b, true
	}
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	c, success = Divide(9, 0)
	fmt.Println(c, success)
}
