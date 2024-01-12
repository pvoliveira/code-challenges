package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "-c" {
		fmt.Println(strings.Join(os.Args, " "))
		fmt.Println("The correct use is `ccwc -c <filename>`.")
		os.Exit(1)
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Println("------------------------")
		fmt.Println(os.Args[0])
		fmt.Println(os.Args[1])
		fmt.Println(os.Args[2])
		fmt.Println("------------------------")
	}

	filename := os.Args[2]

	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Errorf("\nError: %v\n", err)
		os.Exit(2)
	}
	size := len(f)

	/*
		// I guess we can use less memory
		// if we use a buffer and sum the number of bytes read
		buf := make([]byte, 512)
		size := 0
		for {
			s, err := f.Read(buf)
			size += s
			if err == io.EOF {
				break
			}
		}
	*/

	fmt.Printf("%v\t%s\n", size, filename)
}
