package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create file")
		return
	}

	defer fmt.Println("default call")
	defer f.Close()
	defer fmt.Println("file closed")

	fmt.Println("add to file HelloWorld")
	fmt.Fprintln(f, "Hello, world")
}
