#!/bin/bash

set -e

GOPACKAGES=$(go list ./... | grep -v /vendor/ | grep -v /api)
GOFILES=$(find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./api/*")

COVERAGE_REPORT=coverage.txt
PROFILE=profile.out

echo "Running tests"
if [[ -f ${COVERAGE_REPORT} ]]; then
  rm ${COVERAGE_REPORT}
fi
touch ${COVERAGE_REPORT}
for package in ${GOPACKAGES[@]}; do
  go test -v -race -coverprofile=${PROFILE} -covermode=atomic $package
  if [ -f ${PROFILE} ]; then
    cat ${PROFILE} >> ${COVERAGE_REPORT}
    rm ${PROFILE}
  fi
done

echo "Linting packages"
gometalinter.v1 --vendor --exclude=api/* --disable=gas --disable=gotype --sort=linter --deadline=240s ./...

echo "Formatting go files"
if [ ! -z "$(gofmt -l -s ${GOFILES})" ]; then
  exit 1
fi

exit 0
