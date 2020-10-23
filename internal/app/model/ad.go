package model

import "fmt"

type Ad struct {
	ID int
	Link string
	Price int
}

func (ad *Ad) CheckPrice() {
	fmt.Println("IM CHECKING PRICES")
}
