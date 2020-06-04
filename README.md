
# Golang Pipelines

![Actions Status](https://github.com/pabloos/Go-Pipelines/workflows/tests/badge.svg)  [![Coverage Status](https://coveralls.io/repos/github/pabloos/Go-Pipelines/badge.svg?branch=master)](https://coveralls.io/github/pabloos/Go-Pipelines?branch=master) ![dep Badge](https://img.shields.io/badge/dependencies-none-informational) ![MIT License badge](https://img.shields.io/badge/license-MIT-blue)

This repo contains an approach of Golang's Pipelines and its features, which are organized in branches, so you can go through the different milestones:

1. original-sketch: the original approach discussed on the [go blog](https://blog.golang.org/pipelines)
2. refactoring
3. autogen-stages
4. fanInfanOut
5. cancellation
6. order

## Some background

This repo exists because the first post in [my blog](https://pabloos.github.io/concurrency/pipelines/), where I explain the pattern.

## Usage

### A simple example

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

The code aims to be as declarative as possible:

```go
ctx, cancelCtx := context.WithCancel(context.Background())

defer cancelCtx() // always remenber to finish the pipeline running cancelling his context for avoiding memory leaks

numbers := []int{1, 2, 3}

input := Converter(numbers...)

firstStage := Pipeline(ctx, identity)(input)

secondStage := FanOut(ctx, firstStage, RoundRobin, Pipeline(ctx, double), Pipeline(ctx, square))

merged := FanIn(ctx, secondStage...)

thirdStage := Pipeline(ctx, divideBy(2))(merged)

result := Sink(thirdStage)

fmt.Println(result)
```

### Cancellation

```text

>--------------------------------> []
   1, 3, 3    mult(2)     error
```

```go
ctx, cancelCtx := context.WithCancel(context.Background())

defer cancelCtx()

numbers := []int{1, 2, 3}

input := Converter(numbers...)

firstStage := Pipeline(ctx, multBy(2), errFunc)(input)

result := Sink(firstStage)

fmt.Println(result) // []
```

When a stage finds an error the whole context gets cancelled and the whole pipeline closes, being cancelled also all the pipelines with the same context.

### Order Sink (InOrder, Reverse and NoOrder)

When you use fan out and fan in you could specify on the sink phase the order of the results in relation with the input order. There are three options:

|  Order  |  Behaviour  | Example |
|---------|--------------|--------|
| InOrder | Deterministc | 1, 2 < *2 > 2, 4
| Reverse | Deterministc | 1, 2 < *2 > 4, 2
| NoOrder | Indeterministc | 1, 2 < *2 > 2, 4 or 4, 2 (same as the previous example with Sink)

```text
             double
           ----------
         / 1 3 -> 2 6 \
>-------<               >------> [1, 2, 3]
  1 2 3  \   2 -> 4   /   /2
           ----------
             square

1st Stage   2nd Stage   3rd Stage
```

```go
...

result := SinkWithOrder(thirdStage, InOrder)

fmt.Println(result)
```

## Roadmap

- add more shedulers
- logging
- errors on observables
- benchmarks
