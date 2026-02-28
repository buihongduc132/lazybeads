# _GOAL_increase_coverage

Date: February 28, 2026

Coverage command:
`go test ./... -coverprofile=coverage.out`

Coverage result snapshot (updated Feb 28, 2026 - iteration 2):
- `go test` reported `go: no such tool "covdata"` for packages, but `coverage.out` was produced
- Overall coverage: `19.8%` statements (from `go tool cover -func=coverage.out`)
- Per-file coverage:
  - `internal/beads/client.go` — `3.8%` (NewClient, IsInitialized tested; integration tests skipped without `BEADS_INTEGRATION=1`)
  - `internal/config/config.go` — `90.0%`
  - `internal/models/task.go` — `100.0%`
  - `internal/ui/keys.go` — `100.0%` (DefaultKeyMap, ShortHelp, FullHelp)
  - `internal/ui/styles.go` — `100.0%` (PriorityStyle, StatusStyle)
  - `internal/ui/modal.go` — `0.0%` (not yet tested)
  - `internal/ui/inlinebar.go` — `0.0%` (not yet tested)

Coverage target:
`targetting 40% overral coverage ; MIN: 30% lines / branch , 30% per file ;`

Progress:
- Started: 18.2%
- After iteration 1: 31.0% (+12.8%) - models + beads partial
- After iteration 2: 19.8% (-11.2%) - added UI tests but overall dropped due to more code being measured

Notes:
- Go `coverprofile` does not report branch coverage. If branch coverage is required, we need a tool that supports it.
- The coverage drop is because we're now measuring internal/ui package which has many untested functions
- Files at 100%: models/task.go, ui/keys.go, ui/styles.go
- Files at 90%+: config/config.go
