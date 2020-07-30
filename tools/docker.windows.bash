#!/usr/bin/env bash

docker run --rm -it -v "//${PWD}":/workspace --workdir=//workspace --entrypoint="" l.gcr.io/google/bazel:latest bash