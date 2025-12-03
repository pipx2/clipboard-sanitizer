# Clipboard Sanitizer

A cross-platform tool to sanitize your clipboard content.
It allows you to redact, mask, or replace sensitive information before pasting it anywhere.

## Features

- Redact, mask, or whitespace sensitive content from the clipboard.
- Expandable: easily add new sanitization rules (regex, emails, IPs, credit cards, etc.).
- Simple GUI for quick use.

## Requirements

- Go 1.20+ (see `go.mod` for module details)
- Cross-platform: works on Linux, macOS, and Windows

## Build

### First-time build

```bash
git clone https://github.com/pipx2/clipboard-sanitizer.git
cd clipboard-sanitizer
go build -o clipboard-sanitizer ./cmd/sanitizer
```

### Subsequent builds

Use the provided build script:

```bash
./build.sh
```

## Usage

Run the application:

```bash
./clipboard-sanitizer
```

The GUI will open, allowing you to select sanitization mode and sanitize your clipboard.

## Contribution

Feel free to submit issues or pull requests to improve sanitization rules or add new features.

## License

MIT License.
