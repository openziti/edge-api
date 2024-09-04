# !/usr/bin/env python3
import yaml
import semver
import subprocess

# Define the file paths
files = ['client.yml', 'management.yml', 'source/client.yml', 'source/management.yml']


# Function to extract version from YAML
def extract_version(file_path):
    with open(file_path, 'r') as f:
        content = yaml.safe_load(f)
        return content.get('info', {}).get('version')


# Extract versions from all files
versions = [extract_version(file) for file in files]

# Check if all versions are the same
if len(set(versions)) != 1:
    raise ValueError(f"Versions do not match across files: {versions}: please make sure all files {files} have the same version")

# Get the current version
current_version = versions[0]

# Check that current version is a valid semver
if not semver.VersionInfo.is_valid(current_version):
    raise ValueError(f"Current version '{current_version}' is not a valid semver.")

# Parse current version
current_version_info = semver.VersionInfo.parse(current_version)

# Get the latest git tag
result = subprocess.run(['git', 'describe', '--tags', '--abbrev=0', 'origin/main'], capture_output=True, text=True)
latest_tag = result.stdout.strip()
print(f"Latest tag: {latest_tag}")
latest_tag = latest_tag.lstrip('v')
print(f"Latest tag stripped of leading v: {latest_tag}")
# Parse the latest tag as a semver
if not semver.VersionInfo.is_valid(latest_tag):
    raise ValueError(f"Latest git tag '{latest_tag}' is not a valid semver.")

latest_version_info = semver.VersionInfo.parse(latest_tag)

# Calculate possible next versions
next_patch_version = str(latest_version_info.bump_patch())
next_minor_version = str(latest_version_info.bump_minor())
next_major_version = str(latest_version_info.bump_major())

# Check if the current version matches one of the expected next versions
if current_version not in [next_patch_version, next_minor_version, next_major_version]:
    raise ValueError(
        f"Current version '{current_version}' does not match the next expected versions: "
        f"patch '{next_patch_version}', minor '{next_minor_version}', or major '{next_major_version}'."
    )

print(f"All versions match and are valid. Current version '{current_version}' matches one of the next possible versions.")