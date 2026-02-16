package manual

const gettingStartedTmpl = `---
title: Getting Started
description: Quick start guide for new users of this framework
tags: getting-started, quickstart, tutorial
---

# Getting Started

## Installation

` + "```bash" + `
# TODO: Add installation commands
` + "```" + `

## Create Your First Project

` + "```bash" + `
# TODO: Add project scaffolding commands
` + "```" + `

## Project Structure

` + "```" + `
# TODO: Add annotated directory tree for a new project
` + "```" + `

## Hello World Example

` + "```" + `
// TODO: Add a minimal hello world example
` + "```" + `

## Running the Application

` + "```bash" + `
# TODO: Add run commands
` + "```" + `

## Next Steps

- <!-- TODO: Link to key documentation sections -->
`

const pluginSystemTmpl = `---
title: Plugin System
description: Plugin architecture, extension points, and authoring guide
tags: plugins, extensions, hooks
---

# Plugin System

## Architecture

<!-- TODO: Describe the plugin architecture and lifecycle -->

## Extension Points

| Hook | When It Runs | Use Case |
|------|-------------|----------|
|      |             |          |

## Writing a Plugin

### Plugin Interface

` + "```" + `
// TODO: Add plugin interface definition
` + "```" + `

### Minimal Plugin Example

` + "```" + `
// TODO: Add a complete minimal plugin example
` + "```" + `

### Registering a Plugin

` + "```" + `
// TODO: Add plugin registration code
` + "```" + `

## Built-in Plugins

| Plugin | Description | Enabled by Default |
|--------|------------|-------------------|
|        |            |                   |

## Testing Plugins

` + "```" + `
// TODO: Add plugin testing example
` + "```" + `
`

const migrationGuideTmpl = `---
title: Migration Guide
description: Upgrade paths and breaking change migration instructions
tags: migration, upgrade, breaking-changes
---

# Migration Guide

## Upgrade Path

| From | To | Difficulty | Guide |
|------|----|-----------|-------|
|      |    |           |       |

## Latest Migration

### Migrating from vX to vY

#### Breaking Changes

- <!-- TODO: List breaking changes -->

#### Step-by-Step

1. Update dependencies

` + "```bash" + `
# TODO: Add dependency update commands
` + "```" + `

2. Update configuration

<!-- TODO: Document config changes -->

3. Update code

<!-- TODO: Document code changes with before/after examples -->

#### Deprecated Features

| Feature | Replacement | Remove In |
|---------|------------|-----------|
|         |            |           |

## Troubleshooting Upgrades

<!-- TODO: Document common upgrade issues and fixes -->
`
