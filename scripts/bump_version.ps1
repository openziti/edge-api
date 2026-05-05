param(
    [Parameter(Mandatory)][ValidateSet("major", "minor", "fix")][string]$Part
)

$ErrorActionPreference = "Stop"

$rootDir = Join-Path $PSScriptRoot "../" -Resolve

$clientSpec = Join-Path $rootDir "source/client.yml" -Resolve
$managementSpec = Join-Path $rootDir "source/management.yml" -Resolve
$versionFile = Join-Path $rootDir "version" -Resolve

$versionPattern = '^\s*version:\s*(.+)'

$clientVersion = (Select-String -Path $clientSpec -Pattern $versionPattern | Select-Object -First 1).Matches.Groups[1].Value.Trim()
$managementVersion = (Select-String -Path $managementSpec -Pattern $versionPattern | Select-Object -First 1).Matches.Groups[1].Value.Trim()
$versionFileVersion = ((Get-Content $versionFile -TotalCount 1) -as [string]).Trim()

if ($clientVersion -ne $managementVersion) {
    throw "API versions do not match before bump:`n  Client API:     $clientVersion`n  Management API: $managementVersion"
}

$parts = $clientVersion -split '\.'
$major = [int]$parts[0]
$minor = [int]$parts[1]
$patch = [int]$parts[2]
$currentMajorMinor = "$major.$minor"

if ($versionFileVersion -ne $currentMajorMinor) {
    throw "Version file does not match API version before bump:`n  API version:  $clientVersion (major.minor: $currentMajorMinor)`n  version file: $versionFileVersion"
}

switch ($Part) {
    "major" { $newVersion = "$($major + 1).0.0";       $newMajorMinor = "$($major + 1).0" }
    "minor" { $newVersion = "$major.$($minor + 1).0";  $newMajorMinor = "$major.$($minor + 1)" }
    "fix"   { $newVersion = "$major.$minor.$($patch + 1)"; $newMajorMinor = $currentMajorMinor }
}

$escapedVersion = [regex]::Escape($clientVersion)
(Get-Content $clientSpec)     -replace "^  version: $escapedVersion$", "  version: $newVersion" | Set-Content $clientSpec -NoNewline:$false
(Get-Content $managementSpec) -replace "^  version: $escapedVersion$", "  version: $newVersion" | Set-Content $managementSpec -NoNewline:$false
Set-Content -Path $versionFile -Value $newMajorMinor -NoNewline:$false

$newClientVersion = (Select-String -Path $clientSpec -Pattern $versionPattern | Select-Object -First 1).Matches.Groups[1].Value.Trim()
$newManagementVersion = (Select-String -Path $managementSpec -Pattern $versionPattern | Select-Object -First 1).Matches.Groups[1].Value.Trim()
$newVersionFileVersion = ((Get-Content $versionFile -TotalCount 1) -as [string]).Trim()

if ($newClientVersion -ne $newVersion -or $newManagementVersion -ne $newVersion) {
    throw "API versions do not match after bump:`n  Client API:     $clientVersion -> $newClientVersion (expected $newVersion)`n  Management API: $managementVersion -> $newManagementVersion (expected $newVersion)"
}

if ($newVersionFileVersion -ne $newMajorMinor) {
    throw "Version file does not match API version after bump:`n  API version:  $newVersion (major.minor: $newMajorMinor)`n  version file: $versionFileVersion -> $newVersionFileVersion (expected $newMajorMinor)"
}

"Version bumped: $clientVersion -> $newVersion (version file: $versionFileVersion -> $newVersionFileVersion)"
