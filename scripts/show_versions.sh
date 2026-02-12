#!/usr/bin/env bash

scriptDir=$(dirname "$(realpath "$0")")
rootDir=$(realpath "$scriptDir/..")

get_ver() {
  grep -m1 '^\s*version:' | awk '{print $2}'
}

clientCurrent=$(get_ver < "$rootDir/source/client.yml")
managementCurrent=$(get_ver < "$rootDir/source/management.yml")

clientLocalMain=$(git -C "$rootDir" show main:source/client.yml 2>/dev/null | get_ver || echo "(unavailable)")
managementLocalMain=$(git -C "$rootDir" show main:source/management.yml 2>/dev/null | get_ver || echo "(unavailable)")

clientRemoteMain=$(git -C "$rootDir" show origin/main:source/client.yml 2>/dev/null | get_ver || echo "(unavailable)")
managementRemoteMain=$(git -C "$rootDir" show origin/main:source/management.yml 2>/dev/null | get_ver || echo "(unavailable)")

printf "%-22s %-15s %s\n" ""                      "Client API"          "Management API"
printf "%-22s %-15s %s\n" "Remote (origin/main):" "$clientRemoteMain"   "$managementRemoteMain"
printf "%-22s %-15s %s\n" "Local main:"           "$clientLocalMain"    "$managementLocalMain"
printf "%-22s %-15s %s\n" "Current:"              "$clientCurrent"      "$managementCurrent"
