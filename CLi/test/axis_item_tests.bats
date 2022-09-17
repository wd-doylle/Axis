# !/usr/bin/env bats

setup() {
    rm $AXIS_HOME/axis.db || true
    load 'test_helper/bats-support/load'
    load 'test_helper/bats-assert/load'
}