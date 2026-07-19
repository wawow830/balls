package sequence

import "testing"

func BenchmarkGeneratorNext(b *testing.B) {
	generator := New(42)
	b.ReportAllocs()
	for b.Loop() {
		_ = generator.Next()
	}
}
