package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:2]

	fmt.Println(array)
	fmt.Println("slice:", slice, "len:", len(slice), "cap:", cap(slice))

	array[1] = 100

	fmt.Println("After change second element")
	fmt.Println(array)
	fmt.Println("slice:", slice, "len:", len(slice), "cap:", cap(slice))

	slice = append(slice, 500)
	fmt.Println("After 500 append")
	fmt.Println(array)
	fmt.Println("slice:", slice, "len:", len(slice), "cap:", cap(slice))
}
