package main

import (
	"time"

	"github.com/amwolff/gorder"
	"github.com/amwolff/gorder/dagenerator"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.StandardLogger()
	logger.SetLevel(logrus.DebugLevel)

	start := time.Now()
	graph := dagenerator.Generate(10, 50, 30, 50, 30)
	logger.Infof("DAG generation time: %v", time.Since(start))

	start = time.Now()
	s, err := gorder.TopologicalSort(graph, "kahn")
	logger.Infof("Kahn resolve time: %v", time.Since(start))
	if err != nil {
		logger.Fatal(err)
	}
	spew.Dump(s)

	start = time.Now()
	s, err = gorder.TopologicalSort(graph, "dfsbased")
	logger.Infof("DFS-based resolve time: %v", time.Since(start))
	if err != nil {
		logger.Fatal(err)
	}
	spew.Dump(s)
}
