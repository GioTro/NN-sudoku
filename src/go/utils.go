package main

import "math/rand"

func numline(from, to int) []int {
	var out = make([]int, (to - from))
	var count int

	for i := from; i < to; i++ {
		out[count] = int(i)
		count++
	}

	return out
}

func numline_int8(from, to int) []int8 {
	var out = make([]int8, to-from)
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
	for idx, ridx := range rand_idx(len(arr)) {
		out[idx] = arr[ridx]
	}
	return out
}

func shuffled_int(arr []int) (out []int) {
	for idx, ridx := range rand_idx(len(arr)) {
		out[idx] = arr[ridx]
	}
	return out
}

func shuffled_index(arr []index) (out []index) {
	for idx, ridx := range rand_idx(len(arr)) {
		out[idx] = arr[ridx]
	}
	return out
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
