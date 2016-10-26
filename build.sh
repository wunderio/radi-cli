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

# Get OS and architecture so that we can build correct binary for the host.
OS="$(uname | tr 'A-Z' 'a-z')"
ARC="$(uname -m)"

# Convert $ARC to accepted $GOARCH values by hand. Only know cases work.
# See list of options here https://golang.org/doc/install/source.
if [ $ARC == "x86_64" ]
then
  ARC="amd64"
elif [ "$ARC" == "ARMv7" ] || [ "$ARC" == "ARMv6" ]
then
  ARC="arm"
elif [ "$ARC" == "ARMv8" ]
then
  ARC="arm64"
fi

# Only set $GOOS and $GOARCH if they are not set to allow overriding the build target.
if [ -z "$GOOS" ]
then
  GOOS=${OS}
fi
if [ -z "$GOARCH" ]
then
  GOARCH=${ARC}
fi

source build/make.sh

EXEC_PATH="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
INTERNAL_LIBRARY_PATH="github.com/james-nesbitt/wundertools-go"

echo "This will build wundertools as a 'kraut' for '$GOOS-$GOARCH'. Change this by setting \$GOOS and \$GOARCH environment variables.

 **** Building Wundertools in containerized golang environment
 "

docker run --rm -ti -v "${EXEC_PATH}:/usr/src/myapp" -v "${EXEC_PATH}:/go/src/${INTERNAL_LIBRARY_PATH}" -e "GOOS=${GOOS}" -e "GOARCH=${GOARCH}" -w /usr/src/myapp golang:${GOVERSION} make getdeps build

echo " **** Containerized build complete

You should now have a binary ${WUNDERTOOLS_BUILD_BINARY_PATH}
you can run this directly or install it somewhere usefull.

"
#ls -la ${WUNDERTOOLS_BUILD_BINARY_PATH}
