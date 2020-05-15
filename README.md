go-leveldb
=======

[![Build Status](https://travis-ci.org/ccding/go-leveldb.svg?branch=master)](https://travis-ci.org/ccding/go-leveldb)
[![License](https://img.shields.io/badge/License-Apache%202.0-red.svg)](https://opensource.org/licenses/Apache-2.0)
[![GoDoc](https://godoc.org/github.com/ccding/go-leveldb?status.svg)](http://godoc.org/github.com/ccding/go-leveldb/leveldb)

go-leveldb is a Go wrapper for [LevelDB](https://github.com/google/leveldb)

## Building

You can run `make` to build this library. It will download and compile
`leveldb` and `snappy`, then compile `go-leveldb`.

Alternatively, if you have installed `libleveldb-dev` on your machine, you can
use `go build` to compile `go-leveldb`. However, if `libleveldb-dev` is
installed in non-standard folders, you must set `CGO_CFLAGS` and `CGO_LDFLAGS`
such that `cgo` knows the location of `libleveldb-dev`.

## Usage

go-leveldb supports most functions of leveldb's C API, except those passing
function pointers to the API.

Example code:

```go
import "github.com/ccding/go-leveldb/leveldb"

// creates a database
db, err := leveldb.Open(dir, nil)
// put and get
err = db.Put([]byte("foo"), []byte("bar"), nil)
val, err := db.Get([]byte("foo"), nil) // returns "bar"
// closes the database
db.Close()
// deletes the database directory
err = leveldb.DestroyDB(dir, nil)
```


More details, please go to [`main.go`](https://github.com/ccding/go-leveldb/blob/master/main.go)
and [GoDoc](http://godoc.org/github.com/ccding/go-leveldb/leveldb)

## TODO

Write tests and comments
