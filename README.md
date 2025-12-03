# Clipboard Sanitizer

Cross-platform clipboard sanitizer.

## Features

- Minimal example: trims leading/trailing whitespace
- Expandable: add regex rules, email/IP redaction, credit card masking, etc.

## Requirements

See `go.mod`.

## Build

```bash
git clone https://github.com/pipx2/clipboard-sanitizer.git
cd clipboard-sanitizer
go build -o clipboard-sanitizer ./cmd/sanitizer
```

## Usage

### Run

```bash
./clipboard-sanitizer
```

### Sanitize

- Copy text to your clipboard.
- Press the 'Sanitize Now' button.
