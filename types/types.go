package types

import "fmt"

func TheStruct() {
	theStruct := myStruct{id: 4, name: "bola"}
	fmt.Println(theStruct.name)
	theStruct.name = "bose"
	fmt.Println(theStruct.name)

}

// struct is copied by value
// it is comparable
type myStruct struct {
	id   int
	name string
}

type MenuItem struct {
	Name   string
	Prices map[string]float64
}
