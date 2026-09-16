# openpanel-plugin-boilerplate

A starting point for building your own OpenPanel plugin. Fork this repo, rename it, edit `main.go` and `readme.txt`, push - the build workflow rebuilds the binaries for you.

This example does the minimum needed to prove the pattern works: it adds one page at `/new-plugin`, with a sidebar link, a dashboard icon, and a search entry - all from `readme.txt` alone, no OpenPanel code changes required.

## How it works

OpenPanel execs your plugin's binary directly - it never loads plugin code into its own process. Your plugin is a standalone program (any language would technically work, but the build workflow here targets Go) with subcommands OpenPanel calls over a normal subprocess exec:

- **`plugin page`** - called when a user visits the URL your `readme.txt`'s `link` field names. Print the HTML for that page's content to stdout; OpenPanel wraps it in its own chrome (sidebar, topbar, flashes). See `main.go`'s `servePage()`.

(The [captcha plugin](https://github.com/stefanpejcic/captcha) is a second example, using a different pair of subcommands - `widget`/`verify` - for a login-flow plugin instead of a full page. Subcommands are whatever your plugin needs; OpenPanel only knows about the ones a given integration point calls.)

## readme.txt fields

```
name = example-plugin          # folder/binary name - must match the repo name once installed
title = Example Plugin         # shown in the sidebar, dashboard, and search
description = ...              # shown on the OpenAdmin plugin store card and in search
link = /new-plugin              # the URL your plugin's "page" subcommand serves
version = 1.0.0
author = Your Name
category = advanced            # sidebar group to link under - must be an existing one: account, advanced, cache, docker, domains, emails, files, mongodb, mysql, php, postgresql, websites
icon = bi bi-puzzle             # Bootstrap Icons class, shown on the dashboard
show_in_search = 1              # 1 to show in the sidebar search box, 0 to hide
help_link = https://github.com/you/your-plugin
```

## Fork and edit

1. Fork this repo, rename it to your plugin's name.
2. Edit `readme.txt` - at minimum `name`, `title`, `description`, `link`, `author`, `help_link`.
3. Edit `main.go`'s `servePage()` to render your own page content (plain HTML - Tailwind classes are available, since OpenPanel's own CSS is already loaded).
4. Push to `main` - the build workflow compiles `plugin-linux-amd64` and `plugin-linux-arm64` and commits them back into the repo.
5. Publish your fork's URL to the [OpenPanel plugin store](https://github.com/stefanpejcic/OpenPanel/tree/main/plugins) as a submodule (open a PR there), or just tell people to install it directly by URL.

## Installing

Via OpenAdmin's plugin store (*Settings > Modules > Plugin Store*): paste your fork's git URL, or find it in the store list if it's been added there. The store clones your repo into `/etc/openpanel/modules/<name>/` and runs `install.sh` automatically.

Manually, on the server:
```bash
cd /etc/openpanel/modules/ && git clone https://github.com/you/your-plugin <name>
./<name>/install.sh
docker restart openpanel
```

## Build workflow

`.github/workflows/build.yml` builds both architectures and commits the binaries back to the repo (so installing never needs a compiler on the server). It runs:
- on every push that touches Go source
- on demand (`workflow_dispatch`, the "Run workflow" button on the Actions tab)
- weekly, to catch toolchain/base-image drift even with no code changes

`install.sh` picks whichever of `plugin-linux-amd64` / `plugin-linux-arm64` matches the server's `uname -m` and installs it as `<name>` - the exact filename OpenPanel execs.
