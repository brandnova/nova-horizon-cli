## Multi-Language Support Implementation Plan

```markdown
# Multi-Language Support Implementation Plan

## Overview
Extend Nova Horizon to support multiple programming languages (Python, Go, JavaScript/Node.js, Rust, Java, C++, etc.) with language-specific execution, file operations, and tool availability.

## Phase 1: Language Detection & Registry

### 1.1 Language Identifier
- Create a `Language` enum with supported languages
- Implement file extension mapping (`.py` → Python, `.go` → Go, `.js` → JavaScript, etc.)
- Add language detection logic based on file content or explicit flag

### 1.2 Language Configuration Registry
```go
type Language struct {
    Name          string
    Extensions    []string
    Interpreter   string
    RunCommand    string
    Version       string
    SyntaxChecker string
}

type LanguageRegistry struct {
    languages map[string]*Language
}
```

### 1.3 Supported Languages (Initial)

- **Python**: `python3`, use `python -m py_compile` for syntax check
- **Go**: `go run`, `go build`, built-in format check
- **JavaScript/Node.js**: `node`, use `node --check` for syntax validation
- **Rust**: `rustc`, `cargo run`, type-safe compilation
- **Bash/Shell**: `bash`, `sh` with syntax validation
- **Java**: `javac`, `java` with compiled bytecode execution


## Phase 2: Language-Specific Tool Execution

### 2.1 Execution Engine Abstraction

- Create `ExecutionEngine` interface with methods for each language
- Implement language-specific runners with proper:

- Environment setup (PYTHONPATH, GOPATH, NODE_PATH, etc.)
- Module/dependency management detection
- Standard library availability checks





### 2.2 Timeout & Resource Management

- Language-specific timeouts (Go: 30s, Node: 30s, Python: 30s)
- Memory limits per language
- CPU usage constraints


### 2.3 Syntax Validation

- Language-specific pre-execution validation
- Compilation checks for compiled languages
- Linting integration for optional warnings


## Phase 3: File Operations Per Language

### 3.1 Language-Aware File Creation

- Auto-generate language-specific boilerplate
- Standard library imports for each language
- Proper encoding and line endings


### 3.2 Dependency/Package Management Detection

- Detect `requirements.txt`, `go.mod`, `package.json`, `Cargo.toml`, `pom.xml`
- Alert user if dependencies are needed
- Option to install dependencies before execution


### 3.3 Project Structure Recognition

- Python: detect `src/`, `tests/`, `__init__.py` patterns
- Go: recognize module structure, `main.go` entry point
- Node.js: detect `src/`, `package.json` configuration
- Rust: recognize `Cargo.toml` project layout


## Phase 4: Gemini Prompt Enhancement

### 4.1 Language Context Injection

- Include current language context in prompts
- Add language-specific constraints to tool descriptions
- Example: "The user is working with Go. Consider Go idioms and conventions."


### 4.2 Multi-File Project Support

- Allow agents to understand project structure across languages
- Support creating multiple files in related projects
- Handle cross-language dependencies


## Phase 5: CLI Enhancements

### 5.1 New Flags

- `--lang <language>`: Explicitly specify language
- `--no-execute`: Create files without running
- `--install-deps`: Auto-install missing dependencies
- `--show-interpreter`: Display interpreter path before execution


### 5.2 Auto-Detection Strategy

1. Check `--lang` flag
2. Inspect existing files in directory
3. Look for package managers (go.mod, package.json, etc.)
4. Prompt user if ambiguous


## Phase 6: Safety & Guardrails

### 6.1 Language-Specific Restrictions

- Limit file access per language sandbox
- Prevent network requests (unless explicitly allowed)
- Restrict system calls in interpreted languages


### 6.2 Dangerous Pattern Detection

- SQL injection patterns
- System command injection
- Unsafe file operations
- Network requests without warning


## Phase 7: Output & Logging

### 7.1 Language-Aware Output Formatting

- Syntax highlighting per language in output
- Language-specific error message interpretation
- Stack trace parsing and formatting


### 7.2 Execution Metadata

- Interpreter version detection
- Compilation time vs runtime breakdown
- Memory/CPU usage per language


## Implementation Order

1. Create Language Registry & Detection (easy wins)
2. Implement Node.js support (similar to Python)
3. Add Go execution (already have binary knowledge)
4. Bash/Shell support (minimal risk)
5. Add Rust & Java (more complex)
6. Enhance prompts with language context
7. Add dependency management
8. Implement advanced safety features


## Backward Compatibility

- Maintain existing Python behavior as default
- Auto-detect language from files if not specified
- Error gracefully if language not supported yet