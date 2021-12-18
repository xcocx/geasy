package encrypt

import "testing"

func TestSHA1(t *testing.T) {
	input := "hello world"
	want := "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed"
	if res := SHA1(input); res != want {
		t.Errorf("SHA1(%q) = %s, want %s", input, res, want)
	}
}

func BenchmarkSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SHA1("hello world")
	}
}
