.PHONY: build install clean help dev

help:
	@echo "Nova Horizon CLI - Makefile Commands"
	@echo "===================================="
	@echo "make build       - Build the binary"
	@echo "make install     - Build and install to $(PREFIX)/bin"
	@echo "make dev         - Run with hot reload (requires entr)"
	@echo "make clean       - Remove build artifacts"

build:
	CGO_ENABLED=0 go build -o nova-hrzn .

install: build
	@mkdir -p $(HOME)/.local/bin
	cp nova-hrzn $(HOME)/.local/bin/
	@echo "Installed to $(HOME)/.local/bin/nova-hrzn"
	@echo "Add $(HOME)/.local/bin to your PATH if not already added"

clean:
	rm -f nova-hrzn nova-horizon

dev:
	@which entr > /dev/null || (echo "entr not found. Install with: brew install entr (macOS) or apt install entr (Linux)"; exit 1)
	find . -name "*.go" | entr -r make build
