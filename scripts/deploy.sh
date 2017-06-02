#!/bin/bash

set -e

# The deploy stage within Travis-CI dirtys the working tree. This will cause
# the docker tagging to fail. Reset to HEAD as a workaround.
git reset --hard HEAD
if [ -z "$1" ]; then
  make push
else
  make push RELEASE=$1
fi
