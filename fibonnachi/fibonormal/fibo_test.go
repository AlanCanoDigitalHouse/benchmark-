package fibonormal

import "testing"

func TestRecursiveFibonacci(t *testing.T) {
	data := []struct {
		n    uint
		want uint
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{10, 55},
		{42, 267914296},
	}
	for _, d := range data {
		if got := RecursiveFibonacci(d.n); got != d.want {
			t.Errorf("got: %d, want: %d", got, d.want)
		}
	}
}
