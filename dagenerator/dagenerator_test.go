package dagenerator

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGenerate(t *testing.T) {
	const (
		minPerRank = 1
		maxPerRank = 5
		minRanks   = 3
		maxRanks   = 5
		percent    = 30
	)

	g := Generate(minPerRank, maxPerRank, minRanks, maxRanks, percent)

	// check manually
	spew.Dump(g)
}
