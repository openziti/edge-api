#!/usr/bin/env bash

scriptDir=$(dirname "$(realpath "$0")")
rootDir=$(realpath "$scriptDir/..")

get_ver() {
  grep -m1 '^\s*version:' | awk '{print $2}'
}

get_version_file() {
  head -n1 | tr -d '[:space:]'
}

clientCurrent=$(get_ver < "$rootDir/source/client.yml")
managementCurrent=$(get_ver < "$rootDir/source/management.yml")
versionFileCurrent=$(get_version_file < "$rootDir/version")

clientLocalMain=$(git -C "$rootDir" show main:source/client.yml 2>/dev/null | get_ver || echo "(unavailable)")
managementLocalMain=$(git -C "$rootDir" show main:source/management.yml 2>/dev/null | get_ver || echo "(unavailable)")
versionFileLocalMain=$(git -C "$rootDir" show main:version 2>/dev/null | get_version_file || echo "(unavailable)")

clientRemoteMain=$(git -C "$rootDir" show origin/main:source/client.yml 2>/dev/null | get_ver || echo "(unavailable)")
managementRemoteMain=$(git -C "$rootDir" show origin/main:source/management.yml 2>/dev/null | get_ver || echo "(unavailable)")
versionFileRemoteMain=$(git -C "$rootDir" show origin/main:version 2>/dev/null | get_version_file || echo "(unavailable)")

printf "%-22s %-15s %-17s %s\n" ""                      "Client API"          "Management API"          "version file"
printf "%-22s %-15s %-17s %s\n" "Remote (origin/main):" "$clientRemoteMain"   "$managementRemoteMain"   "$versionFileRemoteMain"
printf "%-22s %-15s %-17s %s\n" "Local main:"           "$clientLocalMain"    "$managementLocalMain"    "$versionFileLocalMain"
printf "%-22s %-15s %-17s %s\n" "Current:"              "$clientCurrent"      "$managementCurrent"      "$versionFileCurrent"
