package scanner

import (
	"context"
)

// PackageManager defines the interface for package manager implementations
type PackageManager interface {
	// Name returns the name of the package manager
	Name() string

	// IsAvailable checks if this package manager is installed on the system
	IsAvailable(ctx context.Context) bool

	// Scan discovers all binaries managed by this package manager
	Scan(ctx context.Context) ([]*Binary, error)
}
