package operations

import (
	"github.com/MRickers/math/display"
)

func Add(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func HelloOperations() {
	display.PrintHello("operations")
}
