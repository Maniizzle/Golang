package generics

import (
	"fmt"

	"golang.org/x/exp/maps"
)

func cloneSlice[V any](s []V) []V {
	result := make([]V, len(s))
	copy(result, s)
	// for i, v := range s {
	// 	result[i] = v
	// }
	return result
}

// Note: Keys of a map must be comparable
func cloneMap[K comparable, V any](s map[K]V) map[K]V {
	result := make(map[K]V, len(s))
	maps.Copy(result, s)
	// for i, v := range s {
	// 	result[i] = v
	// }
	return result
}

//Generic Add Function using Type Interface

type addable interface {
	int | float64 | string
}

func add[V addable](s []V) V {
	var result V
	for _, value := range s {
		result += value
	}
	return result
}

func GenericsMain() {
	testScores := []float64{
		87.3,
		105,
		63.5,
		27,
	}
	fmt.Println(testScores)

	clone := cloneSlice(testScores)
	fmt.Println(clone)

	fmt.Println(&testScores[0], &clone[0], clone)

	//map
	testNames := map[string]int32{
		"Bola": 24,
		"Bose": 32,
		"Debo": 18,
	}
	fmt.Println(testNames)

	cloneMap := cloneMap(testNames)
	fmt.Println(cloneMap)

	//fmt.Println(&testScores[0], &clone[0], clone)

	//Testing Generic Add function
	a1 := []float64{87.3, 105, 63.5}
	a2 := []int{1, 2, 3}
	a3 := []string{"bola", "tito", "taiwo"}

	add(a1)
	add(a2)
	add(a3)

}
