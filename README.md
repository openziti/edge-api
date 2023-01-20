# Overview

This folder contains the specifications for the two current edge APIs:

    - `client.yml` - OpenZiti Edge Client REST API
    - `management.yml` - OpenZiti Edge Management REST API

These two specs are built from the sharded files that are maintained in the `source` directory. To consume the
specs as is, single reference `client.yml` or `management.yml` when invoking your tool of choice.

# Updating

To modify or enhance the specifications in any manner, the changes must be made within the files in the `source`
directory. If not, during generation, any modifications to the top level `client.yml` and `management.yml` files
will be lost.

After modifying files within the `source` directory, use either the powershell or bash scripts located in the `scripts`
directory. These scripts will invoke `go-swagger`, so please ensure it is available on your path.

* Download/install `swagger` from https://github.com/go-swagger/go-swagger/releases
* cd to the directory containing this README
* issue either:
  ** `scripts/generate_rest.sh`
  ** `powershell -file "scripts/generate_rest.ps1"`
