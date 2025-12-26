## API Key Setup Explanation

### What Does "Environment Variable" Mean?

An **environment variable** is a temporary value stored in your terminal session. Think of it like a sticky note you put on your computer that only exists while that terminal window is open. Once you close the terminal, the sticky note disappears.

**When you run:**
```bash
export GEMINI_API_KEY="your-api-key-here"
nova-horizon "Your prompt"
```

You're saying: "For this terminal session only, remember this API key. The Nova Horizon application can use it."

### What Does "Config File" Mean?

A **config file** is a permanent storage location on your computer where you write settings. It's like writing something in a notebook that stays on your desk even after you leave. Your application checks this notebook every time it runs to read your settings.

When you create `~/.config/nova-horizon/config.toml`, you're creating a permanent settings file that Nova Horizon will automatically read every time you use it.

---

## Fedora KDE Detailed Setup Guide

### Method 1: Temporary Setup (Testing)

Use this when you want to quickly test the CLI without permanent changes.

#### Step 1: Get Your API Key

1. Open Firefox (or your browser)
2. Visit [https://aistudio.google.com](https://aistudio.google.com)
3. Click the **"Get API key"** button (top-left)
4. Click **"Create API key in new project"**
5. Google will create a new project and show your API key
6. **Click the copy button** next to your API key (it looks like two overlapping squares)


#### Step 2: Open Terminal

1. Right-click on your desktop → **Open Terminal Here** (or press `Ctrl+Alt+T`)
2. A terminal window opens
3. This is where you'll run the commands below


#### Step 3: Set the Environment Variable

Paste this command (replace `your-api-key-here` with your actual key):

```shellscript
export GEMINI_API_KEY="your-api-key-here"
```

**Example with a real key:**

```shellscript
export GEMINI_API_KEY="AIzaSyD8kL9p0Q-xZ5mN3bV7wE8-uH6jK4lO9pR"
```

Press Enter. **Nothing will show up** - that's normal! Your key is now stored in this terminal session.

#### Step 4: Test It Works

```shellscript
nova-horizon "List all files in this directory"
```

If it works, great! If you close this terminal, the key is gone. You'll need to run the `export` command again next time.

---

### Method 2: Permanent Setup (Recommended for Regular Use)

Use this when you want Nova Horizon to always have access to your API key.

#### Step 1: Get Your API Key (same as above)

Visit [https://aistudio.google.com](https://aistudio.google.com) and get your key.

#### Step 2: Create the Config Directory

Open Terminal and run:

```shellscript
mkdir -p ~/.config/nova-horizon
```

This creates a hidden folder in your home directory called `.config/nova-horizon`. The dot (`.`) makes it hidden, so it won't show up in your file manager by default.

#### Step 3: Create the Config File

Run this command to create the configuration file:

```shellscript
nano ~/.config/nova-horizon/config.toml
```

This opens a text editor called `nano`.

#### Step 4: Add Your API Key

Type (or paste) this exactly:

```plaintext
api_key = "your-api-key-here"
```

**Replace** `your-api-key-here` with your actual API key.

**Example:**

```plaintext
api_key = "AIzaSyD8kL9p0Q-xZ5mN3bV7wE8-uH6jK4lO9pR"
```

#### Step 5: Save and Exit

1. Press `Ctrl + X` (the editor asks if you want to save)
2. Press `Y` (for "yes")
3. Press `Enter` (to confirm the filename)


You're done! The file is saved.

#### Step 6: Secure Your Config File

This is **important** because your API key is sensitive:

```shellscript
chmod 600 ~/.config/nova-horizon/config.toml
```

This makes the file readable only by you, not by other users on the system.

#### Step 7: Test It Works

Now open a **new terminal** and run:

```shellscript
nova-horizon "Create a hello world program"
```

It should work! Your API key is now permanently configured.

---

## Verification & Troubleshooting

### Check If Config File Exists

```shellscript
ls -la ~/.config/nova-horizon/config.toml
```

If you see the file listed, it's there. Good!

### Check If File Is Readable

```shellscript
cat ~/.config/nova-horizon/config.toml
```

If you see your API key displayed, the file is readable by you. Good!

### Check Current API Key Setting

```shellscript
echo $GEMINI_API_KEY
```

- If it shows your key: the environment variable is currently set
- If it shows nothing: the environment variable isn't set (but your config file might still work)


### Remove/Reset Your API Key

**If you want to delete the permanent config:**

```shellscript
rm ~/.config/nova-horizon/config.toml
```

**If you want to clear the temporary environment variable:**

```shellscript
unset GEMINI_API_KEY
```

### If Nova Horizon Can't Find Your Config

Make sure:

1. The folder exists: `ls ~/.config/nova-horizon/`
2. The file exists: `ls ~/.config/nova-horizon/config.toml`
3. File is readable: `cat ~/.config/nova-horizon/config.toml` should show your key
4. File has correct name: it must be `config.toml` (exactly that name)


---

## Quick Reference

| Task | Command
|-----|-----
| Set key temporarily | `export GEMINI_API_KEY="your-key"`
| Create config folder | `mkdir -p ~/.config/nova-horizon`
| Create config file | `nano ~/.config/nova-horizon/config.toml`
| Secure config file | `chmod 600 ~/.config/nova-horizon/config.toml`
| View your key (temp) | `echo $GEMINI_API_KEY`
| View your key (permanent) | `cat ~/.config/nova-horizon/config.toml`
| Delete config file | `rm ~/.config/nova-horizon/config.toml`
