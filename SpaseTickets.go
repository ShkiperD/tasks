package main

import (
	"fmt"
	"math/rand"
)
const way = 62100000
const secOnDay = 86400

func main() {
	fmt.Printf("%-18v %-4v %-4v %6v\n", "Spaceline","Days","Trip type","Price")
	fmt.Println("========================================")
	for count:=0;count<10;count++{
		speed,price:=daysOfAdv()
		days:=way/(speed*secOnDay)
		tripType:=tripTypeGen()
		if tripType=="Round Trip"{price*=2}
		fmt.Printf("%-19v %-3v %-12v %v\n",stationPick(),days,tripType,price)
	}
}
func stationPick() string{
	station:=rand.Intn(3)
	var nameStation string
	switch station {
		case 0: nameStation = "SpaceX"
		case 1:	nameStation = "VirginGalactic"
		case 2: nameStation = "Space Adventure"
	}
	return nameStation
}
func daysOfAdv() (int,int) {
	speedShip:= 16 + rand.Intn(14)
	price:=36+(speedShip-16)
	return speedShip,price
}
func tripTypeGen() string{
	var tripType string
	n:=rand.Intn(2)
	switch n{
	case 0:
		tripType="Round Trip"
	case 1:
		tripType="One-way"
	}
	return tripType
}