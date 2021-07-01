package main

import "time"

// 택배회사
type Courier struct {
	Name string
}

// 물품
type Product struct {
	Name  string
	Price int
	ID    int
}

// 소포
type Parcel struct {
	Pdt           *Product
	ShippedTime   time.Time
	DeliveredTime time.Time
}

func (c *Courier) SendProduct(pdt *Product) *Parcel {
	p := &Parcel{}
	p.Pdt = pdt
	p.ShippedTime = time.Now()
	return p
}

func (p *Parcel) Delivered() *Product {
	p.DeliveredTime = time.Now()
	return p.Pdt
}

func main() {

}
