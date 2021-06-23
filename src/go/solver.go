package main

import (
	"fmt"
	"sync"
)

type board struct {
	board          [][]int8
	solution       [][]int8
	solution_count int
	idx            *indexmap
	//get            func(int, int, string, *board) []int8
}

type indexmap struct {
	rmap, cmap [][]index
	sqrmap     [][][]index
	all        []index
	square     func(int, int, [][][]index) []index
	row        func(int, [][]index) []index
	col        func(int, [][]index) []index
}

type index struct {
	row, col int
}

func get(row, col int, s string, board *[][]int8, idx *indexmap) []int8 {
	var out = make([]int8, 9)
	var idxarr []index

	if s == "row" {
		idxarr = (idx).row(row, (idx).rmap)
	} else if s == "col" {
		idxarr = (idx).col(col, (idx).cmap)
	} else if s == "sqr" {
		idxarr = (idx).square(row/3, col/3, (idx).sqrmap)
	} else {
		fmt.Println("danger") // for debugging
	}

	for i, idx := range idxarr {
		out[i] = (*board)[idx.row][idx.col]
	}

	return out
}

func make_board(idx *indexmap) *board {
	var b = make2dint8(9)
	var c = make2dint8(9)
	var out = board{
		board:          b,
		solution:       c,
		solution_count: 0,
		idx:            idx,
	}
	return &out
}

func make_indexmap() indexmap {
	var rows = make([][]index, 9) //[][]index
	var cols = make([][]index, 9) //[][]index
	var all = make([]index, 81)
	var squares = make([][][]index, 3)
	squares[0] = make([][]index, 3)
	squares[1] = make([][]index, 3)
	squares[2] = make([][]index, 3)
	// make row and col map
	var count int
	for i := 0; i < 9; i++ {
		var r = make([]index, 9)
		var c = make([]index, 9)
		for j := 0; j < 9; j++ {
			r[j] = index{row: i, col: j}
			c[j] = index{row: j, col: i}
			all[count] = index{row: i, col: j}
			count++
		}
		rows[i] = r
		cols[i] = c
	}

	for i := 0; i < 9; i++ {
		var s = make([]index, 9)
		var count int
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				orow := 3*(i%3) + j
				ocol := 3*(i/3) + k
				s[count] = index{row: orow, col: ocol}
				count++
			}
		}
		squares[(i % 3)][(i / 3)] = s
	}

	return indexmap{
		rmap:   rows,
		cmap:   cols,
		sqrmap: squares,
		all:    all,

		square: func(row int, col int, sm [][][]index) []index {
			return sm[row][col]
		},

		row: func(idx int, rm [][]index) []index {
			return rm[idx]
		},

		col: func(idx int, cm [][]index) []index {
			return cm[idx]
		},
	}
}

func valid_board(b *[][]int8, idx *indexmap) bool {
	var nline = numline_int8(1, 10)

	// Helper function
	var member = func(a []int8, b []int8) bool {
		var count int
		for _, a_val := range a {
			for _, b_val := range b {
				if (a_val == int8(0)) || (b_val == int8(0)) {
					return false
				}
				if a_val == b_val {
					count++
					break
				}
			}
		}
		if count == 9 {
			return true
		} else {
			return false
		}
	}
	// rows
	for row := 0; row < 9; row++ {
		if !member(get(row, 0, "row", b, idx), nline) {
			return false
		}
	}
	for col := 0; col < 9; col++ {
		if !member(get(0, col, "col", b, idx), nline) {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !member(get(i, j, "sqr", b, idx), nline) {
				return false
			}
		}
	}
	// passed all checks
	return true
}

func valid_move(num int8, idx index, idxmp *indexmap, b *[][]int8) bool {
	if num == 0 {
		return false
	}

	// row
	for _, val := range get(idx.row, idx.col, "row", b, idxmp) {
		if num == val {
			return false
		}
	}

	// col
	for _, val := range get(idx.row, idx.col, "col", b, idxmp) {
		if num == val {
			return false
		}
	}
	// sqr
	for _, val := range get(idx.row, idx.col, "sqr", b, idxmp) {
		if num == val {
			return false
		}
	}
	// passed all checks
	return true
}

