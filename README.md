# Kesh 💶💶

[![Build Status](https://github.com/egregors/kesh/actions/workflows/ci.yml/badge.svg)](https://github.com/egregors/kesh/actions) [![Coverage Status](https://coveralls.io/repos/github/egregors/kesh/badge.svg?branch=main)](https://coveralls.io/github/egregors/kesh?branch=main)

Some cache policies implementation in Go

## Install

`go get github.com/egregors/kesh`

## Usage

All cache policies implementation should satisfy interface:
```go
// Cache is the common interface for all cache data structure implementations
type Cache interface {
    Get(key interface{}) (interface{}, error)
    Put(key interface{}, val interface{})
}
```

## Policies
### Least recently used (LRU)

https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_recently_used_(LRU)

#### Usage

```go
package main

import "github.com/egregors/kesh"

func main() {
	// init lru cache with capacity 2
	lru := kesh.NewLRUCache(2)

	lru.Put(42, "answer")
	lru.Put("super key", "mega value")

	r1, err := lru.Get(42)
	...
}

```

#### Benchmarks
```shell
BenchmarkLRUCache_Get_ifNotExist-12              8688457               129.8 ns/op
BenchmarkLRUCache_Get_ifExist-12                63910112                18.03 ns/op
BenchmarkLRUCache_Put_inCapacity-12             38644390                30.61 ns/op
BenchmarkLRUCache_Put_outOfCapacity-12           8893297               126.0 ns/op
```

