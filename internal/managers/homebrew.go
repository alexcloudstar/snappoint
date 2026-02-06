package managers

import (
	"context"
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/alexcloudstar/snappoint/internal/scanner"
	"github.com/alexcloudstar/snappoint/pkg/system"
)

// Homebrew implements the PackageManager interface for Homebrew
type Homebrew struct {
	executor system.CommandExecutor
}

// NewHomebrew creates a new Homebrew package manager
func NewHomebrew(executor system.CommandExecutor) *Homebrew {
	return &Homebrew{
		executor: executor,
	}
}

// Name returns the name of the package manager
func (h *Homebrew) Name() string {
	return "homebrew"
}

// IsAvailable checks if Homebrew is installed
func (h *Homebrew) IsAvailable(ctx context.Context) bool {
	return h.executor.IsAvailable(ctx, "brew")
}

// Scan discovers all binaries managed by Homebrew
func (h *Homebrew) Scan(ctx context.Context) ([]*scanner.Binary, error) {
	var binaries []*scanner.Binary

	// Get list of installed packages
	output, err := h.executor.Execute(ctx, "brew", "list", "--formula")
	if err != nil {
		return nil, err
	}

	packages := strings.Split(strings.TrimSpace(output), "\n")

	// Get info for each package
	for _, pkg := range packages {
		if pkg == "" {
			continue
		}

		info, err := h.getPackageInfo(ctx, pkg)
		if err != nil {
			// Skip packages that fail to get info
			continue
		}

		binaries = append(binaries, info...)
	}

	return binaries, nil
}

// getPackageInfo retrieves detailed information about a Homebrew package
func (h *Homebrew) getPackageInfo(ctx context.Context, packageName string) ([]*scanner.Binary, error) {
	output, err := h.executor.Execute(ctx, "brew", "info", "--json=v2", packageName)
	if err != nil {
		return nil, err
	}

	var info struct {
		Formulae []struct {
			Name      string `json:"name"`
			Version   string `json:"version"`
			Installed []struct {
				Version string `json:"version"`
			} `json:"installed"`
			Linked bool `json:"linked"`
		} `json:"formulae"`
	}

	if err := json.Unmarshal([]byte(output), &info); err != nil {
		return nil, err
	}

	var binaries []*scanner.Binary

	for _, formula := range info.Formulae {
		version := formula.Version
		if len(formula.Installed) > 0 {
			version = formula.Installed[0].Version
		}

		// Determine Homebrew prefix based on architecture
		platform := system.GetPlatform()
		prefix := "/usr/local"
		if platform.IsARM() {
			prefix = "/opt/homebrew"
		}

		// Common binary path
		binaryPath := filepath.Join(prefix, "bin", formula.Name)

		binaries = append(binaries, &scanner.Binary{
			Name:    formula.Name,
			Path:    binaryPath,
			Manager: h.Name(),
			Version: version,
			Package: packageName,
		})
	}

	return binaries, nil
}
