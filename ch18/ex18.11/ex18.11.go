package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	slice = append(slice, 0)

	idx := 2
	// fmt.Println("slice - 2 = ", len(slice)-2)
	// for i := len(slice) - 2; i >= idx; i-- {
	// 	slice[i+1] = slice[i]
	// }

	copy(slice[idx+1:], slice[idx:])
	slice[idx] = 100
	fmt.Println(slice)
}
