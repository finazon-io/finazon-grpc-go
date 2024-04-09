#!/bin/bash
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
SERVICES_GENERATOR=${SCRIPT_DIR}/tools/services-generator

### Implementation ##################################################
function main {
  local -r version="${1}"
  local -r versionMajor=$(echo ${version} | cut -d. -f1)
  local -r versionMinor=$(echo ${version} | cut -d. -f2)

  if [ -z "$version" ]
  then
    echo "VERSION is not set"
    exit 1
  fi

  printf "package finazon\n\n// ${version}\nvar FINAZON_GRPC_HOST = \"grpc-v${versionMajor}-${versionMinor}.finazon.io:443\"\n" > ${SCRIPT_DIR}/finazon/constants.go

  # generate service wrappers
  npm ci --prefix ${SERVICES_GENERATOR}
  node ${SERVICES_GENERATOR}/src/index.js ${SERVICES_GENERATOR}/templates ${SCRIPT_DIR}/proto ${SCRIPT_DIR}/finazon
}

# Script's entry point: #############################################
main "${@}"
