package main

import (
	"bufio"
	"demo/errors"
	"demo/gofunctions"

	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open(("./menu.txt"))
	io.Copy(w, f)
}
func main() {

	//types.TypesMain()
	gofunctions.FunctionMain()

	return
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

	// goroutine.TestGoroutine()
	//goroutine.SelectWithTwoChannelsExample3()
	//goroutine.LoopingOverChannel()

	errors.ErrorHandling()
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

func Add(l, r int) int {
	return l + r
}
