package gorder

import (
	"testing"

	"github.com/amwolff/gorder/dagenerator"
)

func TestTopologicalSort(t *testing.T) {
	digraph := map[int][]int{
		1: {2, 4},
		2: {3, 5},
		3: {4, 5},
	}

	want := []int{1, 2, 3, 5, 4}


	_, err := TopologicalSort(digraph, 3)
	if err == nil {
		t.Fatal("TopologicalSort: should have returned an error")
	}

	o, err := TopologicalSort(digraph, KAHN)
	if err != nil {
		t.Fatalf("Kahn: %v", err)
	}
	for i, v := range o {
		if v != want[i] {
			t.Fatal("Kahn: output order does not match the desired order")
		}
	}

	o, err = TopologicalSort(digraph, DFS)
	if err != nil {
		t.Fatalf("DFS-based: %v", err)
	}
	for i, v := range o {
		if v != want[i] {
			t.Fatal("DFS-based: output order does not match the desired order")
		}
	}

	graphWithCycles := map[int][]int{
		1: {2, 4},
		2: {3, 5},
		3: {4, 5},
		4: {2},
	}
	_, err = TopologicalSort(graphWithCycles, KAHN)
	if err == nil {
		t.Fatal("Kahn: should have returned an error")
	}
	_, err = TopologicalSort(graphWithCycles, DFS)
	if err == nil {
		t.Fatal("DFS-based: should have returned an error")
	}
}

func BenchmarkTopologicalSort(b *testing.B) {
	digraph := dagenerator.Generate(10, 50, 30, 50, 30)
	b.Run("kahn", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := TopologicalSort(digraph, KAHN)
			if err != nil {
				b.Errorf("Kahn: %v", err)
			}
		}
	})
	b.Run("dfs based", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := TopologicalSort(digraph, DFS)
			if err != nil {
				b.Errorf("DFS-based: %v", err)
			}
		}
	})
}
