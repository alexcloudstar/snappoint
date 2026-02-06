package cli

import (
	"context"
	"fmt"

	"github.com/alexcloudstar/snappoint-cli/internal/managers"
	"github.com/alexcloudstar/snappoint-cli/internal/output"
	"github.com/alexcloudstar/snappoint-cli/internal/scanner"
	"github.com/alexcloudstar/snappoint-cli/pkg/system"
	"github.com/spf13/cobra"
)

var (
	scanManager string
	scanOutput  string
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan system for binaries",
	Long: `Scan your system to discover binaries managed by package managers
(Homebrew, NPM, Pip) and identify ghost binaries that aren't claimed
by any package manager.`,
	RunE: runScan,
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVar(&scanManager, "manager", "", "Filter by package manager (homebrew, npm, pip, manual)")
	scanCmd.Flags().StringVar(&scanOutput, "output", "table", "Output format (table, json)")
}

func runScan(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	executor := system.NewExecutor()

	// Initialize package managers
	mgrs := []scanner.PackageManager{
		managers.NewHomebrew(executor),
		managers.NewNPM(executor),
		managers.NewPip(executor),
	}

	// Create scanner
	s := scanner.NewScanner(mgrs...)

	fmt.Println("Scanning system for binaries...")

	var result *scanner.ScanResult
	var err error

	// Scan specific manager or all
	if scanManager != "" {
		result, err = s.ScanSingle(ctx, scanManager)
	} else {
		result, err = s.Scan(ctx)
	}

	if err != nil {
		fmt.Printf("Warning: %v\n", err)
	}

	// Now scan for manual/ghost binaries if not filtering by a specific manager
	if scanManager == "" || scanManager == "manual" {
		manualMgr := managers.NewManual(executor)
		manualMgr.SetKnownBinaries(result.Binaries)

		manualBinaries, err := manualMgr.Scan(ctx)
		if err != nil {
			fmt.Printf("Warning: manual scan failed: %v\n", err)
		} else {
			for _, binary := range manualBinaries {
				result.AddBinary(binary)
			}
			// Re-detect conflicts after adding manual binaries
			result.DetectConflicts()
		}
	}

	// Format and display results
	formatter := output.NewTableFormatter()
	return formatter.Format(result)
}
