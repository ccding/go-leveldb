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

type WriteOptions struct {
	X *C.leveldb_writeoptions_t
}

func NewWriteOptions() *WriteOptions {
	return &WriteOptions{C.leveldb_writeoptions_create()}
}

func (o *WriteOptions) Destroy() {
	C.leveldb_writeoptions_destroy(o.X)
	o.X = nil
}

func (o *WriteOptions) SetSync(x bool) {
	b := C.uchar(0)
	if x {
		b = C.uchar(1)
	}
	C.leveldb_writeoptions_set_sync(o.X, b)
}
