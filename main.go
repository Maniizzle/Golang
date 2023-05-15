package main

import (
	"bufio"
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

	TheStruct()
	Map()
	//Slice()
	//Array()

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

type myStruct struct {
	id   int
	name string
}
