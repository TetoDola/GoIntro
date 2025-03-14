package tutorials

import (
	"errors"
	"strconv"
)

func Divide(num1 int, num2 int) (int, error) {
	var err error
	if num2 == 0 {
		err = errors.New("cannot divide by zero")
		return 0, err
	}
	return num1 / num2, err
}

func Add(num1 int, num2 int) string {

	return "The result is: " + strconv.Itoa(num1+num2)
}
