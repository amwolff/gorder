package gorder

import (
	"errors"
	"regexp"
)

func TopologicalSort[T comparable, V []T](digraph map[T]V, algorithm string) (solution V, err error) {
	kahnRgxp, err := regexp.Compile(`[Kk]ahn\z`)
	if err != nil {
		return nil, err
	}
	dfsBasedRgxp, err := regexp.Compile(`[Dd][Ff][Ss]-?[Bb]ased\z`)
	if err != nil {
		return nil, err
	}

	if kahnRgxp.MatchString(algorithm) {
		if solution, err = kahn(digraph); err != nil {
			return nil, err
		}
	} else if dfsBasedRgxp.MatchString(algorithm) {
		if solution, err = dfsBased(digraph); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("no such algorithm")
	}
	return solution, nil
}

func kahn[T comparable, V []T](digraph map[T]V) (V, error) {
	indegrees := make(map[T]int)

	// loop through all diagraph and add increase indegrees of values
	for _, iter := range digraph {
		for _, v := range iter {
			indegrees[v]++
		}
	}

	var queue V
	for u := range digraph {
		if _, ok := indegrees[u]; !ok {
			queue = append(queue, u)
		}
	}

	var order V
	for len(queue) > 0 {
		u := queue[len(queue)-1]
		queue = queue[:(len(queue) - 1)]
		order = append(order, u)
		for _, v := range digraph[u] {
			indegrees[v]--
			if indegrees[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	for _, indegree := range indegrees {
		if indegree > 0 {
			return order, errors.New("not a DAG")
		}
	}
	return order, nil
}

func dfsBased[T comparable, V []T](digraph map[T]V) (V, error) {
	var (
		acyclic       = true
		order         V
		permanentMark = make(map[T]bool)
		temporaryMark = make(map[T]bool)
		visit         func(T)
	)

	visit = func(u T) {
		if temporaryMark[u] {
			acyclic = false
		} else if !(temporaryMark[u] || permanentMark[u]) {
			temporaryMark[u] = true
			for _, v := range digraph[u] {
				visit(v)
				if !acyclic {
					return
				}
			}
			delete(temporaryMark, u)
			permanentMark[u] = true
			order = append(V{u}, order...)
		}
	}

	for u := range digraph {
		if !permanentMark[u] {
			visit(u)
			if !acyclic {
				return order, errors.New("not a DAG")
			}
		}
	}
	return order, nil
}
