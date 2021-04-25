package gorder
import "errors"

type Algo int

const (
	DFS Algo = iota
	Kahn
)

func TopologicalSort(digraph map[interface{}][]interface{}, algorithm Algo) (solution []interface{}, err error) {
	switch algorithm {
	case DFS:
		if solution, err = dfsBased(digraph); err != nil {
			return nil, err
		}
	case Kahn:
		if solution, err = kahn(digraph); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("no such algorithm")
	}
	return solution, nil
}

func kahn(digraph map[interface{}][]interface{}) ([]interface{}, error) {
	indegrees := make(map[interface{}]int)
	for u := range digraph {
		if digraph[u] != nil {
			for _, v := range digraph[u] {
				indegrees[v]++
			}
		}
	}

	var queue []interface{}
	for u := range digraph {
		if _, ok := indegrees[u]; !ok {
			queue = append(queue, u)
		}
	}

	var order []interface{}
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

func dfsBased(digraph map[interface{}][]interface{}) ([]interface{}, error) {
	var (
		acyclic       = true
		order         []interface{}
		permanentMark = make(map[interface{}]bool)
		temporaryMark = make(map[interface{}]bool)
		visit         func(interface{})
	)

	visit = func(u interface{}) {
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
			order = append([]interface{}{u}, order...)
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
