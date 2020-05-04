package main

import "testing"

func BenchmarkPipeline(b *testing.B) {

	pip := NewPipeline(addTo(2), square, addTo(2))

	for n := 0; n < b.N; n++ {
		input := Converter(1+n, 2+n, 3+n, 4+n, 5+n)

		Sink(pip(input))
	}
}

func BenchmarkPipeline2(b *testing.B) {
	pip := NewPipeline(addTo(2), square, addTo(2))

	for n := 0; n < b.N; n++ {
		input := Converter(1+n, 2+n, 3+n, 4+n, 5+n, 6+n, 7+n, 8+n, 9+n)

		Sink(pip(input))

	}
}
