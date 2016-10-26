#!/bin/bash

#
# Build wundertools in a container
#
# @NOTE to specify a different os/arch:
#    - GOOS : linux osx windows
#    - GOARCH : 386 x64 rpi arm
#
# @NOTE !does not install it yet
#  (installs it, but inside the container)
#

set -e

source build/make.sh

EXEC_PATH="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
INTERNAL_LIBRARY_PATH="github.com/james-nesbitt/wundertools-go"

echo "This will build wundertools as a 'kraut' for '$GOOS-$GOARCH' (you can change that)

 **** Building Wundertools in containerized golang environment
 "

docker run --rm -ti -v "${EXEC_PATH}:/usr/src/myapp" -v "${EXEC_PATH}:/go/src/${INTERNAL_LIBRARY_PATH}" -e GOOS -e GOARCH -w /usr/src/myapp golang:${GOVERSION} make getdeps build

echo " **** Containerized build complete

You should now have a binary ${WUNDERTOOLS_BUILD_BINARY_PATH}
you can run this directly or install it somewhere usefull.

"
#ls -la ${WUNDERTOOLS_BUILD_BINARY_PATH}
