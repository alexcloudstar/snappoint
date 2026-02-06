package system

import (
	"testing"
)

func TestGetPlatform(t *testing.T) {
	platform := GetPlatform()

	if platform.OS == "" {
		t.Error("Expected OS to be set")
	}

	if platform.Arch == "" {
		t.Error("Expected Arch to be set")
	}
}

func TestPlatformString(t *testing.T) {
	platform := Platform{
		OS:   "darwin",
		Arch: "arm64",
	}

	expected := "darwin-arm64"
	if platform.String() != expected {
		t.Errorf("Expected %s, got %s", expected, platform.String())
	}
}

func TestIsDarwin(t *testing.T) {
	tests := []struct {
		name     string
		platform Platform
		expected bool
	}{
		{"Darwin", Platform{OS: "darwin", Arch: "amd64"}, true},
		{"Linux", Platform{OS: "linux", Arch: "amd64"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.platform.IsDarwin() != tt.expected {
				t.Errorf("Expected IsDarwin() to be %v", tt.expected)
			}
		})
	}
}

func TestIsLinux(t *testing.T) {
	tests := []struct {
		name     string
		platform Platform
		expected bool
	}{
		{"Linux", Platform{OS: "linux", Arch: "amd64"}, true},
		{"Darwin", Platform{OS: "darwin", Arch: "amd64"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.platform.IsLinux() != tt.expected {
				t.Errorf("Expected IsLinux() to be %v", tt.expected)
			}
		})
	}
}

func TestIsARM(t *testing.T) {
	tests := []struct {
		name     string
		platform Platform
		expected bool
	}{
		{"ARM64", Platform{OS: "linux", Arch: "arm64"}, true},
		{"ARM", Platform{OS: "linux", Arch: "arm"}, true},
		{"AMD64", Platform{OS: "linux", Arch: "amd64"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.platform.IsARM() != tt.expected {
				t.Errorf("Expected IsARM() to be %v", tt.expected)
			}
		})
	}
}
