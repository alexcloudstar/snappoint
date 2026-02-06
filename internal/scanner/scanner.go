package scanner

import (
	"context"
	"fmt"
	"sync"
)

// Scanner orchestrates scanning across multiple package managers
type Scanner struct {
	managers []PackageManager
}

// NewScanner creates a new scanner with the given package managers
func NewScanner(managers ...PackageManager) *Scanner {
	return &Scanner{
		managers: managers,
	}
}

// Scan runs all package managers in parallel and aggregates results
func (s *Scanner) Scan(ctx context.Context) (*ScanResult, error) {
	result := NewScanResult()

	var wg sync.WaitGroup
	var mu sync.Mutex
	errChan := make(chan error, len(s.managers))

	for _, manager := range s.managers {
		// Skip if manager is not available
		if !manager.IsAvailable(ctx) {
			continue
		}

		wg.Add(1)
		go func(mgr PackageManager) {
			defer wg.Done()

			binaries, err := mgr.Scan(ctx)
			if err != nil {
				errChan <- fmt.Errorf("%s scan failed: %w", mgr.Name(), err)
				return
			}

			mu.Lock()
			for _, binary := range binaries {
				result.AddBinary(binary)
			}
			mu.Unlock()
		}(manager)
	}

	wg.Wait()
	close(errChan)

	// Collect any errors (but don't fail if one manager fails)
	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}

	// Detect conflicts after all binaries are collected
	result.DetectConflicts()

	// Return errors if any occurred, but still return the results
	if len(errs) > 0 {
		return result, fmt.Errorf("some scans failed: %v", errs)
	}

	return result, nil
}

// ScanSingle scans a single package manager
func (s *Scanner) ScanSingle(ctx context.Context, managerName string) (*ScanResult, error) {
	result := NewScanResult()

	var targetManager PackageManager
	for _, mgr := range s.managers {
		if mgr.Name() == managerName {
			targetManager = mgr
			break
		}
	}

	if targetManager == nil {
		return nil, fmt.Errorf("package manager '%s' not found", managerName)
	}

	if !targetManager.IsAvailable(ctx) {
		return nil, fmt.Errorf("package manager '%s' is not available on this system", managerName)
	}

	binaries, err := targetManager.Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s scan failed: %w", targetManager.Name(), err)
	}

	for _, binary := range binaries {
		result.AddBinary(binary)
	}

	result.DetectConflicts()

	return result, nil
}
