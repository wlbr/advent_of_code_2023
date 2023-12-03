package main

import (
	"testing"
)

type testdata struct {
	fname     string
	expected1 int
	expected2 int
}

var testset []*testdata = []*testdata{{"example1.txt", 4361, 467835}, {"example2.txt", 4361, 467835}}

func TestTaskOne(t *testing.T) {
	for _, test := range testset {
		r := solve1(test.fname)
		if r != test.expected1 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expected1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		r := solve2(test.fname)
		if r != test.expected2 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expected2)
		}
	}
}
