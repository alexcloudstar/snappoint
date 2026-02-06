package scanner

import (
	"fmt"
	"strings"
)

// Binary represents a binary executable found on the system
type Binary struct {
	Name          string
	Path          string
	Manager       string // "homebrew", "npm", "pip", "manual"
	Version       string
	Package       string
	ConflictsWith []*Binary
}

// IsGhost returns true if the binary is not managed by any package manager
func (b *Binary) IsGhost() bool {
	return b.Manager == "manual" || b.Manager == "ghost"
}

// HasConflicts returns true if this binary conflicts with other versions
func (b *Binary) HasConflicts() bool {
	return len(b.ConflictsWith) > 0
}

// String returns a human-readable representation of the binary
func (b *Binary) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s (%s)", b.Name, b.Path))
	if b.Version != "" {
		sb.WriteString(fmt.Sprintf(" v%s", b.Version))
	}
	sb.WriteString(fmt.Sprintf(" [%s]", b.Manager))
	return sb.String()
}

// ScanResult holds the results of a system scan
type ScanResult struct {
	Binaries  []*Binary
	Conflicts map[string][]*Binary // Map of binary name to conflicting versions
	Ghosts    []*Binary
}

// NewScanResult creates a new scan result with initialized fields
func NewScanResult() *ScanResult {
	return &ScanResult{
		Binaries:  make([]*Binary, 0),
		Conflicts: make(map[string][]*Binary),
		Ghosts:    make([]*Binary, 0),
	}
}

// AddBinary adds a binary to the scan result
func (sr *ScanResult) AddBinary(binary *Binary) {
	sr.Binaries = append(sr.Binaries, binary)

	if binary.IsGhost() {
		sr.Ghosts = append(sr.Ghosts, binary)
	}
}

// DetectConflicts finds binaries with the same name but different versions or paths
func (sr *ScanResult) DetectConflicts() {
	nameMap := make(map[string][]*Binary)

	// Group binaries by name
	for _, binary := range sr.Binaries {
		nameMap[binary.Name] = append(nameMap[binary.Name], binary)
	}

	// Find conflicts (multiple versions of the same binary)
	for name, binaries := range nameMap {
		if len(binaries) > 1 {
			sr.Conflicts[name] = binaries

			// Mark each binary as conflicting with the others
			for i, b1 := range binaries {
				for j, b2 := range binaries {
					if i != j {
						b1.ConflictsWith = append(b1.ConflictsWith, b2)
					}
				}
			}
		}
	}
}

// TotalCount returns the total number of binaries found
func (sr *ScanResult) TotalCount() int {
	return len(sr.Binaries)
}

// ConflictCount returns the number of binaries with conflicts
func (sr *ScanResult) ConflictCount() int {
	return len(sr.Conflicts)
}

// GhostCount returns the number of ghost binaries
func (sr *ScanResult) GhostCount() int {
	return len(sr.Ghosts)
}
