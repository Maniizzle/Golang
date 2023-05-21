package controlflow

import "fmt"

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

func Action() {
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
