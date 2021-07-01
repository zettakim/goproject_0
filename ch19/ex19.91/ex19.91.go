package main

import "fmt"

type Cart struct {
	productList string
}

func (c *Cart) AddPoduct(product string) {
	if c.productList != "" {
		c.productList += ". "
	}
	c.productList += product
}

func (c *Cart) Clear() {
	c.productList = ""
}

func (c Cart) GetProductList() string {
	return c.productList
}

func main() {
	c := &Cart{}
	c.AddPoduct("apple")
	c.AddPoduct("kimche")

	fmt.Println(c.GetProductList())

	c.Clear()
	c.AddPoduct("watermelon")
	fmt.Println(c.GetProductList())
}
