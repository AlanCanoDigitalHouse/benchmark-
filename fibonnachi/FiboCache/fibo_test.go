package FiboCache
import "testing"
var result int
func BenchmarkFibComplete(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
			// always record the result of Fib to prevent
			// the compiler eliminating the function call.
			r = fib(10)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}