package main

import "fmt"

type Cal interface {
	Jisuan() int
}
type Add struct {
	num1 int
	num2 int
}

func (a *Add) Jisuan() (value int) {
	value = a.num1 + a.num2
	return
}

type Sub struct {
	num1 int
	num2 int
}

func (s *Sub) Jisuan() (value int) {
	value = s.num1 - s.num2
	return
}

type Factory struct {
}

func Duotai(c Cal) (value int){
	value=c.Jisuan()
	return
}
func (f *Factory) Oprate(num1, num2 int, sig string) (value int) {
	var c Cal
	switch sig {
	case "+":
		Add := Add{num1, num2}
		//value=Add.Jisuan()
		c=&Add
	case "-":
		Sub := Sub{num1, num2}
		//value=Sub.Jisuan()
		c=&Sub
	}
	value=Duotai(c)
	return value
}
func main() {
	var f Factory
	value := f.Oprate(10, 20, "-")
	fmt.Println(value)
}
