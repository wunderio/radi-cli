#!/usr/bin/env bash
set -e

export WUNDERTOOLS_PKG='github.com/wundertools/wundertools-go'

# List of bundles to create when no argument is passed
DEFAULT_BUNDLES=(
    binary
)
bundle() {
    local bundle="$1"; shift
    echo "---> Making bundle: $(basename "$bundle") (in $DEST)"
    source "scripts/$bundle" "$@"
}

if [ $# -lt 1 ]; then
    bundles=(${DEFAULT_BUNDLES[@]})
else
    bundles=($@)
fi
for bundle in ${bundles[@]}; do
    export DEST=.
    ABS_DEST="$(cd "$DEST" && pwd -P)"
    bundle "$bundle"
    echo
done
