package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Option int

const (
	None Option = iota
	Count
)

var (
	ErrNoOption = errors.New("No option found")
)

func main() {
	opt, filename, err := getParams()
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
	}

	err = runOption(opt, filename)

	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		os.Exit(2)
	}
}

func getParams() (Option, string, error) {
	filename := flag.String("c", "", "-c <filename>")

	flag.Parse()

	if len(*filename) != 0 {
		return Count, *filename, nil
	}

	fmt.Println(strings.Join(os.Args, " "))
	fmt.Println("The correct use is `ccwc -c <filename>`.")

	return None, "", nil
}

func runOption(opt Option, filename string) error {
	switch opt {
	case Count:
		return countBytes(filename)
	default:
		return ErrNoOption
	}
}

func countBytes(filename string) error {
	f, err := os.ReadFile(filename)
	if err != nil {
		return err
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
	return nil
}
