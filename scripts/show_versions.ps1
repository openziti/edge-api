$rootDir = Join-Path $PSScriptRoot "../" -Resolve
$sourceDir = Join-Path $rootDir "source" -Resolve

function Get-SpecVersion([string[]]$lines) {
    ($lines | Select-String -Pattern '^\s*version:\s*(.+)' | Select-Object -First 1).Matches.Groups[1].Value.Trim()
}

function Get-VersionFile([string[]]$lines) {
    if ($lines) { ($lines | Select-Object -First 1).Trim() } else { "" }
}

function Get-GitVersion([string]$ref, [string]$path) {
    $lines = git -C $rootDir show "${ref}:${path}" 2>$null
    if ($lines) { Get-SpecVersion $lines } else { "(unavailable)" }
}

function Get-GitVersionFile([string]$ref, [string]$path) {
    $lines = git -C $rootDir show "${ref}:${path}" 2>$null
    if ($lines) { Get-VersionFile $lines } else { "(unavailable)" }
}

$clientCurrent    = Get-SpecVersion (Get-Content (Join-Path $sourceDir "client.yml"))
$mgmtCurrent      = Get-SpecVersion (Get-Content (Join-Path $sourceDir "management.yml"))
$versionCurrent   = Get-VersionFile (Get-Content (Join-Path $rootDir "version"))

$clientLocalMain  = Get-GitVersion     "main"        "source/client.yml"
$mgmtLocalMain    = Get-GitVersion     "main"        "source/management.yml"
$versionLocalMain = Get-GitVersionFile "main"        "version"

$clientRemote     = Get-GitVersion     "origin/main" "source/client.yml"
$mgmtRemote       = Get-GitVersion     "origin/main" "source/management.yml"
$versionRemote    = Get-GitVersionFile "origin/main" "version"

"{0,-22} {1,-15} {2,-17} {3}" -f "",                      "Client API",    "Management API",  "version file"
"{0,-22} {1,-15} {2,-17} {3}" -f "Remote (origin/main):", $clientRemote,   $mgmtRemote,       $versionRemote
"{0,-22} {1,-15} {2,-17} {3}" -f "Local main:",           $clientLocalMain,$mgmtLocalMain,    $versionLocalMain
"{0,-22} {1,-15} {2,-17} {3}" -f "Current:",              $clientCurrent,  $mgmtCurrent,      $versionCurrent
