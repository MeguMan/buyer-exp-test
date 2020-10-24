package model

import "fmt"

type Ad struct {
	ID    int
	Link  string
	Price int
}

func (ad *Ad) ParsePrice(link string) int {
	fmt.Println(link + "AAAAAAAAAAA")
	return 123
}
