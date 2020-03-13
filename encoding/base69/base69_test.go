package base69

import (
	"reflect"
	"testing"
)

var ibPairs = []struct {
	i int
	b []byte
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
	for _, p := range ibPairs {
		got := intToBytes(p.i)
		if !reflect.DeepEqual(got, p.b) {
			t.Errorf("intToBytes(%d) = %q, want %q", p.i, got, p.b)
		}
	}
}

func TestBytesToInt(t *testing.T) {
	for _, p := range ibPairs {
		got := bytesToInt(p.b)
		if got != p.i {
			t.Errorf("bytesToInt(%q) = %d, want %d", p.b, got, p.i)
		}
	}
}

var dePairs = []struct {
	d, e []byte
}{
	{[]byte(""), []byte("")},
	{[]byte("a"), []byte("")},
	{[]byte("abcdef"), []byte("")},
	{[]byte("abcdefg"), []byte("wATBHB2AjAVAHBiB")},
	{[]byte("abcdefgh"), []byte("wATBHB2AjAVAHBiB")},
	{[]byte("abcdefg0123456"), []byte("wATBHB2AjAVAHBiBYAMAmAjAZALBlB2A")},
}

func TestEncodeHead(t *testing.T) {
	for _, p := range dePairs {
		got := encodeHead(p.d)
		if string(got) != string(p.e) {
			t.Errorf("encodeArrayWithLength(%q) = %q, want %q", p.d, got, p.e)
		}
	}
}
