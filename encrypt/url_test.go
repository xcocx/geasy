package encrypt

import "testing"

func TestURLEncode(t *testing.T) {
	input := "name=Jone&age=20"
	want := "name%3DJone%26age%3D20"
	if res := URLEncode(input); res != want {
		t.Errorf("URLEncode(%q) = %s, want %s", input, res, want)
	}
}

func TestURLDecode(t *testing.T) {
	input := "name%3DJone%26age%3D20"
	want := "name=Jone&age=20"
	if res, _ := URLDecode(input); res != want {
		t.Errorf("URLDecode(%q) = %s, want %s", input, res, want)
	}
}

func BenchmarkURLEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		URLEncode("hello world")
	}
}

func BenchmarkURLDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = URLDecode("hello%20world")
	}
}
