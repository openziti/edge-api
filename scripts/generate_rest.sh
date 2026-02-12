#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

GO_SWAGGER_VERSION="v0.33.1"
GO_SWAGGER_HASH="2af7725271cf99ace5d44ab134acb53bffcc5734"
if ! command -v swagger &>/dev/null \
|| [[ "$(swagger version | awk '$1~/^version:/{print $2}')" != "${GO_SWAGGER_VERSION}" \
|| "$(swagger version | awk '$1~/^commit:/{print $2}')" != "${GO_SWAGGER_HASH}" ]]
then
  echo >&2 "Go Swagger executable 'swagger' ${GO_SWAGGER_VERSION} (${GO_SWAGGER_HASH}) is required. Download the binary from GitHub: https://github.com/go-swagger/go-swagger/releases/tag/v0.29.0"
  exit 1
fi

if GO_VERSION="$(go version | awk '{print $3}')"
then
  echo "Using go toolchain version: ${GO_VERSION}"
else
  echo >&2 "Failed to determine go toolchain version"
  exit 1
fi

scriptPath=$(realpath "$0")
scriptDir=$(dirname "$scriptPath")

rootDir=$(realpath "$scriptDir/..")

clientSourceSpec=$(realpath "$rootDir/source/client.yml")
clientSwagSpec=$(realpath -m "$rootDir/client.yml")

managementSourceSpec=$(realpath "$rootDir/source/management.yml")
managementSwagSpec=$(realpath -m "$rootDir/management.yml")

copyrightFile=$(realpath "$scriptDir/template.copyright.txt")

echo "...flattening client spec"
echo swagger flatten "$clientSourceSpec" -o "$clientSwagSpec" --format yaml
swagger flatten "$clientSourceSpec" -o "$clientSwagSpec" --format yaml
echo "...flattening management spec"
swagger flatten "$managementSourceSpec" -o "$managementSwagSpec" --format yaml

codeTarget="$rootDir"

clientServerPath=$(realpath -m "$codeTarget/rest_client_api_server")
echo "...removing any existing server from $clientServerPath"
rm -rf "$clientServerPath"
mkdir -p "$clientServerPath"

clientClientPath=$(realpath -m "$codeTarget/rest_client_api_client")
echo "...removing any existing client from $clientClientPath"
rm -rf "$clientClientPath"
mkdir -p "$clientClientPath"

managementServerPath=$(realpath -m "$codeTarget/rest_management_api_server")
echo "...removing any existing server from $managementServerPath"
rm -rf "$managementServerPath"
mkdir -p "$managementServerPath"

managementClientPath=$(realpath -m "$codeTarget/rest_management_api_client")
echo "...removing any existing client from $managementClientPath"
rm -rf "$managementClientPath"
mkdir -p "$managementClientPath"

modelPath=$(realpath -m "$codeTarget/rest_model")
echo "...removing any existing model from $modelPath"
rm -rf "$modelPath"
mkdir -p "$modelPath"

echo "...generating client api server"
swagger generate server --exclude-main --additional-initialism=jwt -f "$clientSwagSpec" -s rest_client_api_server -t "$codeTarget" -r "$copyrightFile" -m "rest_model"
exit_status=$?
if [ ${exit_status} -ne 0 ]; then
  echo "Failed to generate client api server. See above."
  exit "${exit_status}"
fi

echo "...generating client api client"
swagger generate client --additional-initialism=jwt -f "$clientSwagSpec" -c rest_client_api_client -t "$codeTarget" -r "$copyrightFile" -m "rest_model"
exit_status=$?
if [ ${exit_status} -ne 0 ]; then
  echo "Failed to generate client api client. See above."
  exit "${exit_status}"
fi

echo "...generating management api server"
swagger generate server --exclude-main --additional-initialism=jwt -f "$managementSwagSpec" -s rest_management_api_server -t "$codeTarget" -r "$copyrightFile" -m "rest_model"
exit_status=$?
if [ ${exit_status} -ne 0 ]; then
  echo "Failed to generate management api server. See above."
  exit "${exit_status}"
fi

echo "...generating management api management"
swagger generate client --additional-initialism=jwt -f "$managementSwagSpec" -c rest_management_api_client -t "$codeTarget" -r "$copyrightFile" -m "rest_model"
exit_status=$?
if [ ${exit_status} -ne 0 ]; then
  echo "Failed to generate management api client. See above."
  exit "${exit_status}"
fi

echo "...fixing go module deps"
prev=$(pwd)
cd "$codeTarget"
go mod init github.com/openziti/edge-api || true
go mod tidy
cd "$prev"
