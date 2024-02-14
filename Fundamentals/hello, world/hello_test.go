package main

import "testing"

/*
needs to be in a file with a name like xxx_test.go
test function must start with the word `Test`
test function takes one argument only t *testing.T with import "testing"
*/
func TestHello(t *testing.T) {
	// `go test -v` prints the full output
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Mark")
		want := "Hello, Mark"

		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, Word"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper() // comment this out and modified sub-test to fail, it help developer to spot bug
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
