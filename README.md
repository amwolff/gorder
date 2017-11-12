# gorder

A golang implementation of topological sorting algorithms for fast ordering.

## Installation:

`$ go get -u github.com/amwolff/gorder`

## Usage:

Check `cmd/example/example.go` for example usage.

```
$ go run cmd/example/example.go
Solution (Kahn): [1 {3.141592653589793} 3 5 4]
Solution (DFS-based): [1 {3.141592653589793} 3 5 4]
```

## Notes:

* Maps are one of the common ways to store graphs in Go. The `TopologicalSort` function input supports `map[interface{}][]interface{}` maps.

* Sub-package `dagenerator` was developed and used for generating random DAGs for performance tests (`cmd/measurements` should be rewritten as benchmarks).

* Implementation of Kahn's algorithm does pre-computation in the beginning of its work in order to calculate indegree of every vertex in the graph. This may be split into two separate algorithms/functions in the future:
    - one that doesn't take any additional input but the graph (current implementation)
    - and one that takes additional input of indegrees