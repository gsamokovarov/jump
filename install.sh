#!/usr/bin/env bash
set -euo pipefail

# Jump installer — fetches the latest release binary for the current user.
#
#   curl -fsSL https://raw.githubusercontent.com/gsamokovarov/jump/main/install.sh | bash
#
# Environment overrides:
#   INSTALL_DIR   where to put the binary (default: ~/.local/bin)
#   VERSION       a release tag to pin, e.g. v0.67.0 (default: latest)

REPO="gsamokovarov/jump"
INSTALL_DIR="${INSTALL_DIR:-${HOME}/.local/bin}"
VERSION="${VERSION:-latest}"
BINARY_NAME="jump"

OS="$(uname -s)"
case "$OS" in
  Darwin)               OS="darwin" ;;
  Linux)                OS="linux" ;;
  MINGW*|MSYS*|CYGWIN*) OS="windows" ;;
  *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

ARCH="$(uname -m)"
case "$ARCH" in
  x86_64|amd64)  ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

SUFFIX=""
case "${OS}-${ARCH}" in
  linux-amd64)   ASSET="jump_linux_amd64_binary" ;;
  linux-arm64)   ASSET="jump_linux_arm_binary" ;;
  darwin-amd64)  ASSET="jump_osx" ;;
  darwin-arm64)  ASSET="jump_osx_arm64" ;;
  windows-amd64) ASSET="jump_windows_amd64_binary.exe"; SUFFIX=".exe" ;;
  *)
    echo "No prebuilt jump binary for ${OS}/${ARCH}."
    echo "Install from source instead: go install github.com/${REPO}@latest"
    exit 1
    ;;
esac

if [ "$VERSION" = "latest" ]; then
  URL="https://github.com/${REPO}/releases/latest/download/${ASSET}"
else
  URL="https://github.com/${REPO}/releases/download/${VERSION}/${ASSET}"
fi

echo "Downloading ${BINARY_NAME} (${VERSION}) for ${OS}/${ARCH}..."
TMPDIR="$(mktemp -d)"
trap 'rm -rf "$TMPDIR"' EXIT

if ! curl -fsSL -o "${TMPDIR}/${BINARY_NAME}${SUFFIX}" "$URL"; then
  echo "Error: failed to download ${ASSET} from ${URL}"
  echo "This release may not ship a ${OS}/${ARCH} binary. See:"
  echo "  https://github.com/${REPO}/releases"
  exit 1
fi
chmod +x "${TMPDIR}/${BINARY_NAME}${SUFFIX}"

mkdir -p "$INSTALL_DIR"
install -m 755 "${TMPDIR}/${BINARY_NAME}${SUFFIX}" "${INSTALL_DIR}/${BINARY_NAME}${SUFFIX}"

INSTALLED_VERSION="$("${INSTALL_DIR}/${BINARY_NAME}${SUFFIX}" --version 2>/dev/null || echo "?")"
echo "Installed ${BINARY_NAME} ${INSTALLED_VERSION} to ${INSTALL_DIR}/${BINARY_NAME}${SUFFIX}"

if ! echo "$PATH" | tr ':' '\n' | grep -qx "$INSTALL_DIR"; then
  echo ""
  echo "Add ${INSTALL_DIR} to your PATH:"
  echo "  export PATH=\"${INSTALL_DIR}:\$PATH\""
fi

echo ""
echo "Then enable shell integration (adds the 'j' helper):"
echo "  bash/zsh:    eval \"\$(jump shell)\""
echo "  fish:        jump shell fish | source"
echo "  PowerShell:  Invoke-Expression (&jump shell pwsh | Out-String)"
