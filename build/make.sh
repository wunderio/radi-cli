#!/usr/bin/env bash
set -e

# We should be determining these automatically somehow?
GOOS="${GOOS:-linux}" # Perhaps you would prefer "osx" ?
GOARCH="${GOARCH:-x64}"
GOVERSION="latest"

export WUNDERTOOLS_PKG='github.com/wundertools/wundertools-go'
export WUNDERTOOLS_BUILD_PATH="./bin"
export WUNDERTOOLS_BINARY_NAME="kraut"
export WUNDERTOOLS_BUILD_BINARY_PATH="${WUNDERTOOLS_BUILD_PATH}/${WUNDERTOOLS_BINARY_NAME}"
export WUNDERTOOLS_INSTALL_PATH="${GOPATH}/bin/wundertools"

# Build a bundle
bundle() {
    local bundle="$1"; shift
    echo "---> Making bundle: $(basename "$bundle") (in $DEST)"
    source "build/$bundle" "$@"
}

if [ $# -gt 0 ]; then
    bundles=($@)
    for bundle in ${bundles[@]}; do
        export DEST=.
        ABS_DEST="$(cd "$DEST" && pwd -P)"
        bundle "$bundle"
        echo
    done
fi
