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
#cgo LDFLAGS: -lleveldb
#include "leveldb/c.h"
*/
import "C"

import (
	"unsafe"
)

type WriteBatch struct {
	X *C.leveldb_writebatch_t
}

func NewWriteBatch() *WriteBatch {
	return &WriteBatch{C.leveldb_writebatch_create()}
}

func (wb *WriteBatch) Destroy() {
	C.leveldb_writebatch_destroy(wb.X)
	wb.X = nil
}

func (wb *WriteBatch) Clear() {
	C.leveldb_writebatch_clear(wb.X)
}

func (wb *WriteBatch) Put(key Key, val Value) {
	var cKey, cVal *C.char
	lenKey := len(key)
	lenVal := len(val)
	if lenKey > 0 {
		cKey = (*C.char)(unsafe.Pointer(&key[0]))
	}
	if lenVal > 0 {
		cVal = (*C.char)(unsafe.Pointer(&val[0]))
	}
	C.leveldb_writebatch_put(wb.X, cKey, C.size_t(lenKey), cVal, C.size_t(lenVal))
}

func (wb *WriteBatch) Delete(key Key) {
	var cKey *C.char
	lenKey := len(key)
	if lenKey > 0 {
		cKey = (*C.char)(unsafe.Pointer(&key[0]))
	}
	C.leveldb_writebatch_delete(wb.X, cKey, C.size_t(lenKey))
}

func (wb *WriteBatch) Append(src *WriteBatch) {
	C.leveldb_writebatch_append(wb.X, src.X)
}
