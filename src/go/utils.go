package main

import (
	"fmt"
	"os"
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

func save_data(inp []set, path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	for _, o := range inp {
		var buffer string
		for _, i := range o.solved {
			buffer += strconv.Itoa(int(i))
		}
		buffer += " "
		for _, i := range o.unsolved {
			buffer += strconv.Itoa(int(i))
		}
		buffer += "\n"
		f.WriteString(buffer)
	}
}

func load_data() {
	return
}
