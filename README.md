# Base69

[![Test](https://github.com/eiri/base69/workflows/test/badge.svg?branch=master)](https://github.com/eiri/base69/actions?query=workflow%3Atest)
[![Go Report Card](https://goreportcard.com/badge/github.com/eiri/base69)](https://goreportcard.com/report/github.com/eiri/base69)

Base69 is a binary-to-text encoding scheme. This is port of [pshihn/base69](https://github.com/pshihn/base69). Why Base69 when Base64 is adequate? Because it's _NICE!_

## Usage

```golang
package main

import (
    "fmt"

    "github.com/eiri/base69/encoding/base69"
)

func main() {
    msg := "Hello, 世界"
    encoded := base69.Encode([]byte(msg))
    fmt.Println(string(encoded))
    decoded := base69.Decode(encoded)
    fmt.Println(string(decoded))
}
```

Output:
```
kAZAtABBeB8ATBgAtBuASApB8ARBYA1=
Hello, 世界
```

## Testing

`go test -v ./...`

## Licence

[MIT](https://github.com/eiri/base69/blob/master/LICENSE)
