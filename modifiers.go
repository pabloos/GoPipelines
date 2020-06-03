package pipelines

import (
	"errors"
	"fmt"
)

func square(number int) (int, error) {
	return number * number, nil
}

func cube(number int) (int, error) {
	return number * number * number, nil
}

var double = multBy(2)

func identity(number int) (int, error) {
	return number, nil
}

func addTo(suma int) functor {
	return func(number int) (int, error) {
		return number + suma, nil
	}
}

func subTo(sub int) functor {
	return func(number int) (int, error) {
		return number - sub, nil
	}
}

func multBy(multiplier int) functor {
	return func(number int) (int, error) {
		return number * multiplier, nil
	}
}

func divideBy(divisor int) functor {
	return func(number int) (int, error) {
		return number / divisor, nil
	}
}

func cancel(n int) (int, error) {
	fmt.Println("cancel transformer called")
	return n, errors.New("hehe")
}
