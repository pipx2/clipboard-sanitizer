if [ -f "clipboard-sanitizer" ]; then
    echo "Removing clipboard-sanitizer"
    rm "clipboard-sanitizer"
fi
echo "Build clipboard-sanitizer..."
go build -o clipboard-sanitizer ./cmd/sanitizer/
