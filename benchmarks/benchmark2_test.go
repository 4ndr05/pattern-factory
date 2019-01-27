package benchmarks

import (
	"pattern-factory/algorithms"
	"testing"
)

var (
	test = []string{"a", "b", "c", "d", "e"}
)

func BenchmarkJoinCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = algorithms.Join(test)
	}
}

func BenchmarkSprintfCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = algorithms.Sprintf(test)
	}
}

func BenchmarkConcatCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = algorithms.Concat(test)
	}
}

func BenchmarkConcatOneLineCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = algorithms.ConcatOneLine(test)
	}
}

func BenchmarkBufferCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = algorithms.Buffer(test)
	}
}

func BenchmarkBufferWithResetCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = algorithms.BufferWithReset(test)
	}
}

func BenchmarkBufferFprintfCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = algorithms.BufferFprintf(test)
	}
}

func BenchmarkBufferStringBuilderCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = algorithms.BufferStringBuilder(test)
	}
}
