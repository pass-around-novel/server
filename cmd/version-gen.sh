#!/bin/bash
set -e
cd "$(dirname "$0")"

commit=$(git rev-parse HEAD)

cat <<EOF > version-gen.go
package cmd

// CommitID is the current commit in the Git repository
const CommitID = "$commit"
EOF
