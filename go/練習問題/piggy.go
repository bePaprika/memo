package main

import (
	"fmt"
	"math/rand"
)

func main() {

	all := 0
	dollar := 0
	cent := 0
	coin := 0
	fmt.Printf("Savings balance is $%d.%02d\n", dollar, cent)

	for i := 0; i < 20; i++ {
		var randnum = rand.Intn(3)
		switch randnum {
		case 0:
			coin = 5
		case 1:
			coin = 10
		case 2:
			coin = 25
		}

		all += coin
		dollar = all / 100
		cent = all % 100
		fmt.Printf("Savings balance is $%d.%02d (+%d)\n", dollar, cent, coin)
	}
}