#!/usr/bin/env bash
# Usage: bump_version.sh major|minor|fix

set -o errexit
set -o nounset
set -o pipefail

PART="${1:-}"
if [[ "$PART" != "major" && "$PART" != "minor" && "$PART" != "fix" ]]; then
  echo >&2 "Usage: $0 major|minor|fix"
  exit 1
fi

scriptDir=$(dirname "$(realpath "$0")")
rootDir=$(realpath "$scriptDir/..")

clientSpec="$rootDir/source/client.yml"
managementSpec="$rootDir/source/management.yml"

clientVersion=$(grep -m1 '^\s*version:' "$clientSpec" | awk '{print $2}')
managementVersion=$(grep -m1 '^\s*version:' "$managementSpec" | awk '{print $2}')

if [[ "$clientVersion" != "$managementVersion" ]]; then
  echo >&2 "API versions do not match before bump:"
  echo >&2 "  Client API:     $clientVersion"
  echo >&2 "  Management API: $managementVersion"
  exit 1
fi

IFS='.' read -r major minor patch <<< "$clientVersion"

case "$PART" in
  major) newVersion="$((major + 1)).0.0" ;;
  minor) newVersion="${major}.$((minor + 1)).0" ;;
  fix)   newVersion="${major}.${minor}.$((patch + 1))" ;;
esac

escapedVersion=$(printf '%s' "$clientVersion" | sed 's/\./\\./g')
sed -i "s/^  version: ${escapedVersion}$/  version: ${newVersion}/" "$clientSpec"
sed -i "s/^  version: ${escapedVersion}$/  version: ${newVersion}/" "$managementSpec"

newClientVersion=$(grep -m1 '^\s*version:' "$clientSpec" | awk '{print $2}')
newManagementVersion=$(grep -m1 '^\s*version:' "$managementSpec" | awk '{print $2}')

if [[ "$newClientVersion" != "$newVersion" || "$newManagementVersion" != "$newVersion" ]]; then
  echo >&2 "API versions do not match after bump:"
  echo >&2 "  Client API:     $clientVersion -> $newClientVersion (expected $newVersion)"
  echo >&2 "  Management API: $managementVersion -> $newManagementVersion (expected $newVersion)"
  exit 1
fi

echo "Version bumped: $clientVersion -> $newVersion"
