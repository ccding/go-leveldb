// Copyright 2020 Cong Ding
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package leveldb

/*
#cgo LDFLAGS: -lleveldb -lstdc++
#include <stdlib.h>
#include "leveldb/c.h"
*/
import "C"

import (
	"errors"
	"unsafe"
)

type Key []byte
type Value []byte

type Range struct {
	StartKey Key
	LimitKey Key
}

type LevelDB struct {
	X *C.leveldb_t
}

func Open(name string, opt *Options) (*LevelDB, error) {
	var cErr *C.char
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	if opt == nil {
		opt = NewDefaultOptions()
		defer opt.Destroy()
	}
	db := C.leveldb_open(opt.X, cName, &cErr)
	if cErr != nil {
		defer C.leveldb_free(unsafe.Pointer(cErr))
		return nil, errors.New(C.GoString(cErr))
	}
	return &LevelDB{db}, nil
}

func (db *LevelDB) Close() {
	C.leveldb_close(db.X)
}

func (db *LevelDB) Put(key Key, val Value, opt *WriteOptions) error {
	var cErr *C.char
	var cKey, cVal *C.char
	lenKey := len(key)
	lenVal := len(val)
	if lenKey > 0 {
		cKey = (*C.char)(unsafe.Pointer(&key[0]))
	}
	if lenVal > 0 {
		cVal = (*C.char)(unsafe.Pointer(&val[0]))
	}
	if opt == nil {
		opt = NewWriteOptions()
		defer opt.Destroy()
	}
	C.leveldb_put(db.X, opt.X, cKey, C.size_t(lenKey), cVal, C.size_t(lenVal), &cErr)
	if cErr != nil {
		defer C.leveldb_free(unsafe.Pointer(cErr))
		return errors.New(C.GoString(cErr))
	}
	return nil
}

func (db *LevelDB) Delete(key Key, opt *WriteOptions) error {
	var cErr *C.char
	var cKey *C.char
	lenKey := len(key)
	if lenKey > 0 {
		cKey = (*C.char)(unsafe.Pointer(&key[0]))
	}
	if opt == nil {
		opt = NewWriteOptions()
		defer opt.Destroy()
	}
	C.leveldb_delete(db.X, opt.X, cKey, C.size_t(lenKey), &cErr)
	if cErr != nil {
		defer C.leveldb_free(unsafe.Pointer(cErr))
		return errors.New(C.GoString(cErr))
	}
	return nil
}

func (db *LevelDB) Write(batch WriteBatch, opt *WriteOptions) error {
	var cErr *C.char
	if opt == nil {
		opt = NewWriteOptions()
		defer opt.Destroy()
	}
	C.leveldb_write(db.X, opt.X, batch.X, &cErr)
	if cErr != nil {
		defer C.leveldb_free(unsafe.Pointer(cErr))
		return errors.New(C.GoString(cErr))
	}
	return nil
}

func (db *LevelDB) Get(key Key, opt *ReadOptions) (Value, error) {
	var cErr *C.char
	var cKey *C.char
	var cLen C.size_t
	lenKey := len(key)
	if lenKey > 0 {
		cKey = (*C.char)(unsafe.Pointer(&key[0]))
	}
	if opt == nil {
		opt = NewReadOptions()
		defer opt.Destroy()
	}
	cVal := C.leveldb_get(db.X, opt.X, cKey, C.size_t(lenKey), &cLen, &cErr)
	if cErr != nil {
		defer C.leveldb_free(unsafe.Pointer(cErr))
		return nil, errors.New(C.GoString(cErr))
	}
	if cVal == nil {
		return nil, nil
	}
	defer C.leveldb_free(unsafe.Pointer(cVal))
	return C.GoBytes(unsafe.Pointer(cVal), C.int(cLen)), nil
}

func (db *LevelDB) CreateIterator(opt *ReadOptions) *Iterator {
	if opt == nil {
		opt = NewReadOptions()
		defer opt.Destroy()
	}
	return &Iterator{C.leveldb_create_iterator(db.X, opt.X)}
}

func (db *LevelDB) CreateSnapshot() *Snapshot {
	return &Snapshot{C.leveldb_create_snapshot(db.X)}
}

func (db *LevelDB) ReleaseSnapshot(snapshot *Snapshot) {
	C.leveldb_release_snapshot(db.X, snapshot.X)
}

func (db *LevelDB) PropertyValue(propName string) string {
	cName := C.CString(propName)
	defer C.free(unsafe.Pointer(cName))
	return C.GoString(C.leveldb_property_value(db.X, cName))
}

func (db *LevelDB) ApproximateSizes(rs []Range) []uint64 {
	lenRanges := len(rs)
	startKey := make([]*C.char, lenRanges)
	startKeyLen := make([]C.size_t, lenRanges)
	limitKey := make([]*C.char, lenRanges)
	limitKeyLen := make([]C.size_t, lenRanges)
	for i := 0; i < lenRanges; i++ {
		r := rs[i]
		startKey[i] = C.CString(string(r.StartKey))
		startKeyLen[i] = C.size_t(len(r.StartKey))
		limitKey[i] = C.CString(string(r.LimitKey))
		limitKeyLen[i] = C.size_t(len(r.LimitKey))
	}
	sizes := make([]uint64, lenRanges)
	C.leveldb_approximate_sizes(db.X, C.int(lenRanges), &startKey[0], &startKeyLen[0], &limitKey[0], &limitKeyLen[0], (*C.uint64_t)(&sizes[0]))
	for i := 0; i < lenRanges; i++ {
		C.free(unsafe.Pointer(startKey[i]))
		C.free(unsafe.Pointer(limitKey[i]))
	}
	return sizes
}

func (db *LevelDB) CompactRange(r Range) {
	startKeyLen := len(r.StartKey)
	limitKeyLen := len(r.LimitKey)
	var startKey, limitKey *C.char
	if startKeyLen > 0 {
		startKey = (*C.char)(unsafe.Pointer(&r.StartKey[0]))
	}
	if limitKeyLen > 0 {
		limitKey = (*C.char)(unsafe.Pointer(&r.LimitKey[0]))
	}
	C.leveldb_compact_range(db.X, startKey, C.size_t(startKeyLen), limitKey, C.size_t(limitKeyLen))
}

func DestroyDB(name string, opt *Options) error {
	var cErr *C.char
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	if opt == nil {
		opt = NewDefaultOptions()
		defer opt.Destroy()
	}
	C.leveldb_destroy_db(opt.X, cName, &cErr)
	if cErr != nil {
		defer C.leveldb_free(unsafe.Pointer(cErr))
		return errors.New(C.GoString(cErr))
	}
	return nil
}

func RepairDB(name string, opt *Options) error {
	var cErr *C.char
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	if opt == nil {
		opt = NewDefaultOptions()
		defer opt.Destroy()
	}
	C.leveldb_repair_db(opt.X, cName, &cErr)
	if cErr != nil {
		defer C.leveldb_free(unsafe.Pointer(cErr))
		return errors.New(C.GoString(cErr))
	}
	return nil
}
