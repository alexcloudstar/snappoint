package system

import (
	"os"
)

// FileValidator provides methods for validating file system properties
type FileValidator struct{}

// NewFileValidator creates a new FileValidator instance
func NewFileValidator() *FileValidator {
	return &FileValidator{}
}

// IsBinaryExecutable checks if a file exists and is executable
// Returns true only if:
// - The file exists
// - It's a regular file (not a directory)
// - It has execute permissions
// - If it's a symlink, the target is validated
func (v *FileValidator) IsBinaryExecutable(path string) bool {
	// Get file info (follows symlinks)
	info, err := os.Stat(path)
	if err != nil {
		// File doesn't exist, permission denied, or broken symlink
		return false
	}

	// Must be a regular file, not a directory
	if !info.Mode().IsRegular() {
		return false
	}

	// Check if the file has execute permissions (any of user/group/other)
	if info.Mode()&0111 == 0 {
		return false
	}

	return true
}
