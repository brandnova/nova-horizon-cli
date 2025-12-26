Here’s how you do it once your Go project is fully fixed, builds, and runs locally. I’ll break it down cleanly for **Fedora/Linux** and **Windows**.

---

## **1. Linux / Fedora**

### Step 1: Build the static binary

From the project root:

```bash
CGO_ENABLED=0 go build -o nova-horizon .
```

This produces the `nova-horizon` binary in the current directory.

---

### Step 2: Copy the binary to a directory in your PATH

Common choices:

* User-local: `~/.local/bin`
* System-wide: `/usr/local/bin`

For user-local:

```bash
mkdir -p ~/.local/bin
cp nova-horizon ~/.local/bin/
```

Then make sure `~/.local/bin` is in your PATH:

```bash
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

Check:

```bash
which nova-horizon
```

It should print:

```
/home/yourusername/.local/bin/nova-horizon
```

Now you can run it from **any directory**:

```bash
nova-horizon "Say hello"
```

---

### Step 3: Optional – make it system-wide

If you want all users to have access:

```bash
sudo cp nova-horizon /usr/local/bin/
sudo chmod +x /usr/local/bin/nova-horizon
```

---

## **2. Windows**

Windows works differently, but the principle is the same.

### Step 1: Build the binary for Windows

Open PowerShell in the project directory:

```powershell
setx CGO_ENABLED 0
go build -o nova-horizon.exe
```

This produces `nova-horizon.exe`.

---

### Step 2: Add a folder to your PATH

* Create a folder like `C:\GoApps\` (or reuse an existing bin folder)
* Move `nova-horizon.exe` there
* Add that folder to your PATH:

1. Press **Win + S** → search `Environment Variables`
2. Click **Edit the system environment variables**
3. Under **System Variables**, select **Path** → **Edit** → **New** → add `C:\GoApps\`
4. Open a **new terminal** to pick up the change

---

### Step 3: Test it

```powershell
nova-horizon.exe "Say hello"
```

Or just:

```powershell
nova-horizon "Say hello"
```

It should run from any folder.

---

### Notes

* Linux and Windows can both use **static binaries**; no runtime Go install is required once built.
* If you plan to distribute your CLI, **commit `go.mod` and `go.sum`**, build for each OS/arch, and ship binaries.
* Always verify that your Gemini API key is set (env var or config file) before running.

---

If you want, I can give a **ready-to-go one-line installer command for Fedora** that copies the binary to `~/.local/bin`, sets the PATH if missing, and makes it immediately available. It’s nice for CI or other systems.
