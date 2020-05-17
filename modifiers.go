package pipelines

func square(number int) int {
	return number * number
}

func cube(number int) int {
	return square(number) * number
}

var double = multBy(2)

func identity(number int) int {
	return number
}

func addTo(suma int) functor {
	return func(number int) int {
		return number + suma
	}
}

func subTo(sub int) functor {
	return func(number int) int {
		return number - sub
	}
}

func multBy(multiplier int) functor {
	return func(number int) int {
		return number * multiplier
	}
}

func divideBy(divisor int) functor {
	return func(number int) int {
		return number / divisor
	}
}
