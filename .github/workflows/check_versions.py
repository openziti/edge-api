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

# Check that the version file matches the spec major.minor
with open('version', 'r') as f:
    version_file_content = f.read().strip()

expected_version_file = f"{current_version_info.major}.{current_version_info.minor}"
if version_file_content != expected_version_file:
    raise ValueError(
        f"Version file contains '{version_file_content}' but spec version is '{current_version}'. "
        f"Expected version file to contain '{expected_version_file}'."
    )

# Get the latest git tag by semver ordering (not by reachability, which breaks on shallow clones)
result = subprocess.run(['git', 'tag', '--sort=-v:refname'], capture_output=True, text=True)
all_tags = [t for t in result.stdout.strip().split('\n') if t and semver.VersionInfo.is_valid(t.lstrip('v'))]
if not all_tags:
    raise ValueError("No valid semver tags found in the repository.")
latest_tag = all_tags[0]
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