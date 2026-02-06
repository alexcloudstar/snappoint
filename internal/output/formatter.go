package output

import (
	"github.com/alexcloudstar/snappoint/internal/scanner"
)

// Formatter defines the interface for output formatting
type Formatter interface {
	Format(result *scanner.ScanResult) error
}
