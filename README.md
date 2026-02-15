# pm

A CLI tool for managing project-specific runbooks and manuals stored in `.pm/` directories.

## Install

**Homebrew:**

```bash
brew install hojooneum/tap/pm
```

**Go:**

```bash
go install github.com/hojooneum/pm@latest
```

## Quick Start

```bash
# Initialize a project manual in the current directory
pm init

# List all sections
pm list

# Open a specific section
pm open deploy

# Search across all sections
pm search kubernetes
```

Running `pm` with no arguments shows a project summary, or prompts you to initialize if no `.pm/` directory exists.

## Commands

| Command | Description |
|---|---|
| `pm` | Show project summary (interactive init if no `.pm/`) |
| `pm init` | Scaffold a `.pm/` directory from a template |
| `pm list [group]` | List available sections (alias: `ls`) |
| `pm open <section>` | Display a section's content |
| `pm search <keyword>` | Search for a keyword across all sections |

### pm init

```bash
pm init                          # Use the default template
pm init --template minimal       # Use a built-in preset
pm init --template my-tmpl.json  # Use a custom JSON template
pm init --list-templates         # List available presets
```

## How It Works

`pm` stores project documentation in a `.pm/` directory at your project root.

**Groups** are subdirectories under `.pm/` (e.g., `core/`, `custom/`). They organize sections by category. `core` sorts first, `custom` sorts last, and everything else is alphabetical.

**Sections** are markdown files within groups. Each section can have YAML frontmatter with `title`, `description`, and `tags`:

```markdown
---
title: Deployment Guide
description: Step-by-step deployment procedures
tags: deploy, release
---

# Deployment Guide

Your content here...
```

Section names are resolved case-insensitively, so `pm open Deploy` and `pm open deploy` both work.

## Templates

Templates define which sections to scaffold when running `pm init`.

**Built-in presets:**

| Preset | Sections | Description |
|---|---|---|
| `default` | 7 | overview, deploy, troubleshoot, backup, maintenance, monitoring, contacts |
| `minimal` | 3 | overview, deploy, contacts |

**Custom templates** can be provided as JSON files:

```json
{
  "name": "my-template",
  "sections": [
    {
      "name": "overview",
      "group": "core",
      "title": "Project Overview",
      "description": "High-level summary",
      "tags": ["overview"]
    }
  ]
}
```

Section names must match `[a-z0-9][a-z0-9-]*`.

## License

[MIT](LICENSE)
