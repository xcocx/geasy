package random

import "testing"

func TestString(t *testing.T) {
	if res := String(10); len(res) != 10 {
		t.Errorf("String(10) = %s, want len is 10", res)
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(10)
	}
}
