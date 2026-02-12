$rootDir = Join-Path $PSScriptRoot "../" -Resolve
$sourceDir = Join-Path $rootDir "source" -Resolve

function Get-SpecVersion([string[]]$lines) {
    ($lines | Select-String -Pattern '^\s*version:\s*(.+)' | Select-Object -First 1).Matches.Groups[1].Value.Trim()
}

function Get-GitVersion([string]$ref, [string]$path) {
    $lines = git -C $rootDir show "${ref}:${path}" 2>$null
    if ($lines) { Get-SpecVersion $lines } else { "(unavailable)" }
}

$clientCurrent    = Get-SpecVersion (Get-Content (Join-Path $sourceDir "client.yml"))
$mgmtCurrent      = Get-SpecVersion (Get-Content (Join-Path $sourceDir "management.yml"))

$clientLocalMain  = Get-GitVersion "main"        "source/client.yml"
$mgmtLocalMain    = Get-GitVersion "main"        "source/management.yml"

$clientRemote     = Get-GitVersion "origin/main" "source/client.yml"
$mgmtRemote       = Get-GitVersion "origin/main" "source/management.yml"

"{0,-22} {1,-15} {2}" -f "",                      "Client API",    "Management API"
"{0,-22} {1,-15} {2}" -f "Remote (origin/main):", $clientRemote,   $mgmtRemote
"{0,-22} {1,-15} {2}" -f "Local main:",           $clientLocalMain,$mgmtLocalMain
"{0,-22} {1,-15} {2}" -f "Current:",              $clientCurrent,  $mgmtCurrent
