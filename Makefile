.PHONY: build local get-deps fmt binary install clean

MAKE_SCRIPT="./build/make.sh"

default: all

all: build

build: getdeps binary install

local: fmt binary install
local-full: getdeps local



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
