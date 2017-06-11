package montecarlo

import (
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func monteCarloNaive(n int) float64 {
	in := 0.0
	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()
		d := math.Sqrt(x*x + y*y)
		if d < 1 {
			in++
		}
	}
	return 4.0 * in / float64(n)
}

func monteCarloParallelNaive(n int) float64 {
	mtx := sync.Mutex{}
	in := 0.0
	for i := 0; i < n; i++ {
		go func() {
			x := rand.Float64()
			y := rand.Float64()
			d := math.Sqrt(x*x + y*y)
			if d < 1 {
				mtx.Lock()
				in++
				mtx.Unlock()
			}
		}()
	}
	mtx.Lock()
	defer mtx.Unlock()
	return 4.0 * in / float64(n)
}

func monteCarloBatch1(n int, batchSize int) float64 {
	wg := sync.WaitGroup{}
	in := uint64(0)
	batches := n / batchSize
	wg.Add(batchSize)
	for b := 0; b < batchSize; b++ {
		go func() {
			for i := 0; i < batches; i++ {
				x := rand.Float64()
				y := rand.Float64()
				d := math.Sqrt(x*x + y*y)
				if d < 1 {
					atomic.AddUint64(&in, 1)
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return 4.0 * float64(in) / float64(n)
}

func monteCarloBatch2(n int, batchSize int) float64 {
	wg := sync.WaitGroup{}
	in := uint64(0)
	batches := n / batchSize
	wg.Add(batchSize)
	for b := 0; b < batchSize; b++ {
		go func() {
			localIn := uint64(0)
			for i := 0; i < batches; i++ {
				x := rand.Float64()
				y := rand.Float64()
				d := math.Sqrt(x*x + y*y)
				if d < 1 {
					localIn++
				}
			}
			atomic.AddUint64(&in, localIn)
			wg.Done()
		}()
	}
	wg.Wait()
	return 4.0 * float64(in) / float64(n)
}

func monteCarloBatch3(n int, batchSize int) float64 {
	wg := sync.WaitGroup{}
	in := uint64(0)
	t := time.Now().Unix()
	batches := n / batchSize
	wg.Add(batchSize)
	for b := 0; b < batchSize; b++ {
		go func(s int64) {
			r := rand.New(rand.NewSource(t + s))
			localIn := uint64(0)
			for i := 0; i < batches; i++ {
				x := r.Float64()
				y := r.Float64()
				d := math.Sqrt(x*x + y*y)
				if d < 1.0 {
					localIn++
				}
			}
			atomic.AddUint64(&in, localIn)
			wg.Done()
		}(int64(b))
	}
	wg.Wait()
	return 4.0 * float64(in) / float64(n)
}
