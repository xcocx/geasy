package random

import "testing"

func TestInt(t *testing.T) {
	if res := Int(100, 200); res < 100 || res > 200 {
		t.Errorf(`Int(100, 200) = %d, want between 100 ~ 200`, res)
	}
}

func TestInt64(t *testing.T) {
	if res := Int(1000, 2000); res < 1000 || res > 2000 {
		t.Errorf(`Int(1000, 2000) = %d, want between 1000 ~ 2000`, res)
	}
}

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(100, 200)
	}
}

func BenchmarkInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int64(1000, 2000)
	}
}
