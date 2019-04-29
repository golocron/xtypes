# xtypes

[![GoDoc](https://godoc.org/github.com/golocron/xtypes?status.svg)](https://godoc.org/github.com/golocron/xtypes) [![Go Report Card](https://goreportcard.com/badge/github.com/golocron/xtypes)](https://goreportcard.com/report/github.com/golocron/xtypes)

xtypes provides useful data structures for everyday use.

## Description

The package provides the following data structures:

- Priority Queue based on `container/heap` from standard library
- Queue based on slice
- Semaphore implemented with a channel
- Safe Map

These types are safe for concurrent use.


## Install

```bash
go get github.com/golocron/xtypes
```

## Examples

Examples can be found [here](examples/examples.go)

```bash
go run examples/examples.go

doing work in order by priority...
working in order: task with priority: 18
working in order: task with priority: 22
working in order: task with priority: 25
working in order: task with priority: 29
working in order: task with priority: 45
working in order: task with priority: 53
working in order: task with priority: 64
working in order: task with priority: 67
working in order: task with priority: 81
working in order: task with priority: 86
doing work in parallel...
working in parallel: task with priority: 18
working in parallel: task with priority: 53
working in parallel: task with priority: 29
working in parallel: task with priority: 81
working in parallel: task with priority: 45
working in parallel: task with priority: 64
working in parallel: task with priority: 25
working in parallel: task with priority: 22
working in parallel: task with priority: 86
working in parallel: task with priority: 67
```
