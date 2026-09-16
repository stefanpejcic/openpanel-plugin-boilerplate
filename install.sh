#!/bin/bash
# Installs the correct prebuilt binary for this server's CPU architecture.
#
# Run automatically by OpenAdmin's plugin store right after cloning this
# repo into /etc/openpanel/modules/<name>/. Run it yourself if you installed
# this plugin manually on the CLI:
#   cd /etc/openpanel/modules/ && git clone <your-fork-url> <name> && ./<name>/install.sh
set -e

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
NAME="$(basename "$DIR")"

case "$(uname -m)" in
    x86_64|amd64)  ARCH=amd64 ;;
    aarch64|arm64) ARCH=arm64 ;;
    *) echo "Unsupported architecture: $(uname -m)" >&2; exit 1 ;;
esac

SRC="$DIR/plugin-linux-$ARCH"
if [ ! -f "$SRC" ]; then
    echo "Missing prebuilt binary: $SRC (did the build workflow run?)" >&2
    exit 1
fi

cp "$SRC" "$DIR/$NAME"
chmod +x "$DIR/$NAME"
echo "Installed '$NAME' ($ARCH). Restart OpenPanel to activate: docker restart openpanel"
