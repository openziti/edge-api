param(
    [Parameter(Mandatory)][ValidateSet("major", "minor", "fix")][string]$Part
)

$ErrorActionPreference = "Stop"

$rootDir = Join-Path $PSScriptRoot "../" -Resolve

$clientSpec = Join-Path $rootDir "source/client.yml" -Resolve
$managementSpec = Join-Path $rootDir "source/management.yml" -Resolve

$versionPattern = '^\s*version:\s*(.+)'

$clientVersion = (Select-String -Path $clientSpec -Pattern $versionPattern | Select-Object -First 1).Matches.Groups[1].Value.Trim()
$managementVersion = (Select-String -Path $managementSpec -Pattern $versionPattern | Select-Object -First 1).Matches.Groups[1].Value.Trim()

if ($clientVersion -ne $managementVersion) {
    throw "API versions do not match before bump:`n  Client API:     $clientVersion`n  Management API: $managementVersion"
}

$parts = $clientVersion -split '\.'
$major = [int]$parts[0]
$minor = [int]$parts[1]
$patch = [int]$parts[2]

$newVersion = switch ($Part) {
    "major" { "$($major + 1).0.0" }
    "minor" { "$major.$($minor + 1).0" }
    "fix"   { "$major.$minor.$($patch + 1)" }
}

$escapedVersion = [regex]::Escape($clientVersion)
(Get-Content $clientSpec)     -replace "^  version: $escapedVersion$", "  version: $newVersion" | Set-Content $clientSpec -NoNewline:$false
(Get-Content $managementSpec) -replace "^  version: $escapedVersion$", "  version: $newVersion" | Set-Content $managementSpec -NoNewline:$false

$newClientVersion = (Select-String -Path $clientSpec -Pattern $versionPattern | Select-Object -First 1).Matches.Groups[1].Value.Trim()
$newManagementVersion = (Select-String -Path $managementSpec -Pattern $versionPattern | Select-Object -First 1).Matches.Groups[1].Value.Trim()

if ($newClientVersion -ne $newVersion -or $newManagementVersion -ne $newVersion) {
    throw "API versions do not match after bump:`n  Client API:     $clientVersion -> $newClientVersion (expected $newVersion)`n  Management API: $managementVersion -> $newManagementVersion (expected $newVersion)"
}

"Version bumped: $clientVersion -> $newVersion"
