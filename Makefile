.PHONY: all install

all: install

install:
	go install cmd/enc/enc.go
	go install cmd/dec/dec.go