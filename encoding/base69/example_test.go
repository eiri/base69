package base69_test

import (
	"fmt"

	"github.com/eiri/base69/encoding/base69"
)

func Example() {
	msg := "Hello, 世界"
	encoded := base69.Encode([]byte(msg))
	fmt.Println(string(encoded))
	decoded := base69.Decode(encoded)
	fmt.Println(string(decoded))
	// Output:
	// kAZAtABBeB8ATBgAtBuASApB8ARBYA1=
	// Hello, 世界
}
