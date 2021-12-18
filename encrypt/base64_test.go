package encrypt

import "testing"

func TestBase64Encode(t *testing.T) {
	input := "hello world"
	want := "aGVsbG8gd29ybGQ="
	if res := Base64Encode(input); res != want {
		t.Errorf("Base64Encode(%q) = %s, want %s", input, res, want)
	}
}

func TestBase64Decode(t *testing.T) {
	input := "aGVsbG8gd29ybGQ="
	want := "hello world"
	if res, _ := Base64Decode(input); res != want {
		t.Errorf("Base64Decode(%q) = %s, want %s", input, res, want)
	}
}

func BenchmarkBase64Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Base64Encode("hello world")
	}
}

func BenchmarkBase64Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Base64Encode("aGVsbG8gd29ybGQ=")
	}
}
