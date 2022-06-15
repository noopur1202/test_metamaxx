package main

import (
	"testing"
)

type Num struct {
	Num1 int
	Num2 int
}

func TestSubtract(t *testing.T) {
	s := []Num{{2, 1}, {4, 3}, {6, 5}}

	for _, v := range s {
		result := Subtract(v.Num1, v.Num2)
		testresult := v.Num1 - v.Num2
		if result != testresult {
			t.Errorf("The function is not subtracting correctly.")
		}
	}
}

func TestAdd(t *testing.T) {
	a := []Num{{1, 2}, {3, 4}, {5, 6}}

	for _, v := range a {
		result := Add(v.Num1, v.Num2)
		testresult := v.Num1 + v.Num2
		if result != testresult {
			t.Errorf("The function is not adding correctly.")
		}
	}
}
