package main

import (
	"os"
)

func main() {
	s, sep := "", ""

	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = ""
		i++
		println(i, " - ", os.Args[i])
	}

}
