package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(n, keep int, ch chan set) {
	var ch_out = make(chan board)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		var b board
		go solver(b, ch_out)
	}
	wg.Add(n)
	for i := 20; i > 0; i-- {
		go listen(n, keep, ch_out, ch, &wg)
	}
	wg.Wait()
	close(ch_out)
}

func listen(n, keep int, ch chan board, ch_out chan set, wg *sync.WaitGroup) {
	var c board
	var left int
	for o := range ch {
		for {
			c, left = plucker(keep, o)
			if left <= keep {
				break
			}
		}
		ch_out <- set{
			solved:   o,
			unsolved: c,
		}
		wg.Done()
	}
}

func main() {
	runtime.GOMAXPROCS(8)

	// currently it can generate a valid board in ~.01 seconds
	// There is a lot of overhead in running this concurrently
	// Pick suitable settings for you machine and fly like a falcon.

	const n = int(1e5)
	const n_per_batch = int(1e4)
	var out = make([]set, n)
	//var ch = make(chan set)
	var idx int

	var start = time.Now()
	for i := 0; i < n/n_per_batch; i++ {
		var ch = make(chan set)
		go worker(n_per_batch, 40, ch)

		for i := 0; i < n_per_batch; i++ {
			out[idx] = <-ch
			idx++
		}
		fmt.Println("Generated ", idx, " elapsed time ", time.Since(start).Minutes(), " minutes!")
	}
	var duration = time.Since(start).Seconds()

	fmt.Println(duration)
	save_data(out, "sudoku.txt")
}