// func recursive(limit int, b *board, ref [][]int8) {
// 	/*
// 		Recursively solves the board via backtracking.
// 		Solves the board in place

// 		limit number of solution to find. This is going to be
// 		either 1 for find first solution or 2 to check if the
// 		solution is unique.
// 	*/
// 	if !(b.solution_count < limit) {
// 		return
// 	}

// 	var count int
// 	for row := 0; row < 9; row++ {
// 		for col := 0; col < 9; col++ {
// 			var idx = index{row: row, col: col}
// 			count++
// 			if ref[idx.row][idx.col] == 0 {
// 				var guess = shuffled_int8(numline_int8(1, 10))
// 				for _, num := range guess {
// 					if valid_move(b, num, idx) {
// 						ref[idx.row][idx.col] = num
// 						if !(b.solution_count < limit) {
// 							return
// 						}
// 						recursive(limit, b, ref)

// 						// Check this statement if this would be correct
// 						ref[idx.row][idx.col] = int8(0) // back track
// 					}
// 				}
// 				return
// 			}
// 		}
// 	}
// 	b.board = deep_copy(ref)
// 	if valid_board(b) {
// 		pretty_print(&b.board)
// 		b.solution_count++
// 	}
// }

// findfirst solution, returns true if
// there was a solution to be found
// func find_first(b *board) bool {
// 	if valid_board(b) {
// 		return true
// 	}

// 	var ref = deep_copy(b.board)
// 	recursive(1, b, ref)

// 	if b.solution_count == 0 {
// 		return false
// 	} else {
// 		return true
// 	}
// }

// // Alias
// func find_more(b *board) bool {
// 	var ref = deep_copy(b.board)
// 	recursive(2, b, ref)
// 	if (*b).solution_count > 1 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func make_valid_board(b *board) {
	// this is fast when generating the first board.
	var nline = numline_int8(1, 10)
	for {
		for _, idx := range (*b.idx).all {
			var guess = shuffled_int8(nline)
			for _, g := range guess {
				if valid_move(g, idx, b.idx, &b.board) {
					b.board[idx.row][idx.col] = g
				}
			}
		}
		if valid_board(&b.board, b.idx) {
			break
		}
		b.board = make2dint8(9)
	}
}

/*
	Try removing random points
*/

func pluck(keep int, b *board) bool {
	var idx = shuffled_index(b.idx.all)
	var pointer int
	var reference int

	var left = len(idx)
	for (reference != -1) && left > keep {
		var guess = idx[pointer]
		pointer++

		if (*b).board[guess.row][guess.col] == 0 {
			continue
		}

		var tmp = (*b).board[guess.row][guess.col]
		(*b).board[guess.row][guess.col] = int8(0)

		// if any other choice then tmp is valid then
		// we can't remove tmp

		for _, num := range numline_int8(1, 10) {
			if num == tmp {
				continue
			}
			if valid_move(num, guess, b.idx, &b.board) {
				// if true restore it and break
				(*b).board[guess.row][guess.col] = tmp
				break
			}
		}
		if (*b).board[guess.row][guess.col] == 0 {
			left--
		}

		if pointer >= 80 && left > keep {
			pointer = 0
			if (reference - left) < 5 {
				reference = -1
			} else {
				reference = left
			}
		}
	}
	if left == keep {
		return true
	} else {
		return false
	}
}

func generate(keep int, idx *indexmap) (b *board) {
	for {
		b = make_board(idx)

		//b = tmp
		make_valid_board(b)

		b.solution = deep_copy(b.board)
		break
	}
	return b
}

func process(keep int, idx *indexmap, ch chan<- *sudoku, wg *sync.WaitGroup) {
	defer wg.Done()
	var b = generate(keep, idx)
	ch <- &sudoku{
		solved:   b.solution,
		unsolved: b.board,
	}
}

func process_single(keep int, idx *indexmap) sudoku {
	// defer wg.Done()
	// var b = make_board(idx)
	var b = generate(keep, idx)

	// ch <- [2][9][9]int8
	//out[0] = &b.board
	//out[1] = &b.solved

	return sudoku{
		solved:   b.solution,
		unsolved: b.board,
	}
}
