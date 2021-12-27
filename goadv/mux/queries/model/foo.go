package model

type Foo struct{
	Fname string
	Lname string
}
func NewFoo(fname string,lname string)Foo{
	return Foo{
		Fname:fname,
		Lname:lname,
	}
}
