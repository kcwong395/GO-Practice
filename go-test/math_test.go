package main

import (
	"testing"
)

/*
	=== RUN   TestAdd
	--- PASS: TestAdd (0.00s)
*/
func TestAdd(t *testing.T) {
	expected := 10
	a := 4
	b := 6
	actual := Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected %d", a, b, actual, expected)
	}
}

/*
	=== RUN   TestSub
	--- PASS: TestSub (0.00s)
*/
func TestSub(t *testing.T) {
	expected := -2
	a := 4
	b := 6
	actual := Sub(a, b)
	if actual != expected {
		t.Errorf("Sub(%d, %d) = %d; expected %d", a, b, actual, expected)
	}
}

/*
	=== RUN   TestMulti
		math_test.go:33: Multi(4, 6) = 0; expected 24
	--- FAIL: TestMulti (0.00s)
*/
func TestMulti(t *testing.T) {
	expected := 24
	a := 4
	b := 6
	actual := WrongMulti(a, b)
	if actual != expected {
		t.Errorf("Multi(%d, %d) = %d; expected %d", a, b, actual, expected)
	}
}
