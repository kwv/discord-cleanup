# Contributing to **discord-cleanup**

Thanks for your interest in contributing to **discord-cleanup**!
This guide explains how to develop, build, test, and release new versions of the project.

------------------------------------------------------------------------

## 🛠️ Development Setup

### 1. Clone the Repository

``` sh
git clone https://github.com/kwv4/discord-cleanup.git
cd discord-cleanup
```

### 2. Fast Local Development

``` sh
make build-dev
```

This compiles the binary locally and builds a fresh Docker image for `amd64` in seconds.

### 3. Linting and Verification

``` sh
make lint
make verify-release
```

- `make lint`: Runs `golangci-lint`.
- `make verify-release`: Runs a "snapshot" version of the GoReleaser pipeline locally to verify your configuration without publishing anything.

------------------------------------------------------------------------

## 🚀 Releasing a New Version

The project uses **Git tags** to trigger automated Docker image builds and releases.

### Steps to Publish a New Version

1.  Ensure your changes are committed to `main`.
2.  Choose the version increment type:

``` sh
make bump        # Patch release (v0.0.1 -> v0.0.2)
make bump-minor  # Minor release (v0.0.1 -> v0.1.0)
make bump-major  # Major release (v0.0.1 -> v1.0.0)
```

This automatically updates the version, tags the commit, and pushes the tag to GitHub.

### 🤖 CI/CD Automation

The GitHub Actions workflow triggers on every push and pull request to ensure the code remains clean:

- **Push to main / PRs**: Runs `go test` and `golangci-lint`.
- **Tag Push**: Decodes the version from the tag, runs tests/lint, and then triggers **GoReleaser** to build and publish:
    - GitHub Release binaries (Linux, Windows, Darwin for `amd64` and `arm64`)
    - Multi-arch Docker images (`kwv4/discord-cleanup:v0.0.1` and `latest`)

------------------------------------------------------------------------

## 🧪 Testing Locally

To test the image before publishing:

``` sh
docker build -t discord-cleanup:test .
# Note: Requires valid environment variables
docker run --rm --env-file .env discord-cleanup:test
```
------------------------------------------------------------------------

## 🔄 Contribution Workflow

1. Create a branch for your changes (e.g., `feature/your-feature-name` or `bugfix/your-bug-name`).
2. Commit your changes to this branch.
3. Open a pull request (PR) targeting the `main` branch.
4. Ensure your PR includes a clear description of the changes.

------------------------------------------------------------------------

-   Builds the Docker image
-   Pushes it to Docker Hub as:
    -   `kwv4/discord-cleanup:<version>`
    -   `kwv4/discord-cleanup:latest`

------------------------------------------------------------------------

## 🧩 Required Secrets for CI

Add these secrets in your GitHub repository settings:

  Secret Name            Value
  ---------------------- ---------------------------
  `DOCKERHUB_USERNAME`   `kwv4`
  `DOCKERHUB_TOKEN`      *Docker Hub access token*

------------------------------------------------------------------------

## 💬 Questions?

Feel free to open an issue or start a discussion if you have questions, improvement ideas, or suggestions.
We welcome contributions of all kinds!
