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

type Iterator struct {
	X *C.leveldb_iterator_t
}

func (it *Iterator) Destroy() {
	C.leveldb_iter_destroy(it.X)
	it.X = nil
}

func (it *Iterator) Valid() bool {
	return !(C.leveldb_iter_valid(it.X) == C.uchar(0))
}

func (it *Iterator) SeekToFirst() {
	C.leveldb_iter_seek_to_first(it.X)
}

func (it *Iterator) SeekToLast() {
	C.leveldb_iter_seek_to_last(it.X)
}

func (it *Iterator) Seek(key Key) {
	keyLen := len(key)
	cKey := (*C.char)(unsafe.Pointer(&key[0]))
	C.leveldb_iter_seek(it.X, cKey, C.size_t(keyLen))
}

func (it *Iterator) Next() {
	C.leveldb_iter_next(it.X)
}

func (it *Iterator) Prev() {
	C.leveldb_iter_prev(it.X)
}

func (it *Iterator) Key() Key {
	var cLen C.size_t
	cKey := C.leveldb_iter_key(it.X, &cLen)
	if cKey == nil {
		return nil
	}
	return C.GoBytes(unsafe.Pointer(cKey), C.int(cLen))
}

func (it *Iterator) Value() Value {
	var cLen C.size_t
	cVal := C.leveldb_iter_value(it.X, &cLen)
	if cVal == nil {
		return nil
	}
	return C.GoBytes(unsafe.Pointer(cVal), C.int(cLen))
}

func (it *Iterator) GetError() error {
	var cErr *C.char
	C.leveldb_iter_get_error(it.X, &cErr)
	if cErr != nil {
		defer C.leveldb_free(unsafe.Pointer(cErr))
		return errors.New(C.GoString(cErr))
	}
	return nil
}
