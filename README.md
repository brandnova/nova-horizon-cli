# Nova Horizon

A local AI coding agent powered by Google's Gemini API. Built in Go with safety-first design.

## Features

- **Smart Code Agent**: Uses Gemini API to understand and execute your requests
- **File Operations**: List directories, read files, write files safely
- **Program Execution**: Run Python, Go, Node.js, Bash, and TypeScript scripts
- **Safety First**: Path validation, file size limits, execution timeouts
- **Flexible Configuration**: Environment variables or config file
- **CLI-Optimized**: Static binary, works from anywhere
- **Verbose Mode**: Detailed logging for debugging

## Prerequisites

- Go 1.23+
- Google Gemini API key
- Runtime dependencies based on your needs (Python3, Node.js, Bash, etc.)

## Installation

### Fedora (Linux)

```bash
sudo dnf install golang
git clone https://github.com/brandnova/nova-horizon-cli.git
cd nova-horizon-cli
go build -o nova-horizon .
```

### Windows

1. Install Go from https://golang.org/dl/
2. Open PowerShell or Command Prompt
3. Clone the repository:
   ```powershell
   git clone https://github.com/brandnova/nova-horizon-cli.git
   cd nova-horizon-cli
   ```
4. Build:
   ```powershell
   go build -o nova-horizon.exe .
   ```

### macOS / Ubuntu (using Make)

### 1. Clone and Build

```bash
git clone https://github.com/brandnova/nova-horizon-cli.git
cd nova-horizon-cli
make install
```

This installs the binary to `~/.local/bin/nova-horizon`. Ensure this directory is in your `PATH`:

```bash
export PATH="$HOME/.local/bin:$PATH"
```

Add this to your shell profile (~/.bashrc, ~/.zshrc, etc.) for persistence.

### 2. Set Up API Key

Choose one of these methods:

#### Option A: Environment Variable (Recommended for testing)

```bash
export GEMINI_API_KEY="your-api-key-here"
nova-horizon "Your prompt"
```

#### Option B: Config File (Recommended for permanent setup)

Create `~/.config/nova-horizon/config.toml`:

```toml
api_key = "your-api-key-here"
```

Then use nova-horizon without setting environment variables:

```bash
nova-horizon "Your prompt"
```

#### Getting Your API Key

1. Visit [Google AI Studio](https://aistudio.google.com)
2. Click "Get API key" → "Create API key in new project"
3. Copy the key and use it in one of the methods above

## Usage

### Basic Usage

```bash
nova-horizon "Create a Python script that calculates fibonacci"
```

### With Flags

```bash
# Verbose output
nova-horizon -v "List all Python files in src/"

# Dry run (show what would happen, don't make changes)
nova-horizon --dry-run "Create a config file"

# Specify working directory
nova-horizon --dir ./myproject "Refactor this code"

# Allow program execution
nova-horizon --allow-run "Run my test script"

# Custom model
nova-horizon --model "gemini-1.5-pro" "Your prompt"

# Set maximum steps (default 10)
nova-horizon --max-steps 20 "Complex task"

# Auto-apply changes without confirmation
nova-horizon --apply "Update all files"
```

### Common Examples

```bash
# List project structure
nova-horizon "Show me the structure of this project"

# Create a new file
nova-horizon "Create a .gitignore file for Python"

# Read and understand code
nova-horizon "What does the main function do?"

# Refactor code
nova-horizon "Simplify this function and make it more readable"

# Execute and debug
nova-horizon --allow-run "Run the tests and show me any failures"

# Use different working directory
nova-horizon --dir /path/to/project "Set up this new project"
```

## Command Reference

```
nova-horizon [flags] [prompt]

Flags:
  -d, --dir string          Working directory (default: current directory)
  -v, --verbose             Enable verbose output
      --dry-run             Show what would be done without making changes
      --model string        Model to use (default: "gemini-2.5-flash")
      --max-steps int       Maximum agent loop iterations (default: 10)
      --allow-run           Allow execution of programs
      --apply               Automatically apply file changes without confirmation
  -h, --help                Show this help message
```

## Architecture

```
nova-horizon-cli/
├── cmd/                    # CLI command definitions
│   └── root.go            # Main CLI entry point
├── internal/
│   ├── agent/             # Core agent loop
│   ├── config/            # Configuration management
│   ├── gemini/            # Gemini API client
│   ├── logger/            # Logging utilities
│   └── tools/             # File and execution tools
├── main.go                # Entry point
├── go.mod                 # Go module definition
├── Makefile               # Build automation
└── README.md              # This file
```

## How It Works

1. **Startup**: Loads your Gemini API key and validates the working directory
2. **User Prompt**: Takes your request and sends it to Gemini
3. **Agent Loop**: Gemini decides which tools to use (file read, write, execute)
4. **Execution**: Tools execute safely with path validation and size limits
5. **Feedback**: Results are sent back to Gemini for the next step
6. **Completion**: When done or max steps reached, outputs the result

## Safety Features

- **Path Validation**: All file operations stay within the working directory
- **File Size Limits**: Maximum 100KB per file to prevent abuse
- **Execution Timeout**: 30-second timeout on program execution
- **Extension Allowlist**: Only safe file types can be written
- **Execution Allowlist**: Only specific file types (.go, .py, .sh, .js, .ts) can be run
- **Loop Detection**: Detects and aborts infinite function call loops

## Troubleshooting

### "GEMINI_API_KEY not set"

Make sure you've configured your API key using one of the methods above:

```bash
# Quick test
export GEMINI_API_KEY="your-key"
nova-horizon "hello"
```

### "path traversal not allowed"

This means you tried to access a file outside the working directory. This is a security feature.

### "execution timeout"

A script took longer than 30 seconds. Use `--max-steps` to adjust agent iterations:

```bash
nova-horizon --max-steps 15 "Your prompt"
```

### "Model not found"

Verify your model string is correct. Default is `gemini-2.5-flash`. Other options:

- `gemini-1.5-flash`
- `gemini-1.5-pro`
- `gemini-2.0-flash-exp`

## Building from Source

```bash
# Build binary
make build

# Build and install
make install

# Clean build artifacts
make clean

# Development (requires entr)
make dev
```

## Static Binary

To create a completely static binary (no runtime dependencies):

```bash
CGO_ENABLED=0 go build -a -installsuffix cgo -o nova-horizon .
```
