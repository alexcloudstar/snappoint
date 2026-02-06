#!/bin/bash
set -e

# SnapPoint CLI Installation Script
# Usage: curl -sS https://snappoint.dev/install.sh | sh

REPO="alexcloudstar/snappoint"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="snappoint"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Logging functions
info() {
    printf "${GREEN}[INFO]${NC} %s\n" "$1"
}

warn() {
    printf "${YELLOW}[WARN]${NC} %s\n" "$1"
}

error() {
    printf "${RED}[ERROR]${NC} %s\n" "$1"
    exit 1
}

# Detect platform
detect_platform() {
    local os=$(uname -s | tr '[:upper:]' '[:lower:]')
    local arch=$(uname -m)

    case "$os" in
        darwin*)
            os="darwin"
            ;;
        linux*)
            os="linux"
            ;;
        *)
            error "Unsupported operating system: $os"
            ;;
    esac

    case "$arch" in
        x86_64)
            arch="amd64"
            ;;
        aarch64|arm64)
            arch="arm64"
            ;;
        *)
            error "Unsupported architecture: $arch"
            ;;
    esac

    echo "${os}-${arch}"
}

# Get latest release version
get_latest_version() {
    local api_url="https://api.github.com/repos/${REPO}/releases/latest"
    local version=$(curl -s "$api_url" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

    if [ -z "$version" ]; then
        error "Failed to fetch latest version from GitHub"
    fi

    echo "$version"
}

# Download binary
download_binary() {
    local version=$1
    local platform=$2
    local binary_name="${BINARY_NAME}-${platform}"
    local download_url="https://github.com/${REPO}/releases/download/${version}/${binary_name}"
    local tmp_file="/tmp/${BINARY_NAME}"

    info "Downloading SnapPoint ${version} for ${platform}..."

    if ! curl -sL "$download_url" -o "$tmp_file"; then
        error "Failed to download binary from ${download_url}"
    fi

    chmod +x "$tmp_file"
    echo "$tmp_file"
}

# Install binary
install_binary() {
    local tmp_file=$1
    local install_path="${INSTALL_DIR}/${BINARY_NAME}"

    info "Installing to ${install_path}..."

    if [ -w "$INSTALL_DIR" ]; then
        mv "$tmp_file" "$install_path"
    else
        warn "Requires sudo to install to ${INSTALL_DIR}"
        sudo mv "$tmp_file" "$install_path"
        sudo chmod +x "$install_path"
    fi
}

# Verify installation
verify_installation() {
    if ! command -v "$BINARY_NAME" &> /dev/null; then
        error "Installation failed. ${BINARY_NAME} not found in PATH"
    fi

    local version=$($BINARY_NAME --version)
    info "Successfully installed ${version}"
}

# Main installation flow
main() {
    echo "SnapPoint CLI Installer"
    echo "======================="
    echo

    # Detect platform
    local platform=$(detect_platform)
    info "Detected platform: ${platform}"

    # Get latest version
    local version=$(get_latest_version)
    info "Latest version: ${version}"

    # Check if already installed
    if command -v "$BINARY_NAME" &> /dev/null; then
        local current_version=$($BINARY_NAME --version | awk '{print $2}')
        warn "SnapPoint is already installed (${current_version})"
        read -p "Do you want to continue with the installation? (y/N) " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            info "Installation cancelled"
            exit 0
        fi
    fi

    # Download binary
    local tmp_file=$(download_binary "$version" "$platform")

    # Install binary
    install_binary "$tmp_file"

    # Verify installation
    verify_installation

    echo
    info "Installation complete! Run '${BINARY_NAME} --help' to get started."
}

main
