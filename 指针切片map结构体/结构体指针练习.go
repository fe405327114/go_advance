package main

import "fmt"

type Empolyee struct {
	name string
	id int
	famlily []string
	maryy bool
}
func main(){
	//var p1 Empolyee
	//p1:=new(Empolyee)
	//initFunc(p1)
	p:=initFunc2()
	fmt.Println(p)
}
func initFunc(p *Empolyee){
	//如果写成p=&Employee{}则表示的是将地址赋值给指针，没有意义
	*p=Empolyee{"Nami",21,[]string{"Joba","Bluke"},true}
	p.name="Luffr"
}
func initFunc2()*Empolyee{
	p1:=new(Empolyee)

	*p1=Empolyee{"Nami",21,[]string{"Joba","Bluke"},true}
	p1.name="Luffr"
	return p1  //不可以返回局部变量的地址，因为局部变量调用完后会释放，地址会重新分配
	//如果是用NEW函数创建的在堆区，则不影响
}
