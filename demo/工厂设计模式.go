package main

import (
	"fmt"
)

type Oper interface {
	Jisuan() int
}
type Add struct {
	num1 int
	num2 int
}

func (a *Add) Jisuan() int {
	value := a.num1 + a.num2
	return value
}

type Sub struct {
	num1 int
	num2 int
}

func (s *Sub) Jisuan() int {
	value := s.num1 - s.num2
	return value
}
func Duotai(o Oper) int {
	value := o.Jisuan()
	return value
}

type Factory struct {
}

func (f *Factory) Operation(num1, num2 int, sig string) int {
	var o Oper
	switch sig {
	case "+":
		Add := Add{num1, num2}
		o = &Add
	case "-":
		Sub := Sub{num1, num2}
		o = &Sub
	}
	value := Duotai(o)
	return value
}
func main() {
	var f Factory
	value := f.Operation(10, 20, "+")
	fmt.Println(value)
}
