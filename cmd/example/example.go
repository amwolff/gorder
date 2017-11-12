package main

import (
	"fmt"
	"math"

	"github.com/amwolff/gorder"
	"github.com/sirupsen/logrus"
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

	o, err = gorder.TopologicalSort(digraph, "kahn")
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("Solution (Kahn): %v\n", o)

	o, err = gorder.TopologicalSort(digraph, "dfsbased")
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("Solution (DFS-based): %v\n", o)
}
