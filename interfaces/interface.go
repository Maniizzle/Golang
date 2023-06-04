package interfaces

import (
	"bytes"
	"fmt"
	"strings"
)

type printer interface {
	Print() string
}

type user struct {
	username string
	id       int
}

func (u user) Print() string {
	return fmt.Sprintf("%v [%v]\n", u.username, u.id)
}

type menuItem struct {
	name   string
	prices map[string]float64
}

func (m menuItem) Print() string {
	var b bytes.Buffer
	b.WriteString(m.name + "\n")
	b.WriteString(strings.Repeat("-", 10) + "\n")
	for size, cost := range m.prices {
		fmt.Fprintf(&b, "\t%10s%10.2f\n", size, cost)
	}

	return b.String()
}
func InterfaceMain() {
	var p printer
	p = user{username: "bola", id: 2}
	fmt.Println(p.Print())

	p = menuItem{name: "Coffee", prices: map[string]float64{"single": 1.54, "double": 2.3}}
	fmt.Println(p.Print())

	//type assertion

	//u := p.(user)
	//fmt.Println(u) //throws panic because p is currently holding a menutype
	u, ok := p.(user)
	if ok {
		fmt.Println(u)
	}
	d := p.(menuItem)
	fmt.Println(d)

	//type switches for type assertions
	switch v := p.(type) {
	case user:
		fmt.Println("found a user!", v)

	case menuItem:
		fmt.Println("found a menu item!", v)

	default:
		fmt.Println("I am not found")
	}

}
