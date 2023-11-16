package main

import (
	"bar"
	"foo"
)

func main() {}

func Cmp(f *foo.Foo, b *bar.Bar) bool {
	if f.A != b.A {
		return false
	}

	if f.B != b.B { // want "pointer comparison: f.B != b.B"
		return false
	}

	if *f.B != *b.B {
		return false
	}

	var eq bool
	eq = f.A == b.A
	if !eq {
		return false
	}

	eq = f.B == b.B // want "pointer comparison: f.B == b.B"
	if !eq {
		return false
	}

	eq = *f.B == *b.B
	if !eq {
		return false
	}

	return true
}

func cmpFoos(a, b *foo.Foo) bool {
	if a != b { // want "pointer comparison: a != b"
		return false
	}

	if *a != *b {
		return false
	}

	if a.A != b.A {
		return false
	}

	if a.B != b.B { // want "pointer comparison: a.B != b.B"
		return false
	}

	if *a.B != *b.B {
		return false
	}

	return true
}

func cmpBars(a, b *bar.Bar) bool {
	if a != b { // want "pointer comparison: a != b"
		return false
	}

	if *a != *b {
		return false
	}

	if a.A != b.A {
		return false
	}

	if a.B != b.B { // want "pointer comparison: a.B != b.B"
		return false
	}

	if *a.B != *b.B {
		return false
	}

	return true
}
