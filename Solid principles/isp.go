package main

import "fmt"
// Interface segregation principle
// “Many client specific interfaces are better than one general purpose interface” 
type Details interface{
	getName() string
	//getCard()		//violates isp becaz every student wont have id so we need to segregate interface
}
type FullDetails inetrface{
	Details
	getCard() int
}

type stuWithoutId struct{
	name string
}
type stuWithId struct{
	name string
	cardNo int
}
func (s stuWithoutId)getName() string{
	return s.name
}
func (s stuWithId)getName() string{
	return s.name
}
func (s stuWithId)getCard() int{
	return s.cardNo
}

func main(){
	a:=stuWithoutId{name:"neha"}
	a.getName()
	b:=stuWithId{name:"pooja",cardNo:2}
	b.getName()
	b.getCard()
}