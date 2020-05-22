package pipelines

import "errors"

func square(number int) (error, int) {
	return nil, number * number
}

func cube(number int) (error, int) {
	return nil, number * number * number
}

var double = multBy(2)

func identity(number int) (error, int) {
	return nil, number
}

func addTo(suma int) functor {
	return func(number int) (error, int) {
		return nil, number + suma
	}
}

func subTo(sub int) functor {
	return func(number int) (error, int) {
		return nil, number - sub
	}
}

func multBy(multiplier int) functor {
	return func(number int) (error, int) {
		return nil, number * multiplier
	}
}

func divideBy(divisor int) functor {
	return func(number int) (error, int) {
		return nil, number / divisor
	}
}

func cancel(n int) (error, int) {
	return errors.New("hehe"), n
}
