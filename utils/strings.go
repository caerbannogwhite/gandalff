package utils

import "fmt"

func Truncate(s string, n int) string {
	if len(s) > n {
		return s[:n-3] + "..."
	}
	return s
}

func Center(s string, n int) string {
	r := []rune(s)
	l := len(r)
	if l >= n {
		return s
	}

	n += len(s) - l

	left := (n + l) / 2
	right := n - l - left - 1

	return fmt.Sprintf("%*s%*s", left, string(s), right, "")
}
