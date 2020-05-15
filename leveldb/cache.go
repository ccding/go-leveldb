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

type Cache struct {
	X *C.leveldb_cache_t
}

func NewLRUCache(capacity uint) *Cache {
	return &Cache{C.leveldb_cache_create_lru(C.size_t(capacity))}
}

func (c *Cache) Destroy() {
	C.leveldb_cache_destroy(c.X)
	c.X = nil
}
