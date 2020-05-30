package pipelines

// Order returns a slice with the elements ordered
type Order func([]Element) Less

// Less is the type used in the sort goland std lib to order slices
type Less func(i, j int) bool

// InOrder returns a Less function for sort golang func "in order" to get the same order
func NoOrder(elements []Element) Less {
	return func(i, j int) bool {
		return elements[i].orderNum == elements[j].orderNum
	}
}

// InOrder returns a Less function for sort golang func "in order" to get an ascending orderNum order
func InOrder(elements []Element) Less {
	return func(i, j int) bool {
		return elements[i].orderNum < elements[j].orderNum
	}
}

// Reverse returns a Less function for sort golang func "in order" to get an descending orderNum order
func Reverse(elements []Element) Less {
	return func(i, j int) bool {
		return elements[i].orderNum > elements[j].orderNum
	}
}
