package managers

import (
	"context"
	"os"
	"path/filepath"

	"github.com/alexcloudstar/snappoint/internal/scanner"
	"github.com/alexcloudstar/snappoint/pkg/system"
)

// Manual implements the PackageManager interface for detecting ghost binaries
type Manual struct {
	executor      system.CommandExecutor
	knownBinaries map[string]bool
}

// NewManual creates a new Manual package manager
func NewManual(executor system.CommandExecutor) *Manual {
	return &Manual{
		executor:      executor,
		knownBinaries: make(map[string]bool),
	}
}

// Name returns the name of the package manager
func (m *Manual) Name() string {
	return "manual"
}

// IsAvailable always returns true as manual scanning is always available
func (m *Manual) IsAvailable(ctx context.Context) bool {
	return true
}

// SetKnownBinaries sets the list of binaries that are already managed by other package managers
func (m *Manual) SetKnownBinaries(binaries []*scanner.Binary) {
	for _, binary := range binaries {
		key := filepath.Base(binary.Path)
		m.knownBinaries[key] = true
	}
}

// Scan discovers binaries not managed by any package manager
func (m *Manual) Scan(ctx context.Context) ([]*scanner.Binary, error) {
	var binaries []*scanner.Binary

	// Scan common binary directories
	dirsToScan := system.GetCommonBinaryPaths()

	for _, dir := range dirsToScan {
		bins, err := m.scanDirectory(dir)
		if err != nil {
			// Skip directories that can't be scanned
			continue
		}
		binaries = append(binaries, bins...)
	}

	return binaries, nil
}

// scanDirectory scans a single directory for binaries
func (m *Manual) scanDirectory(dir string) ([]*scanner.Binary, error) {
	var binaries []*scanner.Binary

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		// Skip directories and non-executable files
		if entry.IsDir() {
			continue
		}

		// Check if this binary is already managed by another package manager
		if m.knownBinaries[entry.Name()] {
			continue
		}

		fullPath := filepath.Join(dir, entry.Name())

		// Check if file is executable
		info, err := entry.Info()
		if err != nil {
			continue
		}

		// Check if the file has execute permissions
		if info.Mode()&0111 == 0 {
			continue
		}

		binaries = append(binaries, &scanner.Binary{
			Name:    entry.Name(),
			Path:    fullPath,
			Manager: m.Name(),
			Version: "unknown",
			Package: "",
		})
	}

	return binaries, nil
}
