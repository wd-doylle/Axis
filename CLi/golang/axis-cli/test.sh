# !/usr/bin/sh

env AXIS_BINARY_PATH=$(pwd)/axis-cli $(pwd)/../../CLi/test/bats/bin/bats $(pwd)/../../Axis/CLi/test/*.bats