package main

import (
	"bufio"
	"demo/goroutine"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open(("./menu.txt"))
	io.Copy(w, f)
}
func main() {

	//http.HandleFunc("/", Handler)
	//http.ListenAndServe(":3000", nil)

	//menu.AddItem()
	//Panic1()
	//Loops()
	//TheStruct()
	//Map()
	//Slice()
	//Array()
	//generics.GenericsMain()
	//interfaces.InterfaceMain()

	goroutine.TestGoroutine()
	var a string
	a = "foo"

	fmt.Println("Hello Go")
	fmt.Println(a)

	var b int = 99
	fmt.Println(b)

	var c = true
	fmt.Println(c)

	d := 3.1415
	fmt.Println(d)

	j, k := 10, 5
	l := j + k
	fmt.Println(l)

	//Type conversion
	var e int = int(d)
	fmt.Println(e)

	ab := "name"
	//& is an address operator
	//ac is a pointer
	//*ac is dereferencing a pointer
	ac := &ab
	*ac = "bisi"
	fmt.Println(ab)
	fmt.Println(*ac)

	dc := new(int)
	//ac = new(string)
	fmt.Println(dc)
	fmt.Println(ac)
	fmt.Println(&ab)

	fmt.Println("What is your name? ")
	nameReadeer := bufio.NewReader(os.Stdin)
	name, _ := nameReadeer.ReadString('\n')
	name = strings.TrimSpace(name)
	name = strings.ToUpper(name)
	fmt.Println(name + "!")

}

func Array() {
	var arr [3]int
	fmt.Println(arr)

	arr = [3]int{1, 2, 3}
	fmt.Println(arr)

	arr[2] = 4
	fmt.Println(arr)

}

// reference data types
func Slice() {
	var s []int
	fmt.Println(s) // [] (nil)

	s = []int{1, 2, 3}
	fmt.Println(s)

	// updating s
	s[1] = 7
	fmt.Println(s)

	v := s
	sa := &s
	va := &v
	fmt.Println(v)
	v[0] = 24
	fmt.Println(s)
	fmt.Println(sa)
	fmt.Println(v)
	fmt.Println(va)

	//apppend to a slice
	s = append(s, 9, 10, 11)
	saa := &s
	fmt.Println(saa)
	fmt.Println(va)

	//delete from a slice
	s = slices.Delete(s, 1, 3) //remove indices 1 and 2 from slice
	fmt.Println(s)

}

func Map() {
	var s map[string]int
	fmt.Println(s) // [] (nil)
	s = map[string]int{"bar": 1, "foo": 2}
	v := s
	v["bar"] = 5
	fmt.Println(s)
	fmt.Println(v)

	ss := s["bar"]
	fmt.Println(ss)

	sss, status := s["ba"]
	fmt.Println(sss, status)
	fmt.Println(status)

	s["ba"] = 45
	sss = s["ba"]
	fmt.Println(sss)
	delete(s, "ba")

	anothermap := map[string][]string{
		"coffee": {"coffee", "espresso", "cappuccino"},
		"tea":    {"black tea", "green tea", "lipton"},
	}
	fmt.Println(anothermap)

	anothermap["other"] = []string{"beans"}
	fmt.Println(anothermap)

}

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

type menuItem struct {
	name   string
	prices map[string]float64
}

func Loops() {
	// for { ......} infinite loop
	// for condition { ......}  loop till condition
	// for initialozer; test; post clause { ......}  counter-based loop

	//counter-based loop
	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("done")

	//condition loop
	j := 1
	for j < 3 {
		fmt.Println(j)
		j++
	}

	//infinite loop
	i := 1
	for {
		fmt.Println(i)
		i++
		break
	}

	//looping with collections
	// for key, value := range collection {......} get both key and value
	// for key := range collection {......} get just the key
	// for _,value := range collection {......} get just the value

	arr := [3]int{1, 2, 3}
	for b, v := range arr {
		fmt.Println(b, v)
	}

	menu := []menuItem{
		{name: "Cofee", prices: map[string]float64{"small": 1.65, "medium": 1.80}},
		{name: "Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25}},
	}
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}

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

// Panic section

func Panic1() {
	fmt.Println("main 1")
	func1()
	fmt.Println("main 2")

}

func func1() {
	fmt.Println("func 1")
	panic("oh-oh")
	fmt.Println("func 2")

}

// Panic and Recover section

func Panic2() {
	fmt.Println("main 1")
	func2()
	fmt.Println("main 2")

}

func func2() {
	defer func() {
		fmt.Println(recover())
	}()
	fmt.Println("func 1")
	panic("oh-oh")
	fmt.Println("func 2")

}

func action() {
	dividend, divisor := 10, 5
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))

	dividend, divisor = 10, 0
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))

}

func divide(dividend, divisor int) int {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()
	return divisor / dividend
}

func myGotoFunc() {
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
