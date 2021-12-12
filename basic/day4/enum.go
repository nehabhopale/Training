package main

import "fmt"

type Enum interface {
	name() string
	ordinal() int
}

type GenderType uint

const (
	MALE = iota
	FEMALE
)

var genderTypeStrings = []string{
	"MALE",
	"FEMALE",
}

func (gt GenderType) name() string {
	return genderTypeStrings[gt]
}

func (gt GenderType) ordinal() int {
	return int(gt)
}
func main() {
	var ds GenderType = MALE
	fmt.Printf("The Gender is %s\n", ds.name())
	fmt.Println(ds.ordinal())
}
