package main

type board struct {
	board [9][9]int8
	idx   *idxmap
	get   func([9]idx, [9][9]int8) [9]int8
}

type idxmap struct {
	rmap, cmap [9][9]idx
	sqrmap     [3][3][9]idx
	square     func(int, int, [3][3][9]idx) (out [9]idx)
	row        func(int, [9][9]idx) (out [9]idx)
	col        func(int, [9][9]idx) (out [9]idx)
}

type idx struct {
	row, col int
}

func valid_board(b board) bool {
	var numline [9]int8
	for i := 1; i < 10; i++ {
		numline[i-1] = int8(i)
	}

	var member = func(a [9]int8, b [9]int8) bool {
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
		if !member(b.get((*b.idx).row(row, (*b.idx).rmap), b.board), numline) {
			return false
		}
	}
	for col := 0; col < 9; col++ {
		if !member(b.get((*b.idx).col(col, (*b.idx).cmap), b.board), numline) {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !member(b.get((*b.idx).square(i, j, (*b.idx).sqrmap), b.board), numline) {
				return false
			}
		}
	}
	// passed all checks
	return true
}

func makeidxmap() *idxmap {
	var rows [9][9]idx
	var cols [9][9]idx
	var squares [3][3][9]idx

	// make row and col map
	for i := 0; i < 10; i++ {
		var r [9]idx
		var c [9]idx
		for j := 0; j < 10; j++ {
			r[i] = idx{row: i, col: j}
			c[i] = idx{row: j, col: i}
		}
		rows[i] = r
		cols[i] = c
	}

	for i := 0; i < 9; i++ {
		var s [9]idx
		var count int
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				orow := 3*(i%3) + j
				ocol := 3*(i/3) + k
				s[count] = idx{row: orow, col: ocol}
				count++
			}
		}
		squares[(i % 3)][(i / 3)] = s
	}

	var sqr = func(row int, col int, sm [3][3][9]idx) [9]idx {
		return sm[row][col]
	}

	var row = func(index int, rm [9][9]idx) [9]idx {
		return rm[index]
	}
	var col = func(index int, cm [9][9]idx) [9]idx {
		return cm[index]
	}

	var out = idxmap{
		rmap:   rows,
		cmap:   cols,
		sqrmap: squares,
		square: sqr,
		row:    row,
		col:    col,
	}
	return &out
}

func valid_move(b board, num int8, idx idx) bool {
	if num == 0 {
		return false
	}

	// row
	for _, val := range b.get((*b.idx).row(idx.row, (*b.idx).rmap), b.board) {
		if num == val {
			return false
		}
	}

	// col
	for _, val := range b.get((*b.idx).row(idx.col, (*b.idx).cmap), b.board) {
		if num == val {
			return false
		}
	}
	// sqr
	for _, val := range b.get((*b.idx).square(idx.row, idx.col, (*b.idx).sqrmap), b.board) {
		if num == val {
			return false
		}
	}
	// passed all checks
	return true
}

func solve() {
	return
}

func generate() {
	return
}
