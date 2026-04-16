# scribe — Copilot Instructions

## Overview

`scribe` is a Go library for structured, indented output in CLI tools. It writes to any `io.Writer`
and formats output as a tree of describe blocks and leaf print/error lines, with each nesting level
indented by two spaces. All visual styling is delegated to a caller-supplied `Theme`, so the same
structured output can be rendered plain or with ANSI color without changing the calling code.

## API Surface

### `Scriber` interface (`scriber.go`)

```go
type Scriber interface {
    BeginDescribe(desc string)
    EndDescribe()
    Print(str string)
    PrintLines(buf *bytes.Buffer)
    Error(err error)
}
```

- `BeginDescribe` / `EndDescribe` — open/close a named section; increases/decreases the indent level.
  Each call to `BeginDescribe` emits the section label and then increments indent; `EndDescribe`
  decrements it. Nest calls arbitrarily.
- `Print` — emit a single indented line styled by `Theme.Print`.
- `PrintLines` — scan a `*bytes.Buffer` line-by-line and emit each line with `Theme.Print` styling.
- `Error` — emit a single indented line styled by `Theme.Error`.

### Constructor

```go
func NewScribe(writer io.Writer, theme *Theme) (Scriber, error)
```

Validates the theme (returns an error if any decorator is nil) and returns a `Scriber`.

### `Theme` (`theme.go`)

```go
type Theme struct {
    Describe func(string) string
    Print    func(string) string
    Error    func(error) string
}
```

All three fields are required — `ValidateTheme` returns a sentinel error for any nil field:

```go
var (
    ErrThemeDescribeMissing = fmt.Errorf("theme missing describe decorator")
    ErrThemePrintMissing    = fmt.Errorf("theme missing print decorator")
    ErrThemeErrorMissing    = fmt.Errorf("theme missing error decorator")
)
```

`DefaultTheme` uses `NoopDecorator` / `NoopErrDecorator` — no color, strings pass through unchanged.

### `color` subpackage (`color/color.go`)

Provides raw ANSI escape helpers (`RedFg`, `CyanFg`, `HiBlueFg`, etc.) and the standard and
high-intensity foreground color constants. Intended for use when constructing custom `Theme`
decorator functions. It is not coupled to the `Scribe` struct in any way.

## Build & Validate

```sh
go build ./...
go vet ./...
go test ./...
```

There is no linter config in this repo. Run the above three in order after any Go change.

## Repository Layout

```
scribe.go       Scribe struct; implements the Scriber interface
scriber.go      Scriber interface definition
theme.go        Theme struct, DefaultTheme, sentinel errors, ValidateTheme
scribe_test.go  Tests: Describe nesting, full output tree, themed output
color/
  color.go      ANSI escape color helpers (FgXxx constants + XxxFg wrapping fns)
  color_test.go
vendor/         Vendored dependencies (never edit directly)
```

## Key Conventions

- **Error wrapping**: `fmt.Errorf("methodName: operation: %w", err)` — lowercase, no period
- **No blank identifier**: use `for i := range`, never silently discard errors
- **No goroutines** unless there is a demonstrated need
- **No `init()`** unless strictly unavoidable
- **Vendor is read-only**: `go mod vendor` is the only permitted way to modify `vendor/`
- **After any Go change**: run `go fmt`, `go vet`, `go build ./...`, `go test ./...` in that order
- Tests call `t.Parallel()` at the top of each test function
