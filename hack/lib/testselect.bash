#!/usr/bin/env bash

if [[ -n "${ARTIFACT_DIR:-}" ]]; then
  GO111MODULE=off go get github.com/openshift-knative/hack/cmd/testselect

  # CLONEREFS_OPTIONS var is set in CI
  echo "${CLONEREFS_OPTIONS}" > "${ARTIFACT_DIR}/clonerefs.json"

  cat "${ARTIFACT_DIR}/clonerefs.json"

  rootdir="$(dirname "$(dirname "$(dirname "$(realpath "${BASH_SOURCE[0]}")")")")"

  testselect --testsuites="${rootdir}/test/testsuites.yaml" --clonerefs="${ARTIFACT_DIR}/clonerefs.json" --output="${ARTIFACT_DIR}/tests.txt"

  logger.info 'Tests to be run:'
  cat "${ARTIFACT_DIR}/tests.txt"
fi
