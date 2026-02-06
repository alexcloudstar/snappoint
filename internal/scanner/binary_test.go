package scanner

import (
	"testing"
)

func TestBinaryIsGhost(t *testing.T) {
	tests := []struct {
		name     string
		binary   *Binary
		expected bool
	}{
		{
			name: "Manual manager",
			binary: &Binary{
				Name:    "test",
				Manager: "manual",
			},
			expected: true,
		},
		{
			name: "Ghost manager",
			binary: &Binary{
				Name:    "test",
				Manager: "ghost",
			},
			expected: true,
		},
		{
			name: "Homebrew manager",
			binary: &Binary{
				Name:    "test",
				Manager: "homebrew",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.binary.IsGhost() != tt.expected {
				t.Errorf("Expected IsGhost() to be %v, got %v", tt.expected, tt.binary.IsGhost())
			}
		})
	}
}

func TestScanResultAddBinary(t *testing.T) {
	result := NewScanResult()

	binary := &Binary{
		Name:    "test",
		Path:    "/usr/local/bin/test",
		Manager: "homebrew",
		Version: "1.0.0",
	}

	result.AddBinary(binary)

	if result.TotalCount() != 1 {
		t.Errorf("Expected total count to be 1, got %d", result.TotalCount())
	}

	if len(result.Binaries) != 1 {
		t.Errorf("Expected 1 binary, got %d", len(result.Binaries))
	}
}

func TestScanResultDetectConflicts(t *testing.T) {
	result := NewScanResult()

	// Add two versions of the same binary
	result.AddBinary(&Binary{
		Name:    "node",
		Path:    "/usr/local/bin/node",
		Manager: "homebrew",
		Version: "20.0.0",
	})

	result.AddBinary(&Binary{
		Name:    "node",
		Path:    "/home/user/.nvm/node",
		Manager: "manual",
		Version: "18.0.0",
	})

	result.DetectConflicts()

	if result.ConflictCount() != 1 {
		t.Errorf("Expected 1 conflict, got %d", result.ConflictCount())
	}

	if _, exists := result.Conflicts["node"]; !exists {
		t.Error("Expected conflict for 'node' binary")
	}

	if len(result.Conflicts["node"]) != 2 {
		t.Errorf("Expected 2 conflicting versions, got %d", len(result.Conflicts["node"]))
	}
}

func TestScanResultGhostCount(t *testing.T) {
	result := NewScanResult()

	result.AddBinary(&Binary{
		Name:    "test1",
		Manager: "manual",
	})

	result.AddBinary(&Binary{
		Name:    "test2",
		Manager: "homebrew",
	})

	if result.GhostCount() != 1 {
		t.Errorf("Expected 1 ghost binary, got %d", result.GhostCount())
	}
}
