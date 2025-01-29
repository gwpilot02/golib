package golib

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("1 + 2 != 3")
	}
}

func TestSub(t *testing.T) {
	if Sub(2, 1) != 1 {
		t.Error("2 - 1 != 1")
	}
}
