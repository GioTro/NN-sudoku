package main

import (
	"fmt"
	"strconv"
)

func pretty_print(b board) {
	var s string
	for idx := range b {
		if idx%9 == 0 {
			s += "\n"
			if idx%27 == 0 {
				s += "\n"
			}
		}
		if idx%3 == 0 {
			s += "   "
		}
		s += strconv.Itoa(int(b[idx])) + " "
	}
	fmt.Println(s)
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

// func test() {
// 	var tot = time.Now()

// 	var arr [8][8]int8

// 	var t1 = time.Since(tot).Nanoseconds()
// 	var start = time.Now()

// 	var arr2 = make2dint8(8)

// 	var t2 = time.Since(start).Nanoseconds()

// 	start = time.Now()

// 	var arr3 = deep_copy(arr2)

// 	var t4 = time.Since(start).Nanoseconds()
// 	start = tot
// 	var tot2 = time.Since(start).Nanoseconds()

// 	var flt = [5]float64{float64(tot2), float64(t1), float64(t2), float64(t4)}

// 	fmt.Println(flt[0], flt[1]/flt[0], flt[2]/flt[0], flt[3]/flt[0], flt[4]/flt[0])

// 	arr3 = deep_copy(arr3)
// 	if arr[0][0] == int8(0) {
// 		return
// 	}
// }
