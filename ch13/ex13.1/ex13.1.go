package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var house House
	house.Address = "서울 용인시..."
	house.Size = 45
	house.Price = 10
	house.Type = "Apt"

	fmt.Println("주소: ", house.Address)
	fmt.Println("크기: ", house.Size)
	fmt.Println("가격: ", house.Price)
	fmt.Println("형태: ", house.Type)
}
