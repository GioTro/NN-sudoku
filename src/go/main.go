package main

import (
	"fmt"
	"sync"
	"time"
)

type sudoku struct {
	solved, unsolved [][]int8
}

// func listener(wg *sync.WaitGroup, ch chan sudoku) {
// 	defer close(ch)
// 	wg.Wait()
// }

func batch_process(n int, ch chan *sudoku) {
	//defer close(ch)
	var wg sync.WaitGroup
	var idxmap = make_indexmap()

	for i := 0; i < n; i++ {
		wg.Add(1)
		go process(50, &idxmap, ch, &wg)
		if i%3 == 0 {
			wg.Wait()
		}
	}
	//wg.Wait()
}

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

func fill(ch chan<- *[9][9]int8) {
	var count int8
	var a [9][9]int8
	for i := range a {
		for j := range (a)[i] {
			(a)[i][j] = count
			count++
		}
	}
	ch <- &a
}

func main() {

	const n = int(1e6)

	var out [n][9][9]int8
	var ch = make(chan *[9][9]int8)

	for i := 0; i < n; i++ {
		go fill(ch)
	}

	var start = time.Now()
	for i := 0; i < n; i++ {
		out[i] = *(<-ch)
	}
	var duration = time.Since(start).Seconds()
	fmt.Println(duration)

	for i, idx := range rand_idx(n) {
		var c = &out[idx]
		var d = &out[idx+1]
		fmt.Println(c == d)
		if i > 100 {
			break
		}
	}
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
