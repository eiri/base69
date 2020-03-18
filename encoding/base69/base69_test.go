package base69

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var intPairs = []struct {
	in  int
	out []byte
}{
	{0, []byte("AA")},
	{1, []byte("BA")},
	{68, []byte("|A")},
	{69, []byte("AB")},
	{1890, []byte("bb")},
	{3640, []byte("00")},
	{4340, []byte("++")},
	{4760, []byte("||")},
}

func TestIntToBytes(t *testing.T) {
	for _, p := range intPairs {
		got := intToBytes(p.in)
		if !reflect.DeepEqual(got, p.out) {
			t.Errorf("intToBytes(%d) = %q, want %q", p.in, got, p.out)
		}
	}
}

func TestBytesToInt(t *testing.T) {
	for _, p := range intPairs {
		got := bytesToInt(p.out)
		if got != p.in {
			t.Errorf("bytesToInt(%q) = %d, want %d", p.out, got, p.in)
		}
	}
}

var headPairs = []struct {
	in, out []byte
}{
	{[]byte(""), []byte("")},
	{[]byte("a"), []byte("")},
	{[]byte("abcdef"), []byte("")},
	{[]byte("abcdefg"), []byte("wATBHB2AjAVAHBiB")},
	{[]byte("abcdefgh"), []byte("wATBHB2AjAVAHBiB")},
	{[]byte("abcdefg0123456"), []byte("wATBHB2AjAVAHBiBYAMAmAjAZALBlB2A")},
}

func TestEncodeHead(t *testing.T) {
	for _, p := range headPairs {
		got := encodeHead(p.in)
		if string(got) != string(p.out) {
			t.Errorf("encodeHead(%q) = %q, want %q", p.in, got, p.out)
		}
	}
}

var chunkPairs = []struct {
	in, out []byte
}{
	{[]byte("wA-AAAAAAAAAAA6="), []byte("a\x00\x00\x00\x00\x00\x00")},
	{[]byte("wATBHB2AjAVAHB1="), []byte("abcdef\x00")},
	{[]byte("wATBHB2AjAVAHBiB"), []byte("abcdefg")},
	{[]byte("wATBHB2AjAVAHBiBYAMAmAjAZALBlB2A"), []byte("abcdefg")},
}

func TestDecodeChunk(t *testing.T) {
	for _, p := range chunkPairs {
		got := decodeChunk(p.in)
		if string(got) != string(p.out) {
			t.Errorf("decodeChunk(%q) = %q, want %q", p.in, got, p.out)
		}
	}
}

var pairs = []struct {
	in, out []byte
}{
	{[]byte(""), []byte("")},
	{[]byte("a"), []byte("wA-AAAAAAAAAAA6=")},
	{[]byte("abcdef"), []byte("wATBHB2AjAVAHB1=")},
	{[]byte("abcdefg"), []byte("wATBHB2AjAVAHBiB")},
	{[]byte("abcdefgh"), []byte("wATBHB2AjAVAHBiB0AAAAAAAAAAAAA6=")},
	{[]byte("abcdefg0123456"), []byte("wATBHB2AjAVAHBiBYAMAmAjAZALBlB2A")},
	{[]byte("brown fox"), []byte("xAcAIByB7A4A-AhB3AZBAAAAAAAAAA5=")},
}

func TestEncode(t *testing.T) {
	for _, p := range pairs {
		got := Encode(p.in)
		if string(got) != string(p.out) {
			t.Errorf("Encode(%q) = %q, want %q", p.in, got, p.out)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, p := range pairs {
		got := Decode(p.out)
		if string(got) != string(p.in) {
			t.Errorf("Decode(%q) = %q, want %q", p.out, got, p.in)
		}
	}
}

func TestEncodeDecode(t *testing.T) {
	for l := 6; l <= 69; l += 7 {
		testBytes := make([]byte, l)
		rand.Seed(time.Now().UnixNano())
		rand.Read(testBytes)
		encoded := Encode(testBytes)
		if reflect.DeepEqual(encoded, testBytes) {
			t.Error("Encode output is equal to input")
		}
		decoded := Decode(encoded)
		if !reflect.DeepEqual(decoded, testBytes) {
			t.Error("Decode can't decode encoded bin back")
		}
	}
}
