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

type Options struct {
	X *C.leveldb_options_t
}

func NewOptions() *Options {
	return &Options{C.leveldb_options_create()}
}

func NewDefaultOptions() *Options {
	o := &Options{C.leveldb_options_create()}
	o.SetErrorIfExists(false)
	o.SetCreateIfMissing(true)
	return o
}

func (o *Options) Destroy() {
	C.leveldb_options_destroy(o.X)
	o.X = nil
}

func (o *Options) SetComparator(x *Comparator) {
	C.leveldb_options_set_comparator(o.X, x.X)
}

func (o *Options) SetFilterPolicy(x *FilterPolicy) {
	C.leveldb_options_set_filter_policy(o.X, x.X)
}

func (o *Options) SetCreateIfMissing(x bool) {
	b := C.uchar(0)
	if x {
		b = C.uchar(1)
	}
	C.leveldb_options_set_create_if_missing(o.X, b)
}

func (o *Options) SetErrorIfExists(x bool) {
	b := C.uchar(0)
	if x {
		b = C.uchar(1)
	}
	C.leveldb_options_set_error_if_exists(o.X, b)
}

func (o *Options) SetParanoidChecks(x bool) {
	b := C.uchar(0)
	if x {
		b = C.uchar(1)
	}
	C.leveldb_options_set_paranoid_checks(o.X, b)
}

func (o *Options) SetEnv(x *Env) {
	C.leveldb_options_set_env(o.X, x.X)
}

func (o *Options) SetInfoLog(x *Logger) {
	C.leveldb_options_set_info_log(o.X, x.X)
}

func (o *Options) SetWriteBufferSize(x uint) {
	C.leveldb_options_set_write_buffer_size(o.X, C.size_t(x))
}

func (o *Options) SetMaxOpenFiles(x int) {
	C.leveldb_options_set_max_open_files(o.X, C.int(x))
}

func (o *Options) SetCache(x *Cache) {
	C.leveldb_options_set_cache(o.X, x.X)
}

func (o *Options) SetBlockSize(x uint) {
	C.leveldb_options_set_block_size(o.X, C.size_t(x))
}

func (o *Options) SetBlockRestartInterval(x int) {
	C.leveldb_options_set_block_restart_interval(o.X, C.int(x))
}

func (o *Options) SetMaxFileSize(x uint) {
	C.leveldb_options_set_max_file_size(o.X, C.size_t(x))
}

type Compression int

const (
	NoCompression Compression = iota
	SnappyCompression
)

func (o *Options) SetCompression(x Compression) {
	C.leveldb_options_set_compression(o.X, C.int(x))
}
