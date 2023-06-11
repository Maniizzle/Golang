package gofunctions

import (
	"math"
	"net/http"
	"reflect"
	"strings"
)

type MathExpr = string

const (
	AddExpr      = MathExpr("add")
	SubtractExpr = MathExpr("sub")
	MultiplyExpr = MathExpr("multiply")
)

// variadic parameters
// they also need to be the last parameters
func sum(values ...float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}
	return total
}
func add(a, b float64) float64 {

	return a + b
}
func subtract(a, b float64) float64 {

	return a - b
}
func multiply(a, b float64) float64 {

	return a * b
}

func anonymousFunction() {

	func() {
		println("anonymous functions")
	}()
}
func variableFunction() {

	a := func() {
		println("anonymous functions")
	}
	a()
	a()
}

func passFunction(add func() string) {
	add()
}
func passFunctionTwo(f1, f2 float64, mathexpr func(float64, float64) float64) float64 {
	return 2 * mathexpr(f1, f2)
}

func returnFunction() func(float64, float64) float64 {
	return add
}
func returnFunctionTwo() func(float64, float64) float64 {
	return func(f float64, f2 float64) float64 {
		return f + f2
	}
}

func mathExpression(expr MathExpr) func(float64, float64) float64 {
	switch expr {
	case AddExpr:
		return add
	case MultiplyExpr:
		return multiply
	case SubtractExpr:
		return subtract
	default:
		return returnFunction()
	}
}

func statefulFunction() func() int64 {
	x := 1.0
	println("x is", x)
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))
	}
}

func badStateAnonymousFunction() {

	// i gets updated before the functions are run, thus i is 10 when they run
	var funcs []func() int64
	for i := 0; i < 10; i++ {
		funcs = append(funcs, func() int64 {
			return int64(math.Pow(float64(i), 2))
		})
	}
	for _, f := range funcs {
		println(f())
	}
}
func betterStateAnonymousFunction() {

	var funcs []func() int64
	for i := 0; i < 10; i++ {
		cleanI := i
		funcs = append(funcs, func() int64 {
			return int64(math.Pow(float64(cleanI), 2))
		})
	}
	for _, f := range funcs {
		println(f())
	}
}

func init() {
	println("testing init function")

}

func FunctionMain() {

	var no int = 100
	println(reflect.TypeOf(no))

	numbers := []float64{1, 2.3, 4.5, 66.7}
	result := sum(numbers...)
	println(result)

	sv := NewSemanticVersion(1, 2, 3)
	sv.IncrementMajor()
	sv.IncrementMinor()
	sv.IncrementPatch()
	println(sv.String())

	var tripper http.RoundTripper = &RoundTripCounter{}
	r, _ := http.NewRequest(http.MethodGet, "http://pluralsight.com", strings.NewReader("test call"))
	_, _ = tripper.RoundTrip(r)

	//passing and return functions
	s := returnFunction()
	println(s(2, 3))

	println(returnFunctionTwo()(2.3, 4.6))

	passFunction(func() string {
		println("passing functions")
		return "exit"
	})
	println(passFunctionTwo(1.3, 2.4, add))
	println("add it up", mathExpression(AddExpr)(2, 3))
	println("subtract it up", mathExpression(SubtractExpr)(2, 3))
	println("multiply it up", mathExpression(MultiplyExpr)(2, 3))

	p2 := statefulFunction()
	p3 := statefulFunction()
	values := p3()
	println("p3 is", values)
	value := p2()
	println(value)

	value = p2()
	println(value)

	value = p2()
	println(value)

}
