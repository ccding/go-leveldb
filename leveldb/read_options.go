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

type ReadOptions struct {
	X *C.leveldb_readoptions_t
}

func NewReadOptions() *ReadOptions {
	return &ReadOptions{C.leveldb_readoptions_create()}
}

func (o *ReadOptions) Destroy() {
	C.leveldb_readoptions_destroy(o.X)
	o.X = nil
}

func (o *ReadOptions) SetVerifyChecksums(x bool) {
	b := C.uchar(0)
	if x {
		b = C.uchar(1)
	}
	C.leveldb_readoptions_set_verify_checksums(o.X, b)
}

func (o *ReadOptions) SetFillCache(x bool) {
	b := C.uchar(0)
	if x {
		b = C.uchar(1)
	}
	C.leveldb_readoptions_set_fill_cache(o.X, b)
}

func (o *ReadOptions) SetSnapshot(x *Snapshot) {
	var s *C.leveldb_snapshot_t
	if x != nil {
		s = x.X
	}
	C.leveldb_readoptions_set_snapshot(o.X, s)
}
