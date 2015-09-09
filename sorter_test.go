package bitonic_test

import (
	"math/rand"
	"runtime"
	"testing"
	"time"

	"github.com/farazdagi/bitonic"
)

const SAMPLE_SIZE = 6

var lst []int = make([]int, SAMPLE_SIZE)

func randInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func init() {
	runtime.GOMAXPROCS(4)
	rand.Seed(time.Now().UTC().UnixNano())
}

func shuffle() {
	for i := 0; i < SAMPLE_SIZE; i++ {
		lst[i] = randInt(0, 100)
	}
}

func TestSorter(t *testing.T) {
	shuffle()
	bitonic.Sort(lst, bitonic.SORT_ASC)

	prev := lst[0]
	for _, i := range lst {
		if prev > i {
			t.Fatal("failed to confirm ASC sorting")
		}
		prev = i
	}
}

func BenchmarkSorter(b *testing.B) {
	shuffle()
	bitonic.Sort(lst, bitonic.SORT_ASC)
}
