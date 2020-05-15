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

package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/ccding/go-leveldb/leveldb"
)

func main() {
	log.Printf("LevelDB Version: v%v.%v\n", leveldb.MajorVersion(), leveldb.MajorVersion())
	// creates a database
	dir, err := ioutil.TempDir(os.TempDir(), "go-leveldb-")
	if err != nil {
		log.Fatalf("Open DB error: %v", err)
	}
	db, err := leveldb.Open(dir, nil)
	if err != nil {
		log.Fatalf("Open DB error: %v", err)
	}
	log.Printf("Created and opened database: %v", dir)
	// put and get
	key := []byte("foo")
	val := []byte("bar")
	err = db.Put(key, val, nil)
	if err != nil {
		log.Printf("Put error: %v", err)
	}
	log.Printf("Put %v:%v", string(key), string(val))
	valGet, err := db.Get(key, nil)
	if err != nil {
		log.Printf("Get error: %v", err)
	}
	if !bytes.Equal(val, valGet) {
		log.Printf("Put value and Get value mismatch: %v vs %v", val, valGet)
	}
	log.Printf("Get %v:%v", string(key), string(valGet))
	// closes the database
	db.Close()
	log.Printf("Closed database")
	// deletes the test directory
	err = leveldb.DestroyDB(dir, nil)
	if err != nil {
		log.Fatalf("Destroy DB error: %v", err)
	}
	log.Printf("Deleted database: %v", dir)
}
