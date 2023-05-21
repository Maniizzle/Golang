package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

var in = bufio.NewReader(os.Stdin)

func AddItem() {
	name, _ := in.ReadString('\n')
	menu = append(menu, menuItem{name: name, prices: map[string]float64{"single": 1.69, "double": 3.2}})
}

func PrintMenu() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, cost := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, cost)

		}
	}
	name, _ := in.ReadString('\n')
	menu = append(menu, menuItem{name: name, prices: map[string]float64{"single": 1.69, "double": 3.2}})
}
