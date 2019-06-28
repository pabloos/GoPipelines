package main

import (
	"fmt"
)

func main() {
	for _, number := range NewPipeline(start, final, add2, square, add2).Exec(1, 2, 3, 4, 5, 6, 7, 8, 9) {
		fmt.Println(number)
	}
}
