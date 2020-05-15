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

type Env struct {
	X *C.leveldb_env_t
}

func NewDefaultEnv() *Env {
	return &Env{C.leveldb_create_default_env()}
}

func (e *Env) Destroy() {
	C.leveldb_env_destroy(e.X)
	e.X = nil
}

func (e *Env) GetTestDirectory() string {
	return C.GoString(C.leveldb_env_get_test_directory(e.X))
}
