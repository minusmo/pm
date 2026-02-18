# pm — Project Manual CLI

A CLI tool for managing project-specific runbooks and manuals stored in `.pm/` directories.

## Quick Reference

- **Language:** Go 1.25.6
- **CLI framework:** [cobra](https://github.com/spf13/cobra)
- **Module:** `github.com/hojooneum/pm`
- **License:** MIT
- **Build:** `go build -o pm .`
- **Test:** `go test ./...`
- **Lint:** `go vet ./... && test -z "$(gofmt -l .)"`

## Commands

| Command | Description |
|---------|-------------|
| `pm` | Show project summary (or prompt interactive init if no `.pm/`) |
| `pm init` | Scaffold `.pm/` from a template (`--template`, `--list-templates`) |
| `pm list [group]` | List sections (alias: `ls`) |
| `pm open <section>` | Display a section's content |
| `pm edit <section>` | Open a section in `$EDITOR` for editing |
| `pm search <keyword>` | Case-insensitive keyword search across all sections |

## Project Structure

```
main.go                          Entry point
cmd/
  root.go                        Root command, interactive init flow, loadAllSections helper
  init.go                        pm init — scaffold .pm/ from templates
  list.go                        pm list — list sections
  open.go                        pm open — display a section
  edit.go                        pm edit — open a section in $EDITOR
  search.go                      pm search — keyword search
internal/
  cli/
    format.go                    Output formatting (section lists, search results, summaries)
    prompt.go                    Interactive prompts (ConfirmYesNo, SelectOption)
  fs/
    fs.go                        Filesystem ops (.pm/ detection, read, search, write)
    groups.go                    Group listing with sort order (core first, custom last)
  manual/
    section.go                   Section model and frontmatter parser
    template.go                  DefaultTemplates map and CoreSectionOrder
    template_core.go             Core section templates (overview, deploy, troubleshoot, etc.)
    template_onboarding.go       Onboarding templates (setup-guide, codebase-walkthrough, etc.)
    template_microservice.go     Microservice templates (service-dependencies, api-contracts, etc.)
    template_library.go          Library templates (api-reference, usage-examples, versioning, etc.)
    template_framework.go        Framework templates (getting-started, plugin-system, migration-guide)
    preset.go                    Built-in presets (default, minimal, onboarding, microservice, library, framework)
```

## Key Concepts

- **`.pm/` directory:** Lives at project root. Contains grouped markdown files that serve as the project manual.
- **Groups:** Subdirectories under `.pm/` (e.g., `core/`, `custom/`). `core` sorts first, `custom` sorts last, others alphabetical.
- **Sections:** Markdown files within groups. Have YAML-style frontmatter (`title`, `description`, `tags`) delimited by `---`.
- **Templates:** Define which sections to scaffold on `pm init`. 6 built-in presets: `default` (7 core ops/SRE sections), `minimal` (3 sections), `onboarding` (6 sections), `microservice` (9 sections), `library` (7 sections), `framework` (7 sections). Custom templates via JSON files.

## Conventions

- All commands check for `.pm/` existence before operating; print a helpful message if missing.
- Section names must match `^[a-z0-9][a-z0-9-]*$`.
- `FindSection` resolves names case-insensitively, checking all groups.
- Interactive prompts retry up to 3 times on invalid input, then fall back to defaults.
- Version is injected via ldflags: `-X github.com/hojooneum/pm/cmd.version={{.Version}}`.

## CI

GitHub Actions runs `go vet`, `gofmt`, and `go test` on push to main and PRs.

## Releasing a New Version

1. Make sure all changes are committed and CI is green on `main`.
2. Decide the version number following semver (e.g., `v0.2.0`).
3. Create and push a git tag:
   ```bash
   git tag v0.2.0
   git push origin v0.2.0
   ```
4. The `release` GitHub Actions workflow triggers automatically on `v*` tags.
5. Goreleaser builds binaries for darwin/linux (amd64/arm64), creates a GitHub Release with archives and checksums, and updates the Homebrew formula in `hojooneum/homebrew-tap`.
6. Verify the release at `https://github.com/hojooneum/pm/releases`.
7. Users can then install/upgrade via:
   ```bash
   brew install hojooneum/tap/pm
   # or
   brew upgrade pm
   ```
