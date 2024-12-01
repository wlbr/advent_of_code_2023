package main

import (
	"testing"
)

type testdata struct {
	fname    string
	expected int
}

var testset1 []*testdata = []*testdata{{"example1.txt", 142}}
var testset2 []*testdata = []*testdata{{"example2.txt", 281}}

func TestTaskOne(t *testing.T) {

	for _, test := range testset1 {
		r := solve(test.fname, false)
		if r != test.expected {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expected)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset2 {
		r := solve(test.fname, true)
		if r != test.expected {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expected)
		}
	}
}
