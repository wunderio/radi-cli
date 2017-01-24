#!/bin/bash

#
# Build wundertools in a container
#
# @NOTE to specify a different os/arch:
#    - GOOS : linux darwin windows
#    - GOARCH : amd64 arm arm64
#
# @NOTE !does not install it yet
#  (installs it, but inside the container)
#

source build/.os-detect
source build/make.sh

EXEC_PATH="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
INTERNAL_LIBRARY_PATH="github.com/wunderkraut/radi-cli"

echo "This will build the radi-cli as a 'radi' binary for '$GOOS-$GOARCH'. Override this by setting \$GOOS and \$GOARCH environment variables.

 **** Building in containerized golang environment
 "

docker run --rm -ti \
	-v "${EXEC_PATH}:/usr/src/myapp" \
	-v "${EXEC_PATH}:/go/src/${INTERNAL_LIBRARY_PATH}" \
	-e "GOOS=${GOOS}" \
	-e "GOARCH=${GOARCH}" \
	-w /usr/src/myapp \
	golang:${GOVERSION} \
	make getdeps build

echo " **** Containerized build complete

You should now have a binary ${KRAUT_BUILD_BINARY_PATH}
you can run this directly or install it somewhere usefull.

"
#ls -la ${KRAUT_BUILD_BINARY_PATH}
