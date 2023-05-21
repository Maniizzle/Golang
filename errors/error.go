package errors

import (
	"errors"
	"fmt"
)

func ErrorHandling() {

	err := errors.New("error occurred")
	fmt.Println(err)

	err2 := fmt.Errorf("merging error with first error: %w", err)
	fmt.Println(err2)

	res, err := divide(0, 3)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(res)
	}

}

func divide(dividend, divisor int) (int, error) {
	if dividend == 0 {
		return 0, errors.New("number is not divisible by zero")
	}
	return divisor / dividend, nil
}
