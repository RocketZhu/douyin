package middleware

import "testing"

// 7c4a8d09ca3762af61e59520943dc26494f8941b
func TestSHA1(t *testing.T) {
	print(CalculateSHA1("12345"))
}
