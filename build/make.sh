#!/usr/bin/env bash
set -e

# We should be determining these automatically somehow?
GOOS="${GOOS:-linux}" # Perhaps you would prefer "osx" ?
GOARCH="${GOARCH:-x64}"
GOVERSION="latest"

export KRAUT_PKG='github.com/wunder/kraut-cli'
export KRAUT_BUILD_PATH="./bin"
export KRAUT_BINARY_NAME="kraut"

export KRAUT_BUILD_BINARY_PATH="${KRAUT_BUILD_PATH}/${KRAUT_BINARY_NAME}"
export KRAUT_INSTALL_PATH="${GOPATH}/bin/${KRAUT_BINARY_NAME}"

# Build a bundle
bundle() {
    local bundle="$1"; shift
    echo "---> Make-bundle: $(basename "$bundle") (in $DEST)"
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
