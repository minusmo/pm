package manual

const apiReferenceTmpl = `---
title: API Reference
description: Public API surface and usage documentation
tags: api, reference, documentation
---

# API Reference

## Package Overview

<!-- TODO: Describe the public API surface -->

## Core Types

### Type: Config

` + "```" + `
// TODO: Add type definition
` + "```" + `

### Type: Client

` + "```" + `
// TODO: Add type definition
` + "```" + `

## Functions

### NewClient

` + "```" + `
// TODO: Add function signature and description
` + "```" + `

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
|      |      |             |

**Returns:**

| Type | Description |
|------|-------------|
|      |             |

## Constants & Defaults

| Constant | Value | Description |
|----------|-------|-------------|
|          |       |             |
`

const usageExamplesTmpl = `---
title: Usage Examples
description: Practical examples and common use cases
tags: examples, usage, quickstart
---

# Usage Examples

## Quick Start

` + "```" + `
// TODO: Add minimal working example
` + "```" + `

## Common Use Cases

### Basic Usage

` + "```" + `
// TODO: Add basic usage example
` + "```" + `

### Advanced Configuration

` + "```" + `
// TODO: Add advanced configuration example
` + "```" + `

### Error Handling

` + "```" + `
// TODO: Add error handling example
` + "```" + `

## Integration Examples

### With HTTP Server

` + "```" + `
// TODO: Add HTTP integration example
` + "```" + `

### With CLI Application

` + "```" + `
// TODO: Add CLI integration example
` + "```" + `
`

const versioningTmpl = `---
title: Versioning
description: Version strategy, changelog, and compatibility policy
tags: versioning, semver, changelog
---

# Versioning

## Version Strategy

This project follows [Semantic Versioning](https://semver.org/):

| Change Type | Version Bump | Example |
|-------------|-------------|---------|
| Breaking API change | Major (X.0.0) | Removing a public function |
| New feature (backwards-compatible) | Minor (0.X.0) | Adding a new method |
| Bug fix | Patch (0.0.X) | Fixing incorrect behavior |

## Compatibility Policy

<!-- TODO: Describe supported version range and deprecation policy -->

## Changelog

### Unreleased

- <!-- TODO: Add unreleased changes -->

## Deprecation Process

1. Mark as deprecated in code and docs
2. Log deprecation warning at runtime
3. Maintain for at least one minor release
4. Remove in next major release
`

const publishingTmpl = `---
title: Publishing
description: Release and publishing procedures
tags: publishing, release, distribution
---

# Publishing

## Pre-publish Checklist

- [ ] All tests passing
- [ ] Changelog updated
- [ ] Version bumped
- [ ] Documentation updated
- [ ] Breaking changes documented

## Publishing Steps

### 1. Update Version

` + "```bash" + `
# TODO: Add version bump commands
` + "```" + `

### 2. Build & Verify

` + "```bash" + `
# TODO: Add build and verification commands
` + "```" + `

### 3. Publish

` + "```bash" + `
# TODO: Add publish commands
` + "```" + `

## Distribution Channels

| Channel | URL | Notes |
|---------|-----|-------|
|         |     |       |

## Post-publish Verification

` + "```bash" + `
# TODO: Add verification commands (e.g., install from registry)
` + "```" + `
`

const contributingTmpl = `---
title: Contributing
description: Guidelines for contributing to this project
tags: contributing, guidelines, community
---

# Contributing

## Getting Started

1. Fork the repository
2. Clone your fork
3. Create a feature branch
4. Make your changes
5. Submit a pull request

## Development Setup

` + "```bash" + `
# TODO: Add development setup commands
` + "```" + `

## Pull Request Guidelines

- <!-- TODO: Add PR requirements (tests, docs, etc.) -->

## Code Style

<!-- TODO: Reference the project's style guide or linting configuration -->

## Testing

` + "```bash" + `
# TODO: Add test commands
` + "```" + `

### Test Coverage Requirements

<!-- TODO: Specify minimum coverage or testing expectations -->

## Issue Reporting

<!-- TODO: Describe how to file bug reports and feature requests -->

## License

<!-- TODO: Specify the license and CLA requirements if any -->
`
