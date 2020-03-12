package base69

import (
	"strings"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/-*<>|"

func intToBytes(n int) []byte {
	i := n % 69
	j := n / 69
	return []byte{charset[i], charset[j]}
}

func bytesToInt(b []byte) int {
	i := strings.Index(charset, string(b[0]))
	j := strings.Index(charset, string(b[1]))
	return 69*j + i
}
