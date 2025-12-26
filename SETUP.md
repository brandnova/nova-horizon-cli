# Setup Guide for Nova Horizon

## Step 1: Install Go (if not already installed)

### macOS

```bash
brew install go
```

### Fedora/RHEL

```bash
sudo dnf install golang
```

### Ubuntu/Debian

```bash
sudo apt-get update
sudo apt-get install golang-go
```

### Windows

1. Download installer from https://golang.org/dl/
2. Run the installer
3. Verify installation:
   ```powershell
   go version
   ```

## Step 2: Get Your Gemini API Key

1. Open https://aistudio.google.com in your browser
2. Click on **"Get API key"** (top left)
3. Click **"Create API key in new project"**
4. Copy the generated API key
5. **Keep this secret!** Never share it or commit it to version control

## Step 3: Clone the Project

```bash
git clone https://github.com/brandnova/nova-horizon-cli.git
cd nova-horizon
```

## Step 4: Build and Install

```bash
make install
```

This will:

- Build the Go binary
- Install it to `~/.local/bin/nova-horizon`
- Create a static binary (works without Go installed)

## Step 5: Configure API Key

### Method A: Environment Variable (Simple, for testing)

Add to your shell profile (~/.bashrc, ~/.zshrc, ~/.fish/config.fish):

```bash
export GEMINI_API_KEY="your-actual-api-key-here"
```

Then reload your shell:

```bash
source ~/.bashrc  # or ~/.zshrc
```

### Method B: Config File (Recommended for regular use)

Create the config directory:

```bash
mkdir -p ~/.config/nova-horizon
```

Create `~/.config/nova-horizon/config.toml`:

```toml
api_key = "your-actual-api-key-here"
```

Set proper permissions:

```bash
chmod 600 ~/.config/nova-horizon/config.toml
```

## Step 6: Verify Installation

Test that everything works:

```bash
nova-horizon "List all files in the current directory"
```

You should see:

1. Banner showing "Nova Horizon" and working directory
2. The agent making function calls
3. Results of your request

## Step 7: Verify PATH Setup

Make sure `~/.local/bin` is in your PATH:

```bash
echo $PATH | grep ".local/bin"
```

If not present, add this to your shell profile:

```bash
export PATH="$HOME/.local/bin:$PATH"
```

Then reload:

```bash
source ~/.bashrc  # or ~/.zshrc
```

## Step 8: Optional - Install Runtime Dependencies

For full functionality, install these (depending on what you want to do):

```bash
# Python
python3 --version

# Node.js
node --version

# TypeScript (if needed)
npm install -g ts-node

# Bash (usually pre-installed)
bash --version
```

## Verification Checklist

- [ ] Go is installed (`go version`)
- [ ] nova-horizon binary exists (`which nova-horizon`)
- [ ] API key is set or config file exists
- [ ] Can run: `nova-horizon "hello"` without errors
- [ ] `~/.local/bin` is in your PATH
- [ ] Runtime dependencies installed (Python, Node, etc.)

## Testing

### Test 1: Simple File Listing

```bash
nova-horizon "List all files in this directory"
```

### Test 2: File Creation

```bash
nova-horizon "Create a file called test.txt with the content 'Hello World'"
```

### Test 3: File Reading

```bash
nova-horizon "Read the test.txt file I just created"
```

### Test 4: Code Execution (with --allow-run)

```bash
nova-horizon --allow-run "Create a simple Python script that prints hello world and run it"
```

## Troubleshooting

### Issue: "command not found: nova-horizon"

**Solution:**

1. Check that `make install` completed successfully
2. Check that `~/.local/bin` is in your PATH: `echo $PATH`
3. Try the full path: `~/.local/bin/nova-horizon --help`
4. If full path works, add to shell profile: `export PATH="$HOME/.local/bin:$PATH"`

### Issue: "GEMINI_API_KEY not set"

**Solution:**

1. Check if environment variable is set: `echo $GEMINI_API_KEY`
2. Try config file method (see Step 5B)
3. Make sure you used the correct key from aistudio.google.com

### Issue: "invalid API key"

**Solution:**

1. Double-check the key from aistudio.google.com
2. Make sure there are no extra spaces or quotes
3. The key should start with `AIza...`
4. Try creating a new key if the old one seems invalid

### Issue: "build failed"

**Solution:**

1. Update Go: `go version` should be 1.23 or higher
2. Run: `go mod download`
3. Try: `go mod tidy`
4. Then: `make build`

## Next Steps

Once set up, check the README.md for:

- Usage examples
- Command reference
- How it works
- Advanced features

Happy coding!
