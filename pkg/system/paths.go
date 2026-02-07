package system

import (
	"os"
	"path/filepath"
	"strings"
)

// GetPATH returns the directories in the system's PATH environment variable
func GetPATH() []string {
	pathEnv := os.Getenv("PATH")
	if pathEnv == "" {
		return []string{}
	}
	return strings.Split(pathEnv, string(os.PathListSeparator))
}

// GetHomeDir returns the user's home directory
func GetHomeDir() string {
	home, _ := os.UserHomeDir()
	return home
}

// ExpandPath expands ~ to the home directory in a path
func ExpandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		return filepath.Join(GetHomeDir(), path[2:])
	}
	return path
}

// GetCommonBinaryPaths returns common directories where binaries are installed
func GetCommonBinaryPaths() []string {
	platform := GetPlatform()
	paths := []string{
		"/usr/local/bin",
	}

	// Add macOS specific paths
	if platform.IsDarwin() {
		paths = append(paths,
			"/opt/homebrew/bin", // Apple Silicon Homebrew
		)
	}

	// Add user-specific paths
	homeDir := GetHomeDir()
	if homeDir != "" {
		paths = append(paths,
			filepath.Join(homeDir, ".local/bin"),
			filepath.Join(homeDir, "bin"),
		)
	}

	return paths
}
