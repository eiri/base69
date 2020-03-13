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

func encodeHead(src []byte) []byte {
	dst := make([]byte, 0)
	end := len(src) - len(src)%7
	for i := 0; i < end; i++ {
		shift := (i % 7) + 1
		shifted := src[i] >> shift
		if shift > 1 {
			pre := (src[i-1] & ((2 << (shift - 2)) - 1)) << (8 - shift)
			shifted = pre | shifted
		}
		chars := intToBytes(int(shifted))
		dst = append(dst, chars...)
		if shift == 7 {
			shifted = src[i] & 127
			chars = intToBytes(int(shifted))
			dst = append(dst, chars...)
		}
	}
	return dst
}
