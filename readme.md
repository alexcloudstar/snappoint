# SnapPoint 🛠️

[![Twitter Follow](https://img.shields.io/twitter/follow/alexcloudstar?style=social)](https://x.com/alexcloudstar)
[![GitHub stars](https://img.shields.io/github/stars/alexcloudstar/snappoint?style=social)](https://github.com/alexcloudstar/snappoint/stargazers)
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

* [ ] **Alpha (Hunt):** Basic discovery engine to list all binaries and their managers.
* [ ] **TUI Dashboard:** A beautiful terminal interface (built with Go/Bubble Tea) to toggle system health.
* [ ] **The "Purge":** Intelligent uninstaller that follows "Cleanup Recipes" for popular dev tools.
* [ ] **Snap Profiles (Social):** * [ ] Export your setup to a `snap.json` or `Snapfile`.
    * [ ] **Community Gallery:** Browse and adopt setups from top devs (e.g., "The Ultimate Go Dev Profile" or "Alex's MacOS Essentials").
* [ ] **Sync:** Recreate your environment on a new Mac/Linux box with one command.

---

## 🛠️ Installation (Coming Soon)

```bash
# This is the goal:
curl -sS https://snappoint.dev/install.sh | sh

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
