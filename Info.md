# Nova Horizon Project Overview

**Nova Horizon** (`nova-hrzn`) is a **local-first productivity CLI** written in Go.
It provides structured workflow tooling with **optional AI assistance** via Google’s Gemini API.

The CLI is designed to do real work on the user’s machine using explicit tools and permissions. AI is used selectively to assist with generation, explanation, and acceleration, not as the primary execution engine.

---

## Current Project Structure

The project follows a standard Go CLI layout with a growing separation between **core engine**, **tools**, and **AI integration**:

* **`cmd/`**

  * CLI entry point and command definitions
  * Handles flags, working directory selection, verbosity, and interactive shell
  * Will evolve into the stable interface layer

* **`internal/`**

  * **`core/`** *(to be introduced)*

    * Core engine logic not tied to AI
    * Command execution, permissions, context handling
  * **`agent/`**

    * AI-assisted workflow layer
    * Converts user intent into tool usage when AI is enabled
  * **`gemini/`**

    * Google Gemini API client
    * Will become one of multiple optional AI backends
  * **`config/`**

    * Configuration loading and validation
    * Local-only, no cloud dependency
  * **`tools/`**

    * Explicit tooling primitives
    * File I/O, execution, inspection, formatting, etc.

* **`main.go`**

  * Minimal entry point calling the CLI root

* **`Makefile`**

  * Build, install, clean, and dev helpers

---

## Current Capabilities

### 1. Command Execution

* Run one-off tasks:

  ```bash
  nova-hrzn "Refactor main.go"
  ```
* Optional AI assistance for interpretation and generation

### 2. Interactive Shell

* Session-based CLI shell:

  ```bash
  nova-hrzn
  ```
* Maintains short-lived session context
* Allows conversational task execution

### 3. Safety Model

* Restricted execution extensions (`.go`, `.py`, `.js`, `.ts`, `.sh`)
* Working directory sandbox
* Explicit `--allow-run` flag for execution
* No background execution or hidden state

### 4. Tooling Primitives

* File system tools (read, write, list)
* Command execution with captured output
* Basic validation and error reporting

---

## Immediate Next Steps (Core Completion Phase)

This phase focuses on **finishing the current CLI properly** before expanding scope.

### 1. Decouple Core Logic from AI

**Goal:** Make the CLI fully useful without AI.

Actions:

* Introduce `internal/core/`
* Move:

  * File operations
  * Execution logic
  * Permission checks
  * Context handling
* Ensure every tool can be called directly without the agent loop

Result:

* AI becomes an optional helper, not a dependency
* Easier testing and future expansion

---

### 2. Normalize Tool Interfaces

**Goal:** Treat tools as first-class, inspectable operations.

Actions:

* Define a common `Tool` interface:

  * Name
  * Description
  * Required permissions
  * Dry-run support
* Add structured input/output for tools
* Enforce consistent error handling

Result:

* Predictable behavior
* Foundation for future modules
* Better UX for both humans and AI

---

### 3. Improve Permission & Preview Flow

**Goal:** No surprises. Ever.

Actions:

* Add global dry-run support (not just flags)
* Show:

  * Files to be modified
  * Commands to be executed
* Require confirmation for destructive actions

Result:

* Safer workflows
* Increased user trust
* Easier debugging

---

### 4. Strengthen Interactive Shell UX

**Goal:** Make the shell pleasant, not clever.

Actions:

* Command history
* Tab completion (basic)
* Slash commands:

  * `/reset`
  * `/help`
  * `/config`
  * `/exit`
* Clear separation between:

  * Chat input
  * Tool execution output
  * AI responses

Result:

* Reduced friction
* No “chatbot confusion”

---

### 5. Add Core Workflow Utilities (Free Users)

These are **non-AI**, high-value tools:

* Project inspection

  * Tree view
  * File summaries
* Code hygiene

  * Formatters
  * Linters (if available locally)
* Diff previews
* Template-based file generation

AI may assist in *suggesting* actions, but tools execute them.

---

## Preparation for Future Integration (Without Implementing It Yet)

### 1. Introduce a Module Boundary (No Plugins Yet)

**Do not implement modules yet. Prepare for them.**

Actions:

* Define a `Module` concept in docs only:

  * Manifest format
  * Capability declaration
* Ensure tools are reusable and composable
* Avoid hard-coding “developer-only” logic

Result:

* Clean transition later
* No architectural dead ends

---

### 2. Backend-Aware, Not Backend-Dependent

**Goal:** Allow cloud integration later without refactors.

Actions:

* Abstract:

  * Auth status
  * Feature availability checks
* Stub backend calls:

  * `IsAuthenticated()`
  * `IsFeatureAllowed()`

Result:

* CLI remains local-first
* Cloud becomes additive, not invasive

---

### 3. Configuration Hardening

Actions:

* Validate config on startup
* Clear error messages for missing or invalid values
* Support multiple profiles (local only)

Result:

* Predictable startup
* Less user confusion

---

## How to Work With the Project (Updated)

* **Build**

  ```bash
  make build
  ```

* **Install**

  ```bash
  make install
  ```

* **Develop**

  * Core logic lives in `internal/core`
  * Tools live in `internal/tools`
  * AI logic stays isolated in `internal/agent`

* **Extend**

  * Add tools before adding AI prompts
  * Treat AI as optional glue, not structure

---

## Direction Summary

The current focus is **not**:

* SaaS
* Agents
* Autonomous execution
* Cloud-first features

The focus **is**:

* A reliable local CLI
* Explicit tooling
* Safe execution
* Optional AI acceleration

Finish the core.
Make it boring and correct.
Then build on it.
