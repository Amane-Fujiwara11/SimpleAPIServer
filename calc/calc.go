package calc

import (
	"errors"
)

func Fact(number int) (int, error) {
	if number < 1 {
		return 0, errors.New("正の整数は1以上でなければなりません。")
	}
	result := 1
	for i := number; i > 1; i-- {
		result *= i
	}
	return result, nil
}
