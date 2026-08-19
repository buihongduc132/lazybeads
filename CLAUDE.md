# LazyBeads

A TUI (Terminal User Interface) for managing beads issues, built with Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea).

## Project Structure

```
lazybeads/
├── main.go              # Entry point, CLI flags, initialization
├── internal/
│   ├── app/             # Main Bubble Tea application model
│   ├── beads/           # Client wrapper for bd CLI
│   ├── models/          # Shared data models
│   └── ui/              # Reusable UI components
└── .beads/              # Issue tracking (managed by bd)
```

## Development

### Prerequisites
- Go 1.25+
- `bd` CLI installed and available in PATH

### Build and Install
```bash
# Build locally
go build .

# Install globally (use this to test the app)
go install .
```

### Validation
```bash
# Run headless validation to verify bd integration works
lazybeads --check
```

## Issue Tracking

This project uses **beads** (`bd`) for issue tracking. All work should be tracked through beads.

### Common Commands
```bash
bd ready                  # Find work ready to start
bd show <id>              # View issue details
bd update <id> --claim    # Claim and start work
bd close <id>             # Complete work
bd sync --from-main       # Sync beads from main branch
```

### Workflow
1. Find available work with `bd ready`
2. Claim an issue before starting
3. Implement, test, commit
4. Close the issue when done
5. Sync and commit before ending session

<!-- gitnexus:start -->
# GitNexus — Code Intelligence

This project is indexed by GitNexus as **lazybeads** (786 symbols, 2514 relationships, 57 execution flows). Use the GitNexus MCP tools to understand code, assess impact, and navigate safely.

> If any GitNexus tool warns the index is stale, run `npx gitnexus analyze` in terminal first.

## Always Do

- **MUST run impact analysis before editing any symbol.** Before modifying a function, class, or method, run `gitnexus_impact({target: "symbolName", direction: "upstream"})` and report the blast radius (direct callers, affected processes, risk level) to the user.
- **MUST run `gitnexus_detect_changes()` before committing** to verify your changes only affect expected symbols and execution flows.
- **MUST warn the user** if impact analysis returns HIGH or CRITICAL risk before proceeding with edits.
- When exploring unfamiliar code, use `gitnexus_query({query: "concept"})` to find execution flows instead of grepping. It returns process-grouped results ranked by relevance.
- When you need full context on a specific symbol — callers, callees, which execution flows it participates in — use `gitnexus_context({name: "symbolName"})`.

## Never Do

- NEVER edit a function, class, or method without first running `gitnexus_impact` on it.
- NEVER ignore HIGH or CRITICAL risk warnings from impact analysis.
- NEVER rename symbols with find-and-replace — use `gitnexus_rename` which understands the call graph.
- NEVER commit changes without running `gitnexus_detect_changes()` to check affected scope.

## Resources

| Resource | Use for |
|----------|---------|
| `gitnexus://repo/lazybeads/context` | Codebase overview, check index freshness |
| `gitnexus://repo/lazybeads/clusters` | All functional areas |
| `gitnexus://repo/lazybeads/processes` | All execution flows |
| `gitnexus://repo/lazybeads/process/{name}` | Step-by-step execution trace |

## CLI

| Task | Read this skill file |
|------|---------------------|
| Understand architecture / "How does X work?" | `.claude/skills/gitnexus/gitnexus-exploring/SKILL.md` |
| Blast radius / "What breaks if I change X?" | `.claude/skills/gitnexus/gitnexus-impact-analysis/SKILL.md` |
| Trace bugs / "Why is X failing?" | `.claude/skills/gitnexus/gitnexus-debugging/SKILL.md` |
| Rename / extract / split / refactor | `.claude/skills/gitnexus/gitnexus-refactoring/SKILL.md` |
| Tools, resources, schema reference | `.claude/skills/gitnexus/gitnexus-guide/SKILL.md` |
| Index, status, clean, wiki CLI commands | `.claude/skills/gitnexus/gitnexus-cli/SKILL.md` |

<!-- gitnexus:end -->
