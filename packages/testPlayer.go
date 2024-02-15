package main

import (
	"fmt"
	"go_experiments/packages/player"
)

func main() {
	testStruct := new(player.Player)
	testStruct.On = true
	testStruct.Ammo = 1
	testStruct.Power = 2
	fmt.Println(testStruct.Shoot())
	fmt.Println(testStruct.Shoot())
	fmt.Println(testStruct.RideBike())
	fmt.Println(testStruct.RideBike())
	fmt.Println(testStruct.RideBike())
	//testStruct.do(&testStruct.Ammo) // Can't do this, the method is private and available only inside a package
}
