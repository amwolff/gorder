package main

import (
	"log"
	"time"

	"github.com/amwolff/gorder"
	"github.com/amwolff/gorder/dagenerator"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	start := time.Now()
	graph := dagenerator.Generate(10, 50, 30, 50, 30)
	log.Printf("DAG generation time: %v", time.Since(start))

	start = time.Now()
	s, err := gorder.TopologicalSort(graph, gorder.KAHN)
	log.Printf("Kahn resolve time: %v", time.Since(start))
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(s)

	start = time.Now()
	s, err = gorder.TopologicalSort(graph, gorder.DFS)
	log.Printf("DFS-based resolve time: %v", time.Since(start))
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(s)
}
