#!/bin/sh

set -e

if [ ! -f "bin/build.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/bin/_workspace"
root="$PWD"
epvdir="$workspace/src/github.com/epvchain"
if [ ! -L "$epvdir/go-epvchain" ]; then
    mkdir -p "$epvdir"
    cd "$epvdir"
    ln -s ../../../../../. go-epvchain
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$epvdir/go-epvchain"
PWD="$epvdir/go-epvchain"

# Launch the arguments with the configured environment.
exec "$@"
