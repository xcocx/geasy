package encrypt

import "testing"

func TestSHA256(t *testing.T) {
	input := "hello world"
	want := "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
	if res := SHA256(input); res != want {
		t.Errorf("SHA256(%q) = %s, want %s", input, res, want)
	}
}

func BenchmarkSHA256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SHA256("hello world")
	}
}
