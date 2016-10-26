.PHONY: cli-build build install clean

MAKE_SCRIPT="./build/make.sh"

default: all

all: getdeps fmt cli-build install

fmt:
	${MAKE_SCRIPT} fmt

cli-build:
	${MAKE_SCRIPT} binary

cli-getdeps:
	${MAKE_SCRIPT} getdeps

build: cli-build

getdeps: cli-getdeps

install: build
	${MAKE_SCRIPT} install

clean:
	${MAKE_SCRIPT} clean
