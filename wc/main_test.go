package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4
	got, err := count(b, bufio.ScanWords)
	if err != nil {
		t.Error(err)
	}
	if got != exp {
		t.Errorf("Expected %d, got %d", exp, got)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	exp := 3
	got, err := count(b, bufio.ScanLines)
	if err != nil {
		t.Error(err)
	}
	if got != exp {
		t.Errorf("Expected %d, got %d", exp, got)
	}
}
func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("12345")
	exp := 5
	got, err := count(b, bufio.ScanBytes)
	if err != nil {
		t.Error(err)
	}
	if got != exp {
		t.Errorf("Expected %d, got %d", exp, got)
	}
}
