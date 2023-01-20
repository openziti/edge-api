#!/bin/bash -e

command -v swagger >/dev/null 2>&1 || {
  echo >&2 "Command 'swagger' not installed. See: https://github.com/go-swagger/go-swagger for installation"
  exit 1
}

scriptPath=$(realpath $0)
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

codeTarget="$rootDir/go"

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
pwd
echo $codeTarget
prev=$(pwd)
cd $codeTarget
go mod init github.com/openziti/edge-api/go || true
go get -u ./...
go mod tidy
cd $prev