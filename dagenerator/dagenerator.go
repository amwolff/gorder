// Package dagenerator is based on C code from https://stackoverflow.com/questions/12790337/generating-a-random-dag
package dagenerator

import (
	"math"
	"math/rand"
	"time"
)

func Generate(minPerRank, maxPerRank, minRanks, maxRanks, percent int) (digraph map[int][]int) {
	var j, k, nodes int
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	digraph = make(map[int][]int)

	ranks := minRanks + int(math.Mod(random.Float64(), float64(maxRanks-minRanks+1)))
	for i := 0; i < ranks; i++ {
		newNodes := minPerRank + int(math.Mod(random.Float64(), float64(maxPerRank-minPerRank+1)))

		for j = 0; j < nodes; j++ {
			for k = 0; k < newNodes; k++ {
				if int(math.Mod(random.Float64(), 100)) < percent {
					digraph[j] = append(digraph[j], (k + nodes))
				}
			}
		}

		nodes += newNodes
	}
	return
}
