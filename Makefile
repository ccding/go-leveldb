all: goleveldb

LEVELDB_VERSION ?= 1.22
SNAPPY_VERSION ?= 1.1.8

export LD_LIBRARY_PATH := $(PWD)/build/snappy-$(SNAPPY_VERSION)/build:$(PWD)/build/leveldb-$(LEVELDB_VERSION)/build:$(LD_LIBRARY_PATH)
export CFLAGS ?= -I$(PWD)/build/snappy-$(SNAPPY_VERSION)/include
export CXXFLAGS ?= -I$(PWD)/build/snappy-$(SNAPPY_VERSION)/build
export LDFLAGS ?= -L$(PWD)/build/snappy-$(SNAPPY_VERSION)/build
export CGO_CFLAGS ?= -I$(PWD)/build/snappy-$(SNAPPY_VERSION)/build -I$(PWD)/build/leveldb-$(LEVELDB_VERSION)/include
export CGO_LDFLAGS ?= -L$(PWD)/build/snappy-$(SNAPPY_VERSION)/build -L$(PWD)/build/leveldb-$(LEVELDB_VERSION)/build -lsnappy

goleveldb: *.go build/snappy-$(SNAPPY_VERSION)-STAMP build/leveldb-$(LEVELDB_VERSION)-STAMP
	go get -d .
	go build ./...
	go build .

archive:
	mkdir archive
archive/snappy-$(SNAPPY_VERSION).tar.gz: archive
	curl -L https://github.com/google/snappy/archive/$(SNAPPY_VERSION).tar.gz > $@
archive/leveldb-$(LEVELDB_VERSION).tar.gz: archive
	curl -L https://github.com/google/leveldb/archive/$(LEVELDB_VERSION).tar.gz > $@

build:
	mkdir build
build/leveldb-$(LEVELDB_VERSION): archive/leveldb-$(LEVELDB_VERSION).tar.gz build
	tar xzvf archive/leveldb-$(LEVELDB_VERSION).tar.gz -C build
build/leveldb-$(LEVELDB_VERSION)-STAMP: build/leveldb-$(LEVELDB_VERSION)
	mkdir build/leveldb-$(LEVELDB_VERSION)/build && cd build/leveldb-$(LEVELDB_VERSION)/build && cmake -DCMAKE_BUILD_TYPE=Release  .. && cmake --build .
	touch $@
build/snappy-$(SNAPPY_VERSION): archive/snappy-$(SNAPPY_VERSION).tar.gz build
	tar xzvf archive/snappy-$(SNAPPY_VERSION).tar.gz -C build
build/snappy-$(SNAPPY_VERSION)-STAMP: build/snappy-$(SNAPPY_VERSION)
	mkdir -p build/snappy-$(SNAPPY_VERSION)/build && cd build/snappy-$(SNAPPY_VERSION)/build && cmake .. && cmake --build .
	touch $@

test: goleveldb
	go test -v ./...
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...

clean:
	rm -rf build

.PHONY: all test
