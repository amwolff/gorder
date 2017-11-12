package main

import (
	"fmt"

	"github.com/amwolff/gorder"
	"github.com/sirupsen/logrus"
)

func serialize(m map[int][]int) map[interface{}][]interface{} {
	s := make(map[interface{}][]interface{})
	for k, v := range m {
		for _, val := range v {
			s[k] = append(s[k], val)
		}
	}
	return s
}

func deserialize(o []interface{}) []int {
	d := make([]int, len(o))
	for i := range o {
		d[i] = o[i].(int)
	}
	return d
}

func main() {
	var (
		output []interface{}
		err    error
	)

	digraph := map[int][]int{
		1: []int{2, 4},
		2: []int{3, 5},
		3: []int{4, 5},
	}

	output, err = gorder.TopologicalSort(serialize(digraph), "kahn")
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("Solution (Kahn): %d\n", deserialize(output))

	output, err = gorder.TopologicalSort(serialize(digraph), "dfsbased")
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("Solution (DFS-based): %d\n", deserialize(output))
}
