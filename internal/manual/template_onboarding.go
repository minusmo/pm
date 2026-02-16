package manual

const setupGuideTmpl = `---
title: Setup Guide
description: Local development environment setup instructions
tags: setup, install, environment
---

# Setup Guide

## Prerequisites

| Tool | Version | Install |
|------|---------|---------|
|      |         |         |

## Clone & Install

` + "```bash" + `
# TODO: Add clone and install commands
` + "```" + `

## Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
|          |             |         |

### Local Config Files

<!-- TODO: Document config files that need to be created or modified -->

## Verify Setup

` + "```bash" + `
# TODO: Add commands to verify the setup is working
` + "```" + `

## Common Setup Issues

<!-- TODO: Document frequently encountered setup problems and fixes -->
`

const codebaseWalkthroughTmpl = `---
title: Codebase Walkthrough
description: Guided tour of the project structure and key modules
tags: codebase, structure, architecture
---

# Codebase Walkthrough

## Directory Structure

` + "```" + `
# TODO: Add annotated directory tree
` + "```" + `

## Key Modules

### Module: Core

<!-- TODO: Describe the core module, its responsibilities, and entry points -->

### Module: API

<!-- TODO: Describe the API layer -->

### Module: Data

<!-- TODO: Describe the data/persistence layer -->

## Data Flow

<!-- TODO: Describe how a typical request flows through the system -->

## Key Files to Read First

| File | Why |
|------|-----|
|      |     |
`

const devWorkflowTmpl = `---
title: Development Workflow
description: Day-to-day development process and branch strategy
tags: workflow, git, branching
---

# Development Workflow

## Branch Strategy

| Branch | Purpose | Merges Into |
|--------|---------|-------------|
| main   |         |             |
| develop|         |             |
| feature|         |             |

## Making Changes

### 1. Create a Branch

` + "```bash" + `
# TODO: Add branch creation commands
` + "```" + `

### 2. Develop & Test

` + "```bash" + `
# TODO: Add build and test commands
` + "```" + `

### 3. Submit a Pull Request

<!-- TODO: Describe PR process, reviewers, CI checks -->

## Code Review Guidelines

- <!-- TODO: Add review expectations -->

## CI/CD Pipeline

| Stage | Trigger | What It Does |
|-------|---------|-------------|
|       |         |             |
`

const codingConventionsTmpl = `---
title: Coding Conventions
description: Code style, naming, and best practices for this project
tags: conventions, style, standards
---

# Coding Conventions

## Language & Style

<!-- TODO: Specify the language version and style guide followed -->

## Naming Conventions

| Element | Convention | Example |
|---------|-----------|---------|
| Files   |           |         |
| Functions|          |         |
| Variables|          |         |
| Constants|          |         |

## Project Patterns

### Error Handling

<!-- TODO: Describe the project's error handling approach -->

### Logging

<!-- TODO: Describe logging conventions and levels -->

### Testing

<!-- TODO: Describe testing conventions (unit, integration, naming) -->

## Linting & Formatting

` + "```bash" + `
# TODO: Add lint and format commands
` + "```" + `
`
