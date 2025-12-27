# Nova Horizon: Vision & Roadmap

## 🌟 The Goal

**Nova Horizon** (`nova-hrzn`) is a **local-first CLI tooling system** designed to make common technical and creative workflows faster, safer, and more repeatable.

The core philosophy is simple:

- The **CLI does the work**.
- **Tools come first**.
- **AI assists, it does not orchestrate**.
- The cloud exists to support, not control.

Nova Horizon is not trying to be a magical black box. It is a practical command-line assistant that helps users _do real work_ using reliable local tools, with optional AI guidance where it actually adds value.

**If the internet disappears, the CLI still works. If AI is disabled, the tooling still works.**

---

## 🏗️ Core Design Principles

### 1. Local-First by Default

All meaningful execution happens on the user’s machine. The CLI interacts directly with the filesystem, shell tools, and runtimes the user already has installed.

### 2. Tooling Over Intelligence

Every feature must stand on its own **without AI**. AI is allowed only when it clearly reduces friction (e.g., generating a complex script), not when it replaces structure.

### 3. Explicit Permission Model

Nothing runs unless the user can see what will happen and approve it.

### 4. Cloud as Infrastructure, Not Control

The backend exists for authentication, module discovery, and optional premium services. It never dictates local execution logic.

---

## 🧩 System Architecture

### 1. The CLI Client (Go)

- **Core Engine**: Stateless, fast, deterministic. Handles authentication, and permission validation.
- **Module Runtime**: Executes modules as **subprocesses**, not unstable internal plugins. This ensures isolation and stability across OSs.
- **Local Agent**: Runs locally for privacy. AI calls are scoped, optional, and visible.

### 2. Module System (Declarative)

Modules are **described workflows**, not just loaded code. A module consists of:

- `module.toml` manifest.
- Scripts/Templates (Python, Bash, Node, etc.).
- Declared capabilities (filesystem, network).

This design allows modules to be inspectable and prevents third-party code from crashing the CLI core.

### 3. The Backend API (Django + Python)

- **User Management**: JWT Authentication.
- **Module Registry**: A "Package Manager" for validated Nova modules.
- **Subscription System**: Stripe integration for Free vs. Premium tiers.
- **LLM Gateway**: Proxies requests for Premium users to high-tier models.

### 4. The Web Dashboard (React)

- User portal for subscription management.
- Marketplace for discovering new modules.
- Configuration sync.

---

## 🗺️ Detailed Roadmap

### Phase 1: Foundation (Current Status) ✅

_Goal: A solid, standalone CLI tool for developers._

- [x] Basic Go CLI structure & Gemini integration (BYOK).
- [x] Core Coding Tools (Read/Write Code, Execute Scripts).
- [x] Interactive Shell & ASCII Art.
- [x] Safety mechanisms (Sandbox, Allowlist).

### Phase 2: User Experience & Utility 🚧

_Goal: Polish the Core CLI with high-value free features._

- [ ] **AI-Driven Bootstrapping**: `nova-hrzn --bootstrap prompt`. Automatically generate platform-specific scaffolding (e.g., "Create a Python project with virtualenv and Streamlit"). The AI generates the script, explains it, and runs it upon user approval.
- [ ] **Utility Tools**: Integrated, offline Text-to-ASCII and Image-to-ASCII converters.
- [ ] **Enhanced Shell**: Better history, autocomplete, and slash commands (`/reset`, `/config`, `/exit`).
- [ ] **Context Management**: Ability to "pin" files to context so the AI always remembers them.

### Phase 3: The Modular Engine 🧩

_Goal: Re-architect the CLI to support downloadable modules._

- [ ] **Module Standard**: Define the `module.toml` manifest specification.
- [ ] **Core Modules**: Refactor existing coding features into the "Developer Module" while keeping the Core Engine light.
- [ ] **New Modules**:
  - **Writer Module**: Web scraping, summarization, SEO optimization.
  - **Designer Module**: Asset generation, SVG optimization.
  - **Data Analyst Module**: SQL query generation, local data visualization.

### Phase 4: The Cloud Backend (Django) ☁️

_Goal: Connect the local CLI to a centralized server._

- [ ] **Auth System**: `nova-hrzn login` command flows.
- [ ] **Module Registry API**: Server-side storage for signed module packages.
- [ ] **Database Design**: Users, Plans, Modules, UsageLogs.

### Phase 5: The SaaS Platform 🚀

_Goal: Monetization and User Management._

- [ ] **React Dashboard**: Web UI for account management.
- [ ] **Stripe Integration**: Handling payments for Premium subscriptions.
- [ ] **Premium LLM Gateway**: Routing premium user requests to higher-tier models without requires user keys.

---

## 💡 Module Concepts

### 👨‍💻 Developer Module (Core/Default)

_Supercharged Coding Assistant_

- **Features**: Refactoring, Unit Test Generation, CI/CD Pipeline Generation, Smart Project Scaffolding.
- **Philosophy**: Code creation is deterministic where possible; AI fills the gaps.

### ✍️ Writer Module

_AI Research & Editor Assistant_

- **Features**:
  - **Web Scraper**: `curl` wrapper + readability parsing.
  - **Format Converter**: `pandoc` wrapper aiming for Markdown outputs.
  - **SEO**: Keyword analysis.

### 🎨 Designer Module

_Creative Asset Assistant_

- **Features**:
  - **Image Ops**: `imagemagick` wrappers for resizing/conversion.
  - **Asset Gen**: Placeholder generation.

### 📊 Data Analyst Module

_Local Data Operations_

- **Features**:
  - **Query**: Natural language to Pandas/SQL.
  - **Plot**: Auto-generation of Matplotlib scripts.
  - **Pipeline**: Data cleaning automation.

---

## 💰 Monetization Strategy

| Feature        | Free Tier (Community)                 | Premium Tier (Pro)                        |
| :------------- | :------------------------------------ | :---------------------------------------- |
| **LLM Access** | Bring Your Own Key (Free/Paid Gemini) | Managed Premium Access (High Rate Limits) |
| **Modules**    | Basic / Community Modules             | Exclusive "Pro" Modules                   |
| **Usage**      | Unlimited (Local)                     | Priority Processing                       |
| **Support**    | Community                             | Priority Email Support                    |
| **Cloud Sync** | No                                    | Yes (Config & Preferences)                |

---

## The Intended Outcome

Nova Horizon aims to become:

- A trusted CLI users keep installed for years.
- A toolkit that grows without becoming fragile.
- A product that earns trust before asking for money.

**No hype. No hidden automation. Just solid software.**
