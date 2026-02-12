# Edge APIs

This repository contains the Open API 2.0 specification for the OpenZiti Edge Client and Management REST APIs. It also
contains a generated go module, in the `rest_*` directories that can be used to develop against OpenZiti Controllers.

## Versioning

Versioning of the APIs in this repository is independent of the OpenZiti releases created in the
[`ziti`](https://github.com/openziti/zit) repository. Many versions of these API specifications are compatible with
multiple versions of the OpenZiti release versions. To make it somewhat intuitive, the minor version number of the API
is the *minimum minor version of the `ziti`* repository releases that this API is compatible with. It will also be
compatible up until the next minor version of the specifications. Patch versions are used for internal fixes and release
tags.

For simplicity, each controller hosts the specification version they expect and may be used instead of this repository
for live deployments.

## Client & Server Generation

The root-level `client.yml` and `management.yml` OpenAPI 2.0 specification  files are generated from the `source`
directory. There are scripts in `./scripts/` for re-generating these specifications and the Go client and server
libraries. The scripts require that a specific build of the `swagger` (aka "go-swagger") executable.
[Download the Go-Swagger v0.33.1 binary](https://github.com/go-swagger/go-swagger/releases/tag/v0.33.1) from GitHub.

Do not `go install` the binary because the rendering may introduce whitespace changes in many files.

### Modify the Specification Sources

These files are used to build the specifications.

```bash
./source
├── client
├── client.yml
├── management
├── management.yml
└── shared
```

### Bump the Specification Version

Update the version numbers in `./source/client.yml` and `./source/management.yml` in your PR. The versions must be identical.

### Prepare the Local Filesystem

You must clone two repos in adjacent directories.

1. this repo - https://github.com/openziti/edge-api.git
1. the main ziti repo - https://github.com/openziti/ziti.git

### Generate the Specifications and Test

You must create a workspace for the main repo (`ziti`) to run tests against the checkout in this repo (`edge-api`).

#### Linux/Darwin Example

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

#### Windows Example

```powershell
#powershell
./scripts/generate_rest.ps1
```

## Using the generated go module

Within the go module within the `go` directory is a submodule named `rest_util` containing helper functions for using
its sibling `*_client` submodules. This package is not generated. See `rest_util/examples` for full examples.

Example:

```go
func main() {
	ctrlAddress := "https://localhost:1280"
	caCerts, err := rest_util.GetControllerWellKnownCas(ctrlAddress)

	if err != nil {
		log.Fatal(err)
	}

	caPool := x509.NewCertPool()

	for _, ca := range caCerts {
		caPool.AddCert(ca)
	}

	ok, err := rest_util.VerifyController(ctrlAddress, caPool)

	if err != nil {
		log.Fatal(err)
	}

	if !ok {
		log.Fatal("controller failed CA validation")
	}

	client, err := rest_util.NewEdgeManagementClientWithUpdb("admin", "admin", ctrlAddress, caPool)

	if err != nil {
		log.Fatal(err)
	}

	params := &identity.ListIdentitiesParams{
		Context: context.Background(),
	}

	resp, err := client.Identity.ListIdentities(params, nil)

	if err != nil {
		log.Fatal(err)
	}

	println("\n=== Identity List ===")
	for _, identityItem := range resp.GetPayload().Data {
		println(*identityItem.Name)
	}
}
```
