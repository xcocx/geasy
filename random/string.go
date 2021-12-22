package random

import "math/rand"

const charset = "abcdefghigklmnopqrstuvwxyzABCDEFGHIGKLMNOPQRSTUVWXYZ0123456789"

func String(length int) string {
	buf := make([]byte, 0, length)
	charsetLen := len(charset)
	for i := 0; i < length; i++ {
		index := rand.Intn(charsetLen)
		buf = append(buf, charset[index])
	}
	return string(buf)
}
