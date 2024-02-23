# ABool :bulb:
[![Go Report Card](https://goreportcard.com/badge/github.com/tevino/abool)](https://goreportcard.com/report/github.com/tevino/abool)
[![GoDoc](https://godoc.org/github.com/tevino/abool?status.svg)](https://godoc.org/github.com/tevino/abool)

Atomic Boolean library for Go, optimized for performance yet simple to use.

Use this for cleaner code.

Froked from [github.com/tevino/abool](github.com/tevino/abool).

## Usage

```go
import "github.com/barryz/abool"

cond := abool.New()  // default to false

cond.Set()                 // Set to true
cond.IsSet()               // Returns true
cond.UnSet()               // Set to false
cond.SetTo(true)           // Set to whatever you want
cond.SetToIf(false, true)  // Set to true if it is false, returns false(not set)
cond.Toggle() *AtomicBool  // Negates boolean atomically and returns a new AtomicBool object which holds previous boolean value.


// embedding
type Foo struct {
    cond *abool.AtomicBool  // always use pointer to avoid copy
}
```

## Benchmark:

- Go 1.14.2
- OS X 10.15.5

```shell
# Read
BenchmarkMutexRead-12                   100000000               11.0 ns/op
BenchmarkAtomicValueRead-12             1000000000               0.253 ns/op
BenchmarkAtomicBoolRead-12              1000000000               0.259 ns/op    # <--- This package

# Write
BenchmarkMutexWrite-12                  100000000               10.8 ns/op
BenchmarkAtomicValueWrite-12            132855918                9.12 ns/op
BenchmarkAtomicBoolWrite-12             263941647                4.52 ns/op     # <--- This package

# CAS
BenchmarkMutexCAS-12                    54871387                21.6 ns/op
BenchmarkAtomicBoolCAS-12               267147930                4.50 ns/op     # <--- This package

# Toggle
BenchmarkMutexToggle-12                 55389297                21.4 ns/op
BenchmarkAtomicBoolToggle-12            145680895                8.32 ns/op     # <--- This package

```
