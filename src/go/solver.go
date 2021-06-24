package main

import (
	"math/rand"
	"time"
)

type board [81]byte

type set struct {
	solved, unsolved board
}

func valid_board(b board) bool {
	var tmp byte
	var response bool
	for idx := range b {
		tmp, b[idx] = b[idx], 0
		response = valid_move(tmp, idx, b)
		b[idx] = tmp
		if !response {
			return false
		}
	}
	return true
}

func valid_move(g byte, idx int, b board) bool {

	if g == 0 {
		return false
	}

	var row = idx / 9
	var col = idx % 9
	// check row
	for i := 0; i < 9; i++ {
		if i == col {
			continue
		}
		if b[row*9+i] == g {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if i == row {
			continue
		}
		if b[i*9+col] == g {
			return false
		}
	}

	var sqr = 3 * (row / 3)
	var sqc = 3 * (col / 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if row == (sqr+i) && col == (sqc+j) {
				continue
			}
			if b[(sqr+i)*9+(sqc+j)] == g {
				// pretty_print(b)
				return false
			}
		}
	}
	return true
}

func solver(b board, ch chan<- board) bool {
	var shuffle_2 = func(a [9]byte) [9]byte {
		rand.Seed(time.Now().UnixNano())
		for i := len(a) - 1; i > 0; i-- {
			var j = rand.Intn(i + 1)
			a[i], a[j] = a[j], a[i]
		}
		return a
	}

	var find_empty2 = func(a board) int {
		for i := range a {
			if a[i] == 0 {
				return i
			}
		}
		return -1
	}

	var idx = find_empty2(b)
	if idx == -1 {
		ch <- b
		return true
	}

	for _, g := range shuffle_2([9]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		if valid_move(g, idx, b) {
			b[idx] = g
			if solver(b, ch) {
				return true
			}
			b[idx] = 0
		}
	}
	return false
}

func plucker(keep int, b board) (board, int) {
	var nline = func() (out [81]int) {
		for i := range out {
			out[i] = i
		}
		return out
	}
	var shuffle_idx = func(a [81]int) [81]int {
		rand.Seed(time.Now().UnixNano())
		for i := len(a) - 1; i > 0; i-- {
			var j = rand.Intn(i + 1)
			a[i], a[j] = a[j], a[i]
		}
		return a
	}
	var left = len(b)
	var tmp byte
	var unique bool
	for _, idx := range shuffle_idx(nline()) {
		if left <= keep {
			return b, left
		}
		tmp, b[idx] = b[idx], 0
		unique = true
		for _, val := range [9]byte{1, 2, 3, 4, 5, 6, 7, 8, 9} {
			if val == tmp {
				continue
			} else if valid_move(val, idx, b) {
				unique = false
				break
			}
		}
		if !unique {
			b[idx] = tmp
		} else {
			left--
		}
	}
	return b, left
}
