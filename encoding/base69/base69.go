package base69

import (
	"fmt"
	"strconv"
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

func decodeChunk(src []byte) []byte {
	paddedBytes := src[len(src)-1] == 61
	decoded := make([]byte, 8)
	for i := 0; i < 8; i++ {
		if i == 7 && paddedBytes {
			decoded[i] = 0x0
		} else {
			k, l := i*2, i*2+2
			decoded[i] = byte(bytesToInt(src[k:l]))
		}
	}
	dst := make([]byte, 7)
	for i := 0; i < 7; i++ {
		t1 := decoded[i] << (i + 1)
		t2 := decoded[i+1] >> (7 - i - 1)
		dst[i] = t1 | t2
	}
	return dst
}

func Encode(src []byte) []byte {
	dst := encodeHead(src)
	extraBytes := len(src) % 7
	if extraBytes > 0 {
		extra := make([]byte, 7)
		copy(extra, src[(len(src)-extraBytes):])
		tail := encodeHead(extra)
		dst = append(dst, tail...)
		copy(dst[len(dst)-2:], fmt.Sprintf("%d=", 7-extraBytes))
	}
	return dst
}

func Decode(src []byte) []byte {
	if len(src) == 0 {
		return src
	}
	extraBytes := 0
	if src[len(src)-1] == 61 {
		e, _ := strconv.Atoi(fmt.Sprintf("%c", src[len(src)-2]))
		extraBytes += e
	}
	chunkCount := len(src) / 16
	dst := make([]byte, 0)
	for i := 0; i < chunkCount; i++ {
		chunk := src[i*16 : (i+1)*16]
		decoded := decodeChunk(chunk)
		if extraBytes > 0 && (i == chunkCount-1) {
			dst = append(dst, decoded[:7-extraBytes]...)
		} else {
			dst = append(dst, decoded...)
		}
	}
	return dst
}
