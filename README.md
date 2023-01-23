# Edge APIs

This repository contains the Open API 2.0 specification for the OpenZiti Edge Client and Management REST APIs. It also
contains a generated go module, in the `go` directory that can be used to develop against OpenZiti Controllers.

# Versioning

Versioning of the APIs in this repository are independent of the OpenZiti releases created in the 
[`ziti`](https://github.com/openziti/zit) repository. Many versions of these API specifications are  compatible with 
multiple versions of the OpenZiti release versions. To make it somewhat intuitive, the version number of the API
is the *minimum version of the `ziti`* repository releases that this API is compatible with. It will also be 
compatible up until the next version of the specifications.

For simplicity each controller hosts the specification version they expect and may be used instead of this repository
for live deployments.

# Client & Server Generation

The root level `client.yml` and `management.yml` files are generated from the `source` directory. There are scripts
within the `script` directory that will do the heavy lifting of re-generating them if needed.Both scripts require
that the `swagger` executable be available on your `path` environment variable. Releases of it are available in the
[GitHub Go-Swagger](https://github.com/go-swagger/go-swagger/releases) repository.

```bash
#bash
./scripts/generate_rest.sh
```

```powershell
#powershell
./scripts/generate_rest.ps1
```

# Using the generated go module

Within the go module within the `go` directory is a submodule named `rest_util` with contains
helper functions for using its sibling `*_client` submodules. This package is not generated.


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