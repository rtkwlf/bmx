#!/usr/bin/env bash

echo STABLE_GIT_COMMIT $(git rev-parse HEAD)
echo BMX_VERSION $(git tag --points-at HEAD | head -n 1)