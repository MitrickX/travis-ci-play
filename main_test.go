package main

import "testing"

func TestSumSlice(t *testing.T) {
	res := SumSlice([]int{1,2,3,4,5})
	expected := 15
	if expected != res {
		t.Fatalf("unexpected result %d insteadof %d", res, expected)
	}
}

func TestSum(t *testing.T) {
	res := Sum(1,2,3,4,5)
	expected := 15
	if expected != res {
		t.Fatalf("unexpected result %d insteadof %d", res, expected)
	}
}