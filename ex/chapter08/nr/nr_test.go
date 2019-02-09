package nroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestLoad(t *testing.T) {
	start := time.Now()
	loadN(4, false)
	fmt.Printf("TestLoad finished in %s\n", time.Since(start))

}

func TestLoadParallel(t *testing.T) {
	start := time.Now()
	loadNParallel(4, false)
	fmt.Printf("TestLoadParallel finished in %s\n", time.Since(start))
}

func bench(iterations int, t *testing.B) {
	for i := 0; i < t.N; i++ {
		loadN(iterations, false)
	}
}

func benchp(iterations int, t *testing.B) {
	for i := 0; i < t.N; i++ {
		loadNParallel(iterations, false)
	}
}

func BenchmarkLoad8(t *testing.B)  { bench(8, t) }
func BenchmarkLoad16(t *testing.B) { bench(16, t) }
func BenchmarkLoad32(t *testing.B) { bench(32, t) }
func BenchmarkLoad64(t *testing.B) { bench(64, t) }

func BenchmarkLoadParallel8(t *testing.B)  { benchp(8, t) }
func BenchmarkLoadParallel16(t *testing.B) { benchp(16, t) }
func BenchmarkLoadParallel32(t *testing.B) { benchp(32, t) }
func BenchmarkLoadParallel64(t *testing.B) { benchp(64, t) }
