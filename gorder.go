package gorder

import (
	"errors"
	"regexp"
)

func TopologicalSort(digraph map[int][]int, algorithm string) (solution []int, err error) {
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

func kahn(digraph map[int][]int) ([]int, error) {
	inDegree := make(map[int]int)
	for u := range digraph {
		if digraph[u] != nil {
			for _, v := range digraph[u] {
				inDegree[v]++
			}
		}
	}

	var queue []int
	for u := range digraph {
		if _, ok := inDegree[u]; !ok {
			queue = append(queue, u)
		}
	}

	var order []int
	for len(queue) > 0 {
		u := queue[len(queue)-1]
		queue = queue[:(len(queue) - 1)]
		order = append(order, u)
		for _, v := range digraph[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	for _, in := range inDegree {
		if in > 0 {
			return order, errors.New("Kahn: not a DAG")
		}
	}
	return order, nil
}

func dfsBased(digraph map[int][]int) ([]int, error) {
	var (
		acyclic       = true
		order         []int
		permanentMark = make(map[int]bool)
		temporaryMark = make(map[int]bool)
		visit         func(int)
	)

	visit = func(u int) {
		if temporaryMark[u] {
			acyclic = false
		} else if !permanentMark[u] {
			temporaryMark[u] = true
			for _, v := range digraph[u] {
				visit(v)
				if !acyclic {
					return
				}
			}
			delete(temporaryMark, u)
			permanentMark[u] = true
			order = append([]int{u}, order...)
		}
	}

	for u := range digraph {
		if !permanentMark[u] {
			visit(u)
			if !acyclic {
				return order, errors.New("DFS-based: not a DAG")
			}
		}
	}
	return order, nil
}
