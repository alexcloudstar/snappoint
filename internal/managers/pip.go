package managers

import (
	"context"
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/alexcloudstar/snappoint-cli/internal/scanner"
	"github.com/alexcloudstar/snappoint-cli/pkg/system"
)

// Pip implements the PackageManager interface for Pip
type Pip struct {
	executor system.CommandExecutor
}

// NewPip creates a new Pip package manager
func NewPip(executor system.CommandExecutor) *Pip {
	return &Pip{
		executor: executor,
	}
}

// Name returns the name of the package manager
func (p *Pip) Name() string {
	return "pip"
}

// IsAvailable checks if Pip is installed
func (p *Pip) IsAvailable(ctx context.Context) bool {
	// Try pip3 first, then pip
	if p.executor.IsAvailable(ctx, "pip3") {
		return true
	}
	return p.executor.IsAvailable(ctx, "pip")
}

// Scan discovers all globally installed Pip packages
func (p *Pip) Scan(ctx context.Context) ([]*scanner.Binary, error) {
	// Determine which pip command to use
	pipCmd := "pip3"
	if !p.executor.IsAvailable(ctx, "pip3") {
		pipCmd = "pip"
	}

	// Get list of installed packages
	output, err := p.executor.Execute(ctx, pipCmd, "list", "--format=json")
	if err != nil {
		return nil, err
	}

	var packages []struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}

	if err := json.Unmarshal([]byte(output), &packages); err != nil {
		return nil, err
	}

	var binaries []*scanner.Binary

	// Get information about each package's binaries
	for _, pkg := range packages {
		bins, err := p.getPackageBinaries(ctx, pipCmd, pkg.Name, pkg.Version)
		if err != nil {
			// Skip packages that don't have binaries or fail to get info
			continue
		}
		binaries = append(binaries, bins...)
	}

	return binaries, nil
}

// getPackageBinaries retrieves binaries for a specific pip package
func (p *Pip) getPackageBinaries(ctx context.Context, pipCmd, packageName, version string) ([]*scanner.Binary, error) {
	output, err := p.executor.Execute(ctx, pipCmd, "show", packageName)
	if err != nil {
		return nil, err
	}

	// Parse pip show output to find Location
	lines := strings.Split(output, "\n")
	var location string
	var name string

	for _, line := range lines {
		if strings.HasPrefix(line, "Location:") {
			location = strings.TrimSpace(strings.TrimPrefix(line, "Location:"))
		}
		if strings.HasPrefix(line, "Name:") {
			name = strings.TrimSpace(strings.TrimPrefix(line, "Name:"))
		}
	}

	if location == "" {
		return nil, nil
	}

	// Check if package has a binary in common Python bin directories
	binDirs := []string{
		filepath.Join(filepath.Dir(location), "bin"),
		"/usr/local/bin",
		filepath.Join(system.GetHomeDir(), ".local/bin"),
	}

	var binaries []*scanner.Binary

	// Use the package name as a potential binary name
	binaryName := strings.ToLower(packageName)

	for _, binDir := range binDirs {
		binaryPath := filepath.Join(binDir, binaryName)

		binaries = append(binaries, &scanner.Binary{
			Name:    name,
			Path:    binaryPath,
			Manager: p.Name(),
			Version: version,
			Package: packageName,
		})
		break // Only add one binary per package
	}

	return binaries, nil
}
