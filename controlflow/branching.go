package controlflow

import "fmt"

func BranchingWithIf() {

	i := 5
	if i < 5 {
		fmt.Println("i is less than 5")
	} else {
		fmt.Println("i is AT LEAST  5")
	}
	fmt.Println("After the if else statement")

	j := 11
	if j < 5 {
		fmt.Println("j is less than 5")
	} else if j < 10 {
		fmt.Println("j is less than  10")
	} else {
		fmt.Println("j is at least  10")

	}
	fmt.Println("After the if else statement")

	//if with the initializer
	if k := 4; k < 5 {
		fmt.Println("j is less than 5")
	} else if k < 10 {
		fmt.Println("j is less than  10")
	} else {
		fmt.Println("j is at least  10")

	}
	fmt.Println("After the if else statement")

}
func BranchingWithSwitch() {

	i := 5
	switch i {
	case 1:
		fmt.Println("first case")
	case 2 + 3, 2*i + 3:
		fmt.Println("second case")
	default:
		fmt.Println("default case")

	}
	//switch with initializer
	switch j := 5; j {
	case 1:
		fmt.Println("first case")
	case 2 + 3, 2*j + 3:
		fmt.Println("second case")
	default:
		fmt.Println("default case")

	}

	//logical switch : true is defautly implied after the initializer
	switch m := 8; {
	case m < 5:
		fmt.Println("m is less than 5")
	default:
		fmt.Println("m is greater than 10")

	}
}

func DeferredFunction() {
	//last deferred function is called first i.e defer 2 will be called first
	fmt.Println("main 1")

	defer fmt.Println("defer 1")

	fmt.Println("main 2")

	defer fmt.Println("defer 2")

	//output
	//main 2
	//defer 2
	//defer 1

}

func GotoFunc() {
	i := 10
	if i < 15 {
		goto myLabel
	}
myLabel:
	j := 42
	for ; i < 15; i++ {
		fmt.Println(i, j)
	}
}
