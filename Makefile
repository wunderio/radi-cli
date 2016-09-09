.PHONY: cli-build build install clean

default: all

all: getdeps fmt cli-build install

fmt:
	./scripts/make.sh fmt	

cli-build:
	./scripts/make.sh binary

cli-getdeps:
	./scripts/make.sh getdeps

build: cli-build

getdeps: cli-getdeps

install: build
	./scripts/make.sh install

clean:
	./scripts/make.sh clean
