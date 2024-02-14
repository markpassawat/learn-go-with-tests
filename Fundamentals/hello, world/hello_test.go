package main

import "testing"

/*
needs to be in a file with a name like xxx_test.go
test function must start with the word `Test`
test function takes one argument only t *testing.T with import "testing"
*/
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
