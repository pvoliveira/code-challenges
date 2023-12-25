package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println(strings.Join(os.Args, " "))
		fmt.Println("The correct use is `wc <filename>`.")
		os.Exit(1)
	}

	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 512)
	size := 0
	for {
		s, err := f.Read(buf)
		size += s
		if err == io.EOF {
			break
		}
	}

	fmt.Printf("%v %s\n", size, filename)
}
