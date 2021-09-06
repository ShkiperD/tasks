package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var Number = 42
	var CheckingMin = 0
	var CheckingMax = 100
	var Checking int
	fmt.Println("Загадайте число: ")
	for{
		Checking = Random(CheckingMin,CheckingMax)
		break
	}
}
func Random(checkMn, checkMx int) int {
	rand.Seed(time.Now().UnixNano())
	n:=rand.Intn(checkMn,checkMx)
	return n
}