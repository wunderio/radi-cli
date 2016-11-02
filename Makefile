.PHONY: cli-build build install clean

MAKE_SCRIPT="./build/make.sh"

default: all

all: getdeps fmt cli-build install

fmt:
	${MAKE_SCRIPT} fmt

build:
	${MAKE_SCRIPT} binary

getdeps:
	${MAKE_SCRIPT} getdeps

install:
	${MAKE_SCRIPT} install

clean:
	${MAKE_SCRIPT} clean
