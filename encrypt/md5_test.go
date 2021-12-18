package encrypt

import (
	"testing"
)

func TestMD5(t *testing.T) {
	input := "hello world"
	want := "5eb63bbbe01eeed093cb22bb8f5acdc3"
	if res := MD5(input); res != want {
		t.Errorf("MD5(%q) = %s, want %s", input, res, want)
	}
}

func BenchmarkMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MD5("hello world")
	}
}
