// package add contains methods related to adding
package math

type Adder interface {
	Add(x, y int) int
}

type adder struct{}

func NewAdder() *adder {
	return &adder{}
}

func (a adder) Add(x, y int) int {
	return Add(x, y)
}

// Add sums two integers and return the result
func Add(x, y int) int {
	var sum int

	if x > y {
		sum = x

		for y > 0 {
			y--
			sum += 1
		}
	} else {
		sum = y

		for x > 0 {
			sum += 1
			x--
		}

	}

	return sum
}
