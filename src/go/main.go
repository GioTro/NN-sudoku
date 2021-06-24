package main

import (
	"fmt"
	"time"
)

type sudoku struct {
	solved, unsolved [][]int8
}

// func listener(wg *sync.WaitGroup, ch chan sudoku) {
// 	defer close(ch)
// 	wg.Wait()
// }

// func batch_process(n int, ch chan sudoku) {
// 	//defer close(ch)
// 	var wg sync.WaitGroup
// 	var idxmap = make_indexmap()

// 	for i := 0; i < n; i++ {
// 		wg.Add(1)
// 		go process(50, &idxmap, ch, &wg)
// 		if i%3 == 0 {
// 			wg.Wait()
// 		}
// 	}
// 	//wg.Wait()
// }

// func test_process(n int) bool {
// 	var idxmap = make_indexmap()
// 	//var bb = make2dint8(9)
// 	var hold = make([]sudoku, n)
// 	for i := 0; i < n; i++ {
// 		var start = time.Now()
// 		hold[i] = process(50, &idxmap)
// 		var after = time.Since(start).Seconds()
// 		fmt.Println(after)
// 	}

// 	// check validity
// 	for _, i := range hold {
// 		if !valid_board(&i.solved, &idxmap) {
// 			pretty_print(&i.solved)
// 			pretty_print(&i.unsolved)
// 			return false
// 		}
// 	}
// 	return true
// }

// func push(n int, ch chan<- bool) {
// 	for i := 0; i < n; i++ {
// 		ch <- true
// 	}
// }
// func plistener(n_workers int, n int, ch_in chan board, ch_out chan sudoku) {
// 	var idx = make_indexmap()
// 	var signal = make(chan bool)
// 	go push(n, signal)

// 	for i := 0; i < n_workers; i++ {
// 		go solver(50, ch_in, ch_out)
// 		go worker(n_workers, &idx, ch_in, signal)
// 	}
// }

// func single(keep int, idx *indexmap, ch_in chan<- [][]int8) {
// 	var b = make_board(idx)
// 	make_valid_board(b)
// 	ch_in <- (*b).board
// }

func batch_process(n, keep int, ch chan set) {
	for i := 0; i < n; i++ {
		go process(keep, ch)
	}
}

func main() {

	// currently it can generate a valid board in
	// .1 seconds I would like to make it a bit faster
	// then make the goroutines more efficient.
	// but I will likely use this setup because this
	// is taking too much time.

	const n = int(1000)
	var out [n]set
	var ch = make(chan set)

	go batch_process(n, 40, ch)
	var start = time.Now()
	for i := 0; i < n; i++ {
		out[i] = <-ch
	}
	var duration = time.Since(start).Seconds()

	for _, o := range out {
		fmt.Println(valid_board(o.solved))
	}
	fmt.Println(duration)

	// var out [n][][]int8
	// //var ch_out = make(chan sudoku)
	// var ch_out = make(chan [][]int8)

	// var start = time.Now()
	// var idx = make_indexmap()
	// //go plistener(1, n, ch_in, ch_out)

	// for i := 0; i < n; i++ {
	// 	go single(50, &idx, ch_out)
	// }

	// var count int
	// for i := 0; i < n; i++ {
	// 	out[count] = <-ch_out
	// 	count++
	// }
	// var duration = time.Since(start).Seconds()

	// for _, b := range out {
	// 	var idx = make_indexmap()
	// 	fmt.Println(valid_board(&b, &idx))
	// }

	// fmt.Println(duration)
	// var n_train = int(1e3) // how training examples to generate
	// var n_val = int(0)     // how val examples to generate
	// var n_test = int(0)    // how test examples to generate

	// fmt.Println([3]int{n_train, n_val, n_test})
	// var n = 1000
	// //var wg sync.WaitGroup

	// var ch = make(chan *sudoku)

	// var out = make([]*sudoku, n)

	// var start = time.Now()

	// go batch_process(n, ch)

	// for i := 0; i < n; i++ {
	// 	out[i] = <-ch
	// }

	// var duration = time.Since(start)

	// fmt.Println(duration)

	// var idxmap = make_indexmap()
	// var out2 = make([]sudoku, n)
	// //count = 0
	// start = time.Now()
	// for i := 0; i < n; i++ {
	// 	out2[i] = process_single(50, &idxmap)
	// }
	// duration = time.Since(start)
	// fmt.Println(duration)

	// // for _, n := range [3]int{n_train, n_val, n_test} {
	// // 	//var ch = make(chan [2][9][9]int8)
	// // 	//var wg sync.WaitGroup
	// // 	//wg.Add(1)

	// // }
	// // if test_process(300) {
	// // 	fmt.Println("success!")
	// // } else {
	// // 	fmt.Println("failure!")
	// // }
}
