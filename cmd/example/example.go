package main

import (
	"fmt"
	"log"
	"math"

	"github.com/amwolff/gorder"
)

func main() {
	var (
		a   struct{ Pi float64 }
		o   []interface{}
		err error
	)

	a.Pi = math.Pi

	digraph := map[interface{}][]interface{}{
		1: []interface{}{a, "4"},
		a: []interface{}{3, "5"},
		3: []interface{}{"4", "5"},
	}

	o, err = gorder.TopologicalSort(digraph, gorder.KAHN)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Solution (Kahn): %v\n", o)

	o, err = gorder.TopologicalSort(digraph, gorder.DFS)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Solution (DFS-based): %v\n", o)
}
