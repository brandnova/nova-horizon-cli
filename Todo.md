# Todo

<!-- Immediate Todos - Being Implemented -->

* [ ] Create `internal/core/` and move non-AI execution logic (filesystem, command execution, permissions) out of `agent`
* [ ] Define a `Tool` interface (name, description, permissions, dry-run, execute)
* [ ] Refactor existing tools to implement the new `Tool` interface
* [ ] Make the CLI runnable with tools **without** invoking the agent loop
* [ ] Add global dry-run support and unified preview output
* [ ] Separate AI prompt handling from tool execution paths
* [ ] Add basic interactive shell improvements (history + `/exit`)
* [ ] Write minimal tests for core tools (file read/write, exec safety)
* [ ] Update `info.md` and internal docs to reflect the new core/agent split
