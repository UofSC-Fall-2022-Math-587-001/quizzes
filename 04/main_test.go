package main

import (
	"testing"
)

func Test1(t *testing.T) {
	got:= encode("abc")
	want := "bdf" 

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
