.PHONY: cli-build build install clean

default: cli-build

fmt:
	./scripts/make.sh fmt	

cli-build: fmt 
	./scripts/make.sh binary

build: cli-build

install: build
	./scripts/make.sh install

clean:
	./scripts/make.sh clean
