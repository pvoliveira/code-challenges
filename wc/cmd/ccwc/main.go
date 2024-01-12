package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
)

type Option int

const (
	None Option = iota
	Bytes
	Lines
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
		os.Exit(1)
	}
}

func getParams() (Option, string, error) {
	fileToBytes := flag.String("c", "", "-c <filename>")
	fileToLines := flag.String("l", "", "-l <filename>")

	flag.Parse()

	if len(*fileToBytes) != 0 {
		return Bytes, *fileToBytes, nil
	}

	if len(*fileToLines) != 0 {
		return Lines, *fileToLines, nil
	}

	return None, "", nil
}

func runOption(opt Option, filename string) error {
	switch opt {
	case Bytes:
		return countBytes(filename)
	case Lines:
		return countLines(filename)
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

func countLines(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)

	lines := 0
	for s.Scan() {
		lines += 1
	}

	fmt.Printf("%v\t%s\n", lines, filename)
	return nil
}
