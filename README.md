
# Golang Pipelines

![Actions Status](https://github.com/pabloos/Go-Pipelines/workflows/tests/badge.svg)

This repo contains an approach of Golang's Pipelines and its features. It's organized on branches, so you can go through the different milestones:

1. original-sketch: the original approach discussed on the [go blog](https://blog.golang.org/pipelines)
2. refactoring
3. autogen-stages
4. fanInfanOut

## Some background

This repo exists because the first post in [my blog](https://pabloos.github.io/concurrency/pipelines/), where I explain the pattern.

## Usage

A simple example:

```text
             double
           ----------
         / 1 3 -> 2 6 \
>-------<               >------> [1, 2, 3] (maybe in other order)
  1 2 3  \   2 -> 4   /   /2
           ----------
             square

1st Stage   2nd Stage   3rd Stage
```

```go
numbers := []int{1, 2, 3}

input := Converter(numbers...)

firstStage := NewPipeline(identity)(input)

secondStage := FanOut(firstStage, RoundRobin, NewPipeline(double), NewPipeline(square))

merged := FanIn(secondStage...)

thirdStage := NewPipeline(divideBy(2))(merged)

result := Sink(thirdStage)

fmt.Println(result)
```

## Roadmap

- add more shedulers
- converters
- cancellation
- errors on observables
- order in merge
- tests
- benchmarks
- buffered channels?
