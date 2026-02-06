package system

import (
	"runtime"
)

// Platform represents the current system's operating system and architecture
type Platform struct {
	OS   string
	Arch string
}

// GetPlatform returns the current platform information
func GetPlatform() Platform {
	return Platform{
		OS:   runtime.GOOS,
		Arch: runtime.GOARCH,
	}
}

// String returns a human-readable platform identifier (e.g., "darwin-arm64")
func (p Platform) String() string {
	return p.OS + "-" + p.Arch
}

// IsDarwin returns true if running on macOS
func (p Platform) IsDarwin() bool {
	return p.OS == "darwin"
}

// IsLinux returns true if running on Linux
func (p Platform) IsLinux() bool {
	return p.OS == "linux"
}

// IsARM returns true if running on ARM architecture
func (p Platform) IsARM() bool {
	return p.Arch == "arm64" || p.Arch == "arm"
}
