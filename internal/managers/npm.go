package managers

import (
	"context"
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/alexcloudstar/snappoint/internal/scanner"
	"github.com/alexcloudstar/snappoint/pkg/system"
)

// NPM implements the PackageManager interface for NPM
type NPM struct {
	executor system.CommandExecutor
}

// NewNPM creates a new NPM package manager
func NewNPM(executor system.CommandExecutor) *NPM {
	return &NPM{
		executor: executor,
	}
}

// Name returns the name of the package manager
func (n *NPM) Name() string {
	return "npm"
}

// IsAvailable checks if NPM is installed
func (n *NPM) IsAvailable(ctx context.Context) bool {
	return n.executor.IsAvailable(ctx, "npm")
}

// Scan discovers all globally installed NPM packages
func (n *NPM) Scan(ctx context.Context) ([]*scanner.Binary, error) {
	// Get global packages list
	output, err := n.executor.Execute(ctx, "npm", "list", "-g", "--depth=0", "--json")
	if err != nil {
		// npm list can return non-zero exit codes even on success if there are issues
		// Try to parse the output anyway
		if output == "" {
			return nil, err
		}
	}

	var data struct {
		Dependencies map[string]struct {
			Version string `json:"version"`
		} `json:"dependencies"`
	}

	if err := json.Unmarshal([]byte(output), &data); err != nil {
		return nil, err
	}

	// Get global bin directory
	binDir, err := n.executor.Execute(ctx, "npm", "bin", "-g")
	if err != nil {
		return nil, err
	}
	binDir = strings.TrimSpace(binDir)

	var binaries []*scanner.Binary

	for name, pkg := range data.Dependencies {
		binaryPath := filepath.Join(binDir, name)

		binaries = append(binaries, &scanner.Binary{
			Name:    name,
			Path:    binaryPath,
			Manager: n.Name(),
			Version: pkg.Version,
			Package: name,
		})
	}

	return binaries, nil
}
