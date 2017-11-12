package gorder

import "testing"

func TestTopologicalSort(t *testing.T) {
	digraph := map[interface{}][]interface{}{
		1: []interface{}{2, 4},
		2: []interface{}{3, 5},
		3: []interface{}{4, 5},
	}

	want := []int{1, 2, 3, 5, 4}

	_, err := TopologicalSort(digraph, "ultrasuperfast")
	if err == nil {
		t.Fatal("TopologicalSort: should have returned an error")
	}

	o, err := TopologicalSort(digraph, "kahn")
	if err != nil {
		t.Fatalf("Kahn: %v", err)
	}
	for i, v := range o {
		if v != want[i] {
			t.Fatal("Kahn: output order does not match the desired order")
		}
	}

	o, err = TopologicalSort(digraph, "dfsbased")
	if err != nil {
		t.Fatalf("DFS-based: %v", err)
	}
	for i, v := range o {
		if v != want[i] {
			t.Fatal("DFS-based: output order does not match the desired order")
		}
	}

	graphWithCycles := map[interface{}][]interface{}{
		1: []interface{}{2, 4},
		2: []interface{}{3, 5},
		3: []interface{}{4, 5},
		4: []interface{}{2},
	}
	_, err = TopologicalSort(graphWithCycles, "kahn")
	if err == nil {
		t.Fatal("Kahn: should have returned an error")
	}
	_, err = TopologicalSort(graphWithCycles, "dfsbased")
	if err == nil {
		t.Fatal("DFS-based: should have returned an error")
	}
}
