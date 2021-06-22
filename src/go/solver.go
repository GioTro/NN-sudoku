package main

import (
	"fmt"
)

type board struct {
	board          [9][9]int8
	solved         [9][9]int8
	solution_count int
	idx            *indexmap
	get            func(int, int, string, *board) [9]int8
}

type indexmap struct {
	rmap, cmap [9][9]index
	sqrmap     [3][3][9]index
	all        [81]index
	square     func(int, int, [3][3][9]index) [9]index
	row        func(int, [9][9]index) [9]index
	col        func(int, [9][9]index) [9]index
}

type index struct {
	row, col int
}

func make_board(idx *indexmap) board {
	var b [9][9]int8
	var c [9][9]int8
	return board{
		board:          b,
		solved:         c,
		solution_count: 0,
		idx:            idx,
		get: func(row, col int, s string, b *board) [9]int8 {
			var idx = (*b).idx
			var out [9]int8
			var idxarr [9]index

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
				out[i] = (*b).board[idx.row][idx.col]
			}

			return out
		},
	}
}

func make_indexmap() indexmap {
	var rows [9][9]index
	var cols [9][9]index
	var all [81]index
	var squares [3][3][9]index

	// make row and col map
	var count int
	for i := 0; i < 10; i++ {
		var r [9]index
		var c [9]index
		for j := 0; j < 10; j++ {
			r[i] = index{row: i, col: j}
			c[i] = index{row: j, col: i}
			all[count] = index{row: i, col: j}
			count++
		}
		rows[i] = r
		cols[i] = c
	}

	for i := 0; i < 9; i++ {
		var s [9]index
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

		square: func(row int, col int, sm [3][3][9]index) [9]index {
			return sm[row][col]
		},

		row: func(idx int, rm [9][9]index) [9]index {
			return rm[idx]
		},

		col: func(idx int, cm [9][9]index) [9]index {
			return cm[idx]
		},
	}
}

func valid_board(b *board) bool {
	var nline = numline_int8(1, 10)

	// Helper function
	var member = func(a [9]int8, b []int8) bool {
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
		if !member((*b).get(row, 0, "row", b), nline) {
			return false
		}
	}
	for col := 0; col < 9; col++ {
		if !member((*b).get(0, col, "col", b), nline) {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !member((*b).get(i, j, "sqr", b), nline) {
				return false
			}
		}
	}
	// passed all checks
	return true
}

func valid_move(b *board, num int8, idx index) bool {
	if num == 0 {
		return false
	}

	// row
	for _, val := range (*b).get(idx.row, idx.col, "row", b) {
		if num == val {
			return false
		}
	}

	// col
	for _, val := range b.get(idx.row, idx.col, "col", b) {
		if num == val {
			return false
		}
	}
	// sqr
	for _, val := range b.get(idx.row, idx.col, "sqr", b) {
		if num == val {
			return false
		}
	}
	// passed all checks
	return true
}

func recursive(limit int, b *board) {
	/*
		Recursively solves the board via backtracking.
		Solves the board in place

		limit number of solution to find. This is going to be
		either 1 for find first solution or 2 to check if the
		solution is unique.
	*/
	if !((*b).solution_count < limit) {
		return
	}
	var aidx = (*b).index.all
	for _, ridx := range shuffled(len((*b).idx.all)) {
		var row = aidx[ridx].row
		var col = aidx[ridx].col
		if (*b).board[row][col] == 0 && (*b).solution_count < limit {
			for _, num := range numline_int_eight(1, 10, true) {
				if !((*b).solution_count < limit) || !valid_move(b, int8(num), aidx[ridx]) {
					continue
				}

				(*b).board[row][col] = num
				recursive(limit, b)

				// Check this statement if this would be correct
				(*b).board[row][col] = int8(0) // back track
			}

		}
	}
	if valid_board(b) {
		(*b).solution_count++
	}
	return
}

// findfirst solution, returns true if
// there was a solution to be found
func find_first(b *board) bool {
	recursive(1, b)

	if (*b).solution_count == 0 {
		return false
	} else {
		return true
	}
}

// Alias
func find_more(b *board) bool {
	recursive(2, b)
	if (*b).solution_count > 1 {
		return true
	} else {
		return false
	}
}

func deep_copy(source [9][9]int8) (copy [9][9]int8) {
	for i, arr := range source {
		for j, num := range arr {
			copy[i][j] = num
		}
	}
	return copy

}

/*
	Try removing random points
*/

func pluck(limit int, b *board) bool {
	var plucked int
	for _, idx := range shuffled_index((*b).idx.all) {
		if (*b).board[idx.row][idx.col] != 0 {
			// try
			var tmp = (*b).board[idx.row][idx.col]
			(*b).board[idx.row][idx.col] = 0

			// if there is more than one solution
			if find_more(b) {
				// restore value
				(*b).board[idx.row][idx.col] = tmp
			} else {
				plucked++
			}
		}
	}
	if plucked > limit {
		return false
	} else {
		return true
	}
}

func generate(limit int, b *board) {
	for true {
		// Its faster to start over
		// compared to backtracking
		// if pluck fails start over

		var new [9][9]int8
		(*b).board = new

		if !find_first(b) {
			continue
		}

		(*b).solved = deep_copy((*b).board)

		if pluck(limit, b) {
			break
		}
	}
}

func process(limit int, idx *indexmap) (*[9][9]int8, *[9][9]int8) {
	// defer wg.Done()

	b = make_board(idx)
	generate(limit, &b)

	// ch <- [2][9][9]int8

	return &b.board, &b.solved
}
