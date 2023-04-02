#!/bin/bash
set -x

generate_protobuf_for() {
    local FILE="$1"
	local FILE_DIR="$(dirname $FILE)"

    echo "Generating protobuf file from $FILE in $FILE_DIR"
    protoc -I $REPO_ROOT/client/api $FILE --go_out=plugins="grpc:$REPO_ROOT"
}

export -f generate_protobuf_for

echo "$REPO_ROOT is repo root"

# using xargs instead of -exec since the latter doesn't propagate exit statuses
find $REPO_ROOT -not -path "$REPO_ROOT/vendor/*" -name '*.proto' | xargs -n 1 protoc -I $REPO_ROOT --go_out=plugins=grpc:$REPO_ROOT --go_opt=paths=source_relative
