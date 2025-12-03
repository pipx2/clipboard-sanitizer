# Clipboard Sanitizer (Go)

Cross-platform clipboard sanitizer CLI written in Go.

## Features

- Automatically monitors clipboard and sanitizes new content
- Minimal example: trims leading/trailing whitespace
- Expandable: add regex rules, email/IP redaction, credit card masking, etc.
- Single-binary distribution, no Python dependencies

## Requirements

- Go 1.21+
- `github.com/atotto/clipboard` (declared in go.mod)

## Build

```bash
git clone https://github.com/yourusername/clipboard-sanitizer.git
cd clipboard-sanitizer
go build -o clipboard-sanitizer ./cmd/sanitizer
```

## Usage

```bash
./clipboard-sanitizer
```

- Copy text to your clipboard.
- Sanitizer automatically updates the clipboard.
- Stop with Ctrl+C.

## How to run

1. Clone and build:

```bash
git clone https://github.com/yourusername/clipboard-sanitizer.git
cd clipboard-sanitizer
go build -o clipboard-sanitizer ./cmd/sanitizer
```

2. Run the sanitizer

```bash
./clipboard-sanitizer
```

3. Copy something to the clipboard - it will sanitize automatically.

4. Press `Ctrl+C` to stop.
