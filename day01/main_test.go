package main

import (
	"testing"
)

type testdata struct {
	fname    string
	expected int
}

var testset1 []*testdata = []*testdata{{"example.txt", 11}}
var testset2 []*testdata = []*testdata{{"example.txt", 31}}

func TestTaskOne(t *testing.T) {

	for _, test := range testset1 {
		r := solve1(test.fname)
		if r != test.expected {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expected)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset2 {
		r := solve2(test.fname)
		if r != test.expected {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expected)
		}
	}
}
