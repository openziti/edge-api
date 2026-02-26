# Scripts

- **`generate_rest.sh` / `generate_rest.ps1`** — Flattens the source OpenAPI specs and regenerates all Go server, client, and model code. Run after making any changes to files under `source/`. Checks that `swagger` is installed at the expected version and that Go is available. If the API version has not been bumped relative to `main`, prompts to bump it before proceeding.

- **`bump_version.sh` / `bump_version.ps1`** — Bumps the version in both `source/client.yml` and `source/management.yml` together, ensuring they stay in sync. Takes a single argument: `major`, `minor`, or `fix`. Run before `generate_rest` when making API changes that require a version increment. Errors if the two source files are out of sync before or after the bump.

- **`show_versions.sh` / `show_versions.ps1`** — Displays the current Client API and Management API versions from three perspectives: the remote `origin/main`, the local `main` branch, and the current working files under `source/`. Useful for checking whether the version has been bumped relative to main before regenerating.
