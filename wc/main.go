package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	var fn bufio.SplitFunc
	switch {
	case *lines && *bytes:
		fmt.Println("Cannot use both -l and -b")
		os.Exit(1)
	case *lines:
		fn = bufio.ScanLines
	case *bytes:
		fn = bufio.ScanBytes
	default:
		fn = bufio.ScanWords
	}
	wc, err := count(os.Stdin, fn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(wc)
}

func count(r io.Reader, fn bufio.SplitFunc) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(fn)
	wc := 0
	for scanner.Scan() {
		wc++
	}
	if err := scanner.Err(); err != nil {
		return wc, err
	}
	return wc, nil
}
