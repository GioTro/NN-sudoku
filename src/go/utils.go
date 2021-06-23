package main

import (
	"fmt"
	"math/rand"
	"time"
)

func check_equality(a *[][]int8, b *[][]int8) bool {
	for i := range *a {
		for j := range (*a)[i] {
			if (*a)[i][j] != (*b)[i][j] {
				return false
			}
		}
	}
	return true
}

func deep_copy(source [][]int8) [][]int8 {
	var copy = make2dint8(9)
	for i, arr := range source {
		for j, num := range arr {
			(copy)[i][j] = num
		}
	}
	return copy

}

func make2dint8(length int) [][]int8 {
	var out = make([][]int8, length)
	for idx := range out {
		out[idx] = make([]int8, length)
	}
	return out
}

func numline(from, to int) []int {
	var out = make([]int, (to - from))
	for i := range out {
		out[i] = from + i
	}
	return out
}

func numline_int8(from, to int) []int8 {
	var out = make([]int8, to-from)
	for i := range out {
		out[i] = int8(from + i)
	}
	return out
}

func rand_idx(length int) []int {
	var index_arr = numline(0, length)

	var swap = func(from int, arr []int) {
		var to = rand.Intn(len(arr))

		if to == from {
			if to > 1 {
				to--
			} else {
				to++
			}
		}

		var tmp = arr[from]
		arr[from] = arr[to]
		arr[to] = tmp
	}

	// Shuffle
	for idx := range index_arr {
		swap(idx, index_arr)
	}
	return index_arr
}

func shuffled_int8(arr []int8) (out []int8) {
	out = make([]int8, len(arr))
	for idx, ridx := range rand_idx(len(arr)) {
		out[idx] = arr[ridx]
	}
	return out
}

func shuffled_int(arr []int) (out []int) {
	out = make([]int, len(arr))
	for idx, ridx := range rand_idx(len(arr)) {
		out[idx] = arr[ridx]
	}
	return out
}

func shuffled_index(arr []index) (out []index) {
	out = make([]index, len(arr))
	for idx, ridx := range rand_idx(len(arr)) {
		out[idx] = arr[ridx]
	}
	return out
}

func pretty_print(b *[][]int8) {
	for i := range *b {
		for j := range (*b)[i] {
			fmt.Print((*b)[i][j], " ")
			if (j+1)%3 == 0 {
				fmt.Print("  ")
			}
		}
		fmt.Println()
		if (i+1)%3 == 0 {
			fmt.Println()
		}
	}
}

func load_data() {
	return
}

func save_data() {
	return
}

func preprocess() {
	return
}

func test() {
	var tot = time.Now()

	var arr [8][8]int8

	var t1 = time.Since(tot).Nanoseconds()
	var start = time.Now()

	var arr2 = make2dint8(8)

	var t2 = time.Since(start).Nanoseconds()

	start = time.Now()

	var arr3 = deep_copy(arr2)

	var t4 = time.Since(start).Nanoseconds()
	start = tot
	var tot2 = time.Since(start).Nanoseconds()

	var flt = [5]float64{float64(tot2), float64(t1), float64(t2), float64(t4)}

	fmt.Println(flt[0], flt[1]/flt[0], flt[2]/flt[0], flt[3]/flt[0], flt[4]/flt[0])

	arr3 = deep_copy(arr3)
	if arr[0][0] == int8(0) {
		return
	}
}
