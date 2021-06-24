package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// func batch_process(n, keep int, ch chan set) {
// 	for i := 0; i < n; i++ {
// 		go process(keep, ch)
// 	}
// }

func worker(n, keep int, ch chan set) {
	var ch_out = make(chan board)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		var b board
		go solver(b, ch_out)
	}
	wg.Add(n)
	for i := 5; i > 0; i-- {
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

// func process(keep int, ch_out chan<- set) {
// 	var b, c board
// 	var left int

// 	// var solved board
// 	var out = make(chan board)
// 	go solver(b, out)
// 	b = <-out
// 	c, left = plucker(keep, b)
// 	if left > keep {
// 		process(keep, ch_out)
// 	} else {
// 		ch_out <- set{
// 			solved:   b,
// 			unsolved: c,
// 		}
// 	}
// }

// func process_single(keep int) set {
// 	var b, c board
// 	var left int
// 	var out = make(chan board)
// 	go solver(b, out)
// 	b = <-out
// 	c, left = plucker(keep, b)

// 	if left > keep {
// 		return process_single(keep)
// 	} else {
// 		return set{
// 			solved:   b,
// 			unsolved: c,
// 		}
// 	}
// }

func main() {

	// currently it can generate a valid board in
	// .1 seconds I would like to make it a bit faster
	// then make the goroutines more efficient.
	// but I will likely use this setup because this
	// is taking too much time.

	const n = int(10000)
	var out [n]set
	var ch = make(chan set)

	go worker(n, 40, ch)

	var start = time.Now()
	for i := 0; i < n; i++ {
		out[i] = <-ch
	}
	var duration = time.Since(start).Seconds()

	for _, o := range out {
		fmt.Println(valid_board(o.solved))
	}
	var rnd = rand.Intn(len(out))
	pretty_print(out[rnd].solved)
	pretty_print(out[rnd].unsolved)
	fmt.Println(duration)
}
