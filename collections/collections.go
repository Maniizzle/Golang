package collections

import (
	"fmt"

	"golang.org/x/exp/slices"
)

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
	fmt.Println(v)
	v[0] = 24
	fmt.Println(s)
	fmt.Println(v)

	//apppend to a slice
	s = append(s, 9, 10, 11)
	saa := &s
	fmt.Println(saa)

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
