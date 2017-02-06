.PHONY: build local get-deps fmt binary install clean

MAKE_SCRIPT="./build/make.sh"

default: all

all: build

build: getdeps fmt binary install

local: clean getdeps fmt binary install



fmt:
	${MAKE_SCRIPT} fmt

binary:
	${MAKE_SCRIPT} binary

getdeps:
	${MAKE_SCRIPT} getdeps

install:
	${MAKE_SCRIPT} install

clean:
	${MAKE_SCRIPT} clean
