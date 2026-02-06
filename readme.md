# SnapPoint 🛠️

<a href="https://www.producthunt.com/products/snappoint?embed=true&amp;utm_source=badge-featured&amp;utm_medium=badge&amp;utm_campaign=badge-snappoint" target="_blank" rel="noopener noreferrer"><img alt="SnapPoint - Make Your System Snap Back Into Alignment | Product Hunt" width="250" height="54" src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=1075242&amp;theme=light&amp;t=1770411245928"></a>

[![Twitter Follow](https://img.shields.io/twitter/follow/alexcloudstar?style=social)](https://x.com/alexcloudstar)
[![GitHub stars](https://img.shields.io/github/stars/alexcloudstar/snappoint?style=social)](https://github.com/alexcloudstar/snappoint/stargazers)
[![Release](https://img.shields.io/github/v/release/alexcloudstar/snappoint)](https://github.com/alexcloudstar/snappoint/releases/latest)
[![CI](https://github.com/alexcloudstar/snappoint/actions/workflows/ci.yml/badge.svg)](https://github.com/alexcloudstar/snappoint/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexcloudstar/snappoint)](https://goreportcard.com/report/github.com/alexcloudstar/snappoint)
![Go Version](https://img.shields.io/github/go-mod/go-version/alexcloudstar/snappoint)
[![License](https://img.shields.io/github/license/alexcloudstar/snappoint)](LICENSE)
![macOS](https://img.shields.io/badge/macOS-000000?style=flat&logo=apple&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=flat&logo=linux&logoColor=black)

**Make your system snap back into alignment.**

`snappoint` is an open-source, interactive system auditor and package manager "manager." It finds the ghosts of old installs, traces where that stray `python` binary came from, and deep-cleans the debris left behind by Homebrew, NPM, Pip, and manual `curl` installs.

If **LazyVim** is a meta-distro for your editor, **SnapPoint** is the meta-distro for your local machine.

---

## 💻 OS Support

SnapPoint currently supports:
- **macOS** (Intel & Apple Silicon)
- **Linux** (Most distributions)

*Windows support is on the roadmap!* 🚀

---

## ⚡ Why SnapPoint?

Your dev machine is a graveyard of "tutorial remnants." You've got:

* **Ghosts:** Binaries in your `/usr/local/bin` that no package manager claims.
* **Conflicts:** Three versions of Node fighting for your `$PATH`.
* **Bloat:** GBs of cache in `~/.npm` and `~/.cache/pip` for tools you deleted months ago.
* **Orphans:** Packages installed as dependencies that stayed after the parent was removed.

**SnapPoint identifies the mess and snaps it back to a clean state.**

---

## ✨ Key Features

* **🔍 The Hunt:** Scan your system to identify every global binary and its "Origin Story" (Brew, NVM, Pip, or Manual).
* **👻 Ghost Busting:** Find and purge orphaned binaries that aren't managed by any tool.
* **⚖️ Align:** Resolve `$PATH` conflicts. If you have multiple versions of a tool, SnapPoint lets you pick the "Source of Truth."
* **🧹 Deep Clean:** Don't just `uninstall`. SnapPoint wipes the associated `~/.config` and `~/.cache` folders too.
* **🩺 Doctor:** Find broken symlinks and redundant global packages that are already handled by your local project.
* **👯 Community Profiles:** Share your "System DNA." Export your list of curated tools and configs so others can snap their system into the same professional alignment.

---

## 🚀 Roadmap

* [x] **v0.1.0 - Alpha (Hunt):** Basic discovery engine to list all binaries and their managers.
  * [x] Scan Homebrew, NPM, Pip packages
  * [x] Identify ghost binaries
  * [x] Detect PATH conflicts
  * [x] CLI with scan, list, and doctor commands
  * [x] Cross-platform binaries (macOS Intel/ARM, Linux x64/ARM)
  * [x] Installation script
* [ ] **v0.2.0 - TUI Dashboard:** A beautiful terminal interface (built with Go/Bubble Tea) to toggle system health.
* [ ] **v0.3.0 - The "Purge":** Intelligent uninstaller that follows "Cleanup Recipes" for popular dev tools.
* [ ] **v0.4.0 - Snap Profiles (Social):**
    * [ ] Export your setup to a `snap.json` or `Snapfile`.
    * [ ] **Community Gallery:** Browse and adopt setups from top devs (e.g., "The Ultimate Go Dev Profile" or "Alex's MacOS Essentials").
* [ ] **v0.5.0 - Sync:** Recreate your environment on a new Mac/Linux box with one command.
* [ ] **v1.0.0 - Production Ready:**
    * [ ] Support more package managers (Cargo, Go install, RubyGems, APT, YUM, Pacman)
    * [ ] JSON output format
    * [ ] Configuration file support
    * [ ] Cache management

---

## 🛠️ Installation

### Quick Install

```bash
curl -sS https://snappoint.dev/install.sh | sh
```

### Manual Install

Download the latest binary for your platform from the [releases page](https://github.com/alexcloudstar/snappoint/releases):

```bash
# macOS (Apple Silicon)
curl -L https://github.com/alexcloudstar/snappoint/releases/latest/download/snappoint-darwin-arm64 -o snappoint
chmod +x snappoint
sudo mv snappoint /usr/local/bin/

# macOS (Intel)
curl -L https://github.com/alexcloudstar/snappoint/releases/latest/download/snappoint-darwin-amd64 -o snappoint
chmod +x snappoint
sudo mv snappoint /usr/local/bin/

# Linux (x64)
curl -L https://github.com/alexcloudstar/snappoint/releases/latest/download/snappoint-linux-amd64 -o snappoint
chmod +x snappoint
sudo mv snappoint /usr/local/bin/

# Linux (ARM)
curl -L https://github.com/alexcloudstar/snappoint/releases/latest/download/snappoint-linux-arm64 -o snappoint
chmod +x snappoint
sudo mv snappoint /usr/local/bin/
```

### Build from Source

```bash
git clone https://github.com/alexcloudstar/snappoint.git
cd snappoint
make build
sudo make install
```

---

## 📖 Usage

### Check System Health

```bash
snappoint doctor
```

This command checks which package managers are available and shows your system configuration.

### Scan for Binaries

```bash
# Scan all package managers
snappoint scan

# Scan specific package manager
snappoint scan --manager homebrew
snappoint scan --manager npm
snappoint scan --manager pip
```

### List Binaries

```bash
# List all binaries
snappoint list

# Show only ghost binaries
snappoint list --orphans

# Show only conflicts
snappoint list --conflicts
```

### Example Output

```
┌──────────┬────────────────────────┬─────────────┬──────────┐
│ NAME     │ PATH                   │ MANAGER     │ VERSION  │
├──────────┼────────────────────────┼─────────────┼──────────┤
│ node     │ /usr/local/bin/node    │ homebrew    │ v20.11.0 │
│ node     │ ~/.nvm/versions/...    │ 👻 ghost    │ v18.19.0 │
│ python3  │ /usr/local/bin/python3 │ homebrew    │ 3.12.1   │
│ aws      │ /usr/local/bin/aws     │ pip         │ 1.32.50  │
│ random   │ /usr/local/bin/random  │ 👻 ghost    │ unknown  │
└──────────┴────────────────────────┴─────────────┴──────────┘

ℹ Total binaries found: 5

⚠️  Found 1 conflicts:
  • node: 2 versions detected
    - /usr/local/bin/node (homebrew)
    - ~/.nvm/versions/... (manual)

👻 Found 2 ghost binaries:
  • node: No package manager claims this (~/.nvm/versions/...)
  • random: No package manager claims this (/usr/local/bin/random)
```

---

## 🤝 Contributing

This is **Open Source** because no one person knows where every single dev tool hides its trash. We need your "Cleanup Recipes" and "System Profiles"!

1. **The Goal:** Does this tool solve a pain you have? Does it make system setup easier?
2. **Fork & Code:** We're building this in **Go** for that sweet TUI performance and single-binary portability.
3. **Submit:** Help us map out more package managers, "Ghost" locations, or share your own `Snapfile`.

---

## ⭐ Stargazers

Thank you to all our stargazers! [View all →](STARGAZERS.md)

[![Stargazers](https://reporoster.com/stars/notext/alexcloudstar/snappoint)](https://github.com/alexcloudstar/snappoint/stargazers)

---

## License

Built with ☕ and frustration by [@alexcloudstar](https://x.com/alexcloudstar)

---

test
