package main

import "testing"

func TestFactory(t *testing.T) {
	f := Factory{Blue, Orange}

	if f.Count(Blue) != 1 {
		t.Error("{Blue, Orange} factory does not contain Blue")
	}
	if f.Count(Orange) != 1 {
		t.Error("{Blue, Orange} factory does not cotain Orange")
	}
	if f.Count(Black) != 0 {
		t.Error("{Blue, Orange} factory cotains Black")
	}

	if !f.Take(Orange) {
		t.Error("Cannot take Orange from {Blue,Orange} factory")
	}

	if f.Count(Orange) != 0 {
		t.Error("Count of Orange in {Blue,Orange} factory after taking Orange is not zero")
	}
}
