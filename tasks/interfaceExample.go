package main

/*
Create the Calculator interface with a couple of sub-interfaces
*/

import (
	"fmt"
	"reflect"
)

type Calculator interface {
	Adder
	Multer
	AdderMulter // Interface can inherit a methods with the same name, but they must have the same API
	sub(float32, float32) float32
	div(float64, float64) float64
}

type Adder interface {
	add(float64, float64) float64
}

type Multer interface {
	mul(float64, float64) float64
}

type AdderMulter interface {
	add(float64, float64) float64
	mul(float64, float64) float64
}

type myCalc struct{}

func (m myCalc) add(a, b float64) float64 {
	return a + b
}

func (m myCalc) sub(a, b float32) float32 {
	return a - b
}

func (m myCalc) mul(a, b float64) float64 {
	return a * b
}

func (m myCalc) div(a, b float64) float64 {
	return a / b
}

func main() {
	c := new(myCalc)
	interfaceType := reflect.TypeOf((*Calculator)(nil)).Elem()
	fmt.Println(reflect.TypeOf(c).Implements(interfaceType))

	fmt.Println(c.add(12.0, 15.0))
	fmt.Println(c.sub(12.0, 15.0))
	fmt.Println(c.mul(12.0, 15.0))
	fmt.Println(c.div(12.0, 15.0))
}
