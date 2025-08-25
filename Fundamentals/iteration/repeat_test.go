package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// When the benchmark code is executed, it runs b.N times and measures how long it takes.
	// use `go test -bench=.`
	// use `go test -bench=. -benchmem`, reports information about memory allocations
	for b.Loop() {
		Repeat("a")
	}
	/*
		output: 16066174                72.50 ns/op
		72.50 ns/op means is our fn takes on average 72.5 nanosecond to run (locally)
		and to get this result it ran 16,066,174 times
		default benchmarks are run sequentially
	*/

	/* output with -benchmem
	10000000           25.70 ns/op           8 B/op           1 allocs/op
	B/op: the number of bytes allocated per iteration (less is better)
	allocs/op: the number of memory allocations per iteration (less is better)
	*/
}

func ExampleRepeat() {
	repeat := Repeat("b")
	fmt.Println(repeat)
	// output: bbbbb
}
