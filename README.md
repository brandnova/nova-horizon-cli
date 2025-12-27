# Nova Horizon

A local AI workflow agent powered by Google's Gemini API. Built in Go with safety-first design. The future of CLI workflow boost with AI integration. 

## Features

- **Smart Code Agent**: Uses Gemini API to understand and execute your requests
- **File Operations**: List directories, read files, write files safely
- **Program Execution**: Run Python, Go, Node.js, Bash, and TypeScript scripts
- **Safety First**: Path validation, file size limits, execution timeouts
- **Interactive Shell**: Built-in shell for seamless interaction
- **Flexible Configuration**: Environment variables or config file
- **CLI-Optimized**: Static binary, works from anywhere

## Prerequisites

- **Go 1.23+** (Required for building)
- **Google Gemini API Key** (Required for usage)
- **Runtime Dependencies**: Python3, Node.js, Bash, etc., depending on what code you want the agent to write and run.

## Installation

### 1. Build the Binary

#### Linux (Fedora / RHEL)

```bash
sudo dnf install golang
git clone https://github.com/brandnova/nova-horizon-cli.git
cd nova-horizon-cli
go build -o nova-hrzn .
```

#### Windows

1.  Install Go from [golang.org/dl](https://golang.org/dl/).
2.  Clone the repository:
    ```powershell
    git clone https://github.com/brandnova/nova-horizon-cli.git
    cd nova-horizon-cli
    ```
3.  Build the executable:
    ```powershell
    go build -o nova-hrzn.exe .
    ```

#### macOS / Ubuntu (using Make)

```bash
git clone https://github.com/brandnova/nova-horizon-cli.git
cd nova-horizon-cli
make install
```

### 2. Install to PATH

#### Linux / macOS

Copy the binary to a directory in your PATH (e.g., `~/.local/bin`):

```bash
mkdir -p ~/.local/bin
cp nova-hrzn ~/.local/bin/
```

Ensure `~/.local/bin` is in your PATH in `~/.bashrc` or `~/.zshrc`:

```bash
export PATH="$HOME/.local/bin:$PATH"
```

#### Windows

Move `nova-hrzn.exe` to a folder (e.g., `C:\GoApps\`) and add that folder to your System PATH environment variable.

## Configuration (Free usage)

You need a Google Gemini API key to use Nova Horizon.

### 1. Get Your API Key

1.  Visit [Google AI Studio](https://aistudio.google.com).
2.  Click **Get API key** -> **Create API key in new project**.
3.  Copy the key.

### 2. Set Up the Key

Choose **one** of the following methods:

#### Method A: Config File (Recommended)

Create `~/.config/nova-horizon/config.toml`:

```toml
api_key = "your-api-key-here"
```

_Note: Ensure this file is readable only by you (`chmod 600`)._

#### Method B: Environment Variable

```bash
export GEMINI_API_KEY="your-api-key-here"
```

## Usage

### Interactive Shell

Run without arguments to enter the interactive mode:

```bash
nova-hrzn
```

### Single Command

```bash
nova-hrzn "Create a Python script in this dir that calculates fibonacci"
```

### Command Flags

```bash
# Show info about Nova Horizon
nova-hrzn --info

# Verbose output (for debugging)
nova-hrzn -v "List all Python files"

# Dry run (show what would be done without making changes)
nova-hrzn --dry-run "Create a config file"

# Specify working directory
nova-hrzn --dir ./myproject "Refactor this code"

# Allow program execution (Safety: restricted to specific extensions)
nova-hrzn --allow-run "Run the test script"

# Auto-apply changes without confirmation
nova-hrzn --apply "Update all files"
```

## Troubleshooting

- **"command not found: nova-hrzn"**: Ensure the binary is in a folder included in your `PATH`.
- **"GEMINI_API_KEY not set"**: Check your `config.toml` or environment variable.
- **Build Failures**: Ensure you are using Go 1.23+.

## Architecture

```
nova-hrzn-cli/
├── cmd/           # CLI commands
├── internal/      # Core logic (Agent, Gemini Client, Tools)
├── main.go        # Entry point
└── README.md      # Documentation
```
