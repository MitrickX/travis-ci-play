package main

import "testing"

func TestSumSlice1(t *testing.T) {
	res := SumSlice([]int{1,2,3,4,5})
	expected := 15
	if expected != res {
		t.Fatalf("unexpected result %d insteadof %d", res, expected)
	}
}

func TestSumSlice2(t *testing.T) {
	res := SumSlice([]int{1,2,3,4,5,6})
	expected := 21
	if expected != res {
		t.Fatalf("unexpected result %d insteadof %d", res, expected)
	}	
}

func TestSum1(t *testing.T) {
	res := Sum(1,2,3,4,5)
	expected := 15
	if expected != res {
		t.Fatalf("unexpected result %d insteadof %d", res, expected)
	}
}

func TestSum2(t *testing.T) {
	res :=   Sum(1,2,3,4,5,6)
	expected := 21
	if expected != res {
		t.Fatalf("unexpected result %d insteadof %d", res, expected)
	}
}
