package managers

import (
	"context"
	"encoding/json"
	"os"
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
	validator := system.NewFileValidator()
	prefixes := []string{"/opt/homebrew", "/usr/local"}

	for _, formula := range info.Formulae {
		version := formula.Version
		if len(formula.Installed) > 0 {
			version = formula.Installed[0].Version
		}

		for _, prefix := range prefixes {
			// Skip prefixes whose Cellar doesn't exist
			cellarDir := filepath.Join(prefix, "Cellar")
			if _, err := os.Stat(cellarDir); err != nil {
				continue
			}

			// Try to find all binaries in the Cellar bin directory
			cellarBinDir := filepath.Join(cellarDir, formula.Name, version, "bin")
			entries, err := os.ReadDir(cellarBinDir)
			if err == nil && len(entries) > 0 {
				for _, entry := range entries {
					if entry.IsDir() {
						continue
					}
					linkedPath := filepath.Join(prefix, "bin", entry.Name())
					if validator.IsBinaryExecutable(linkedPath) {
						binaries = append(binaries, &scanner.Binary{
							Name:    entry.Name(),
							Path:    linkedPath,
							Manager: h.Name(),
							Version: version,
							Package: packageName,
						})
					}
				}
			} else {
				// Fallback for library-only formulae: check prefix/bin/<formula>
				binaryPath := filepath.Join(prefix, "bin", formula.Name)
				if validator.IsBinaryExecutable(binaryPath) {
					binaries = append(binaries, &scanner.Binary{
						Name:    formula.Name,
						Path:    binaryPath,
						Manager: h.Name(),
						Version: version,
						Package: packageName,
					})
				}
			}
		}
	}

	return binaries, nil
}
