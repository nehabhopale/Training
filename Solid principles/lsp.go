
package main
//Liskov substitution

type human interface{
	getName() string
}

type human struct{
	name string
}
type teacher struct{
	human
	id int
}
func(h human) getName()string{
	return  h.name
}
func(t teacher) getName()string{
	return  "teacher"
}

func main(){
	teacher:=teacher{
		human:human{name:"neha"},
		id:1
	}
	//as per principle child should implement all the methods that are defined on parent struct
	fmt.Println(teacher.getName())
	fmt.Println(teacher.human.getName())

}