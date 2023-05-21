package controlflow

import (
	"demo/types"
	"fmt"
	"strings"
)

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

	menu := []types.MenuItem{
		{Name: "Cofee", Prices: map[string]float64{"small": 1.65, "medium": 1.80}},
		{Name: "Espresso", Prices: map[string]float64{"single": 1.90, "double": 2.25}},
	}
	for _, item := range menu {
		fmt.Println(item.Name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.Prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}
