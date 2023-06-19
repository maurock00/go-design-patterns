package main

import (
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	dat, _ := os.Open("./" + filename)

	io.Copy(os.Stdout, dat)
}
