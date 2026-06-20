<p align="center"><img src="https://raw.githubusercontent.com/go-composites/brand/main/social/go-composites.png" alt="go-composites/null" width="720"></p>

# null

[![ci](https://github.com/go-composites/null/actions/workflows/ci.yml/badge.svg)](https://github.com/go-composites/null/actions/workflows/ci.yml)

The Null-Object primitive of [go-composites](https://github.com/go-composites).
`Null` is the inert, never-nil value that the other composites reach for when
they need an *absence* that still honours an interface rather than a bare Go
`nil`. It is the default payload of a fresh
[`Result`](https://github.com/go-composites/result), so "no value yet" is itself
a value.

## Install

```sh
go get github.com/go-composites/null
```

## API

| symbol | kind | notes |
| --- | --- | --- |
| `Interface` | type | alias for `interface{}` — a `Null` satisfies every interface |
| `New()` | `Interface` | returns a fresh, inert Null-Object |
| `IsNul()` | `bool` (value method) | always `true` on a `Null` |

## Usage

```go
package main

import (
	"fmt"

	Null "github.com/go-composites/null/src"
	Result "github.com/go-composites/result/src"
)

func main() {
	n := Null.New()
	fmt.Println(n) // an inert Null-Object, not nil

	// A fresh Result carries a Null payload by default.
	r := Result.New()
	fmt.Println(r.HasError())  // false
	fmt.Println(r.Payload())   // the Null-Object
}
```

## License

BSD-3-Clause © the go-composites/null authors.
