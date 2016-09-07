.PHONY: binary

default: binary

all:

binary: 
	./scripts/make.sh binary

install: binary
	./scripts/make.sh install

clean:
	./scripts/make.sh clean
