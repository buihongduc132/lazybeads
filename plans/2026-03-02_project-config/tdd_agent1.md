# TDD Test Cases: Project-Level Config (Loading, Precedence, CustomCommands)

Verbatim user intention:
Create TDD test cases for implementing project-level configuration in lazybeads. Focus on config loading, precedence, and customCommands behavior.

Assumptions (must be confirmed before implementation)
1. Project config filename is `.lazybeads.yml` and lives at the project root (or nearest ancestor directory).
2. Discovery starts at the process working directory and walks up parent directories until repo root (or filesystem root if repo root is not detectable).
3. Precedence order is: explicit `LAZYBEADS_CONFIG` path (highest, no project search) -> global user config (`~/.config/lazybeads/config.yml`) -> project config overlays global.
4. `customCommands` merge rule is: start with global list, overlay project list by `key` (project wins on conflicts), preserve stable order for bindings.

## Config Loading
1. Loads project config when `.lazybeads.yml` exists in the current working directory.
2. Walks up parent directories and uses the nearest `.lazybeads.yml` when multiple exist (child overrides parent).
3. When no project config exists, falls back to global config only (existing behavior stays intact).
4. When both global and project configs exist, loader returns a merged config (not just one source).
5. When `LAZYBEADS_CONFIG` is set to an explicit path, project discovery is skipped and only that file is loaded.
6. When explicit `LAZYBEADS_CONFIG` path does not exist, loader returns a clear error (path in message).
7. When project config exists but is unreadable, loader returns a clear error (path in message).
8. When project config exists but has invalid YAML, loader returns a clear parse error (path and line/column if available).

## Precedence
1. Project config overrides global config for overlapping fields (verify with conflicting `customCommands` entries by key).
2. Global config values remain when project config omits a field (verify with a non-overlapping `customCommands` entry).
3. Precedence is deterministic across runs (same global + project inputs always produce same output ordering and values).

## CustomCommands Behavior
1. `customCommands` from both global and project configs are included in the final config (union by key).
2. Duplicate `customCommands` keys across global and project resolve to the project version (same key, different description/command).
3. `customCommands` with missing `context` default to `list` for both global and project sources.
4. Custom command bindings reflect the merged list order (global first, then project overrides/replacements), ensuring keybinding help uses the project version for overridden keys.
5. Empty `customCommands` in project config does not erase global commands unless explicitly specified by the merge rule (if explicit clear is desired, add a separate test and behavior decision).

Documented changes:
TDD test cases for implementing project-level configuration in lazybeads. Focus on config loading, precedence, and customCommands behavior.
