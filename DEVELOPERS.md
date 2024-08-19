# How to Update the Specifications

## Go Swagger

Install the `swagger` v0.29.0 executable. Using this version minimizes pull request differences. Updating the required version in a separate PR is as simple as substituting the new version throughout the repo and re-generating the specs.

## Modify the Specification Sources

These files are used to build the specifications.

```bash
./source
├── client
├── client.yml
├── management
├── management.yml
└── shared
```

## Bump the Specification Version

Update the version numbers in source/client.yml and source/management.yml in your PR. The versions must be identical.

## Prepare the Local Filesystem

You must clone two repos in adjacent directories.

1. this repo - https://github.com/openziti/edge-api.git
1. the main ziti repo - https://github.com/openziti/ziti.git

## Generate the Specifications and Test

You must create a workspace for the main repo (`ziti`) to run tests against the checkout in this repo (`edge-api`).

```bash
(
    set -euxo pipefail;
    ./scripts/generate_rest.sh;
    (
        cd ../ziti;
        [[ -s ./go.work ]] && mv -v ./go.work{,.$(date --utc --iso-8601=seconds)};
        go work init;
        go work use .;
        go work use ../edge-api;
        go test ./... -tags apitests;
    )
)
```
