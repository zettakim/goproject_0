package main

import "fmt"

type ParkingLot struct {
	LotSize int
}

func (lot *ParkingLot) ParkCar(carSize int) {
	lot.LotSize -= carSize
}

func main() {
	lot := &ParkingLot{100}
	lot.ParkCar(10)
	fmt.Println(lot.LotSize)
}
