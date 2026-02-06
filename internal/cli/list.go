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
	listOrphans   bool
	listConflicts bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List binaries found in previous scans",
	Long: `List binaries discovered on your system. By default, shows all binaries.
Use --orphans to show only ghost binaries or --conflicts to show only
binaries with version conflicts.`,
	RunE: runList,
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&listOrphans, "orphans", false, "Show only ghost binaries")
	listCmd.Flags().BoolVar(&listConflicts, "conflicts", false, "Show only conflicting versions")
}

func runList(cmd *cobra.Command, args []string) error {
	// For now, list command performs a fresh scan
	// In the future, this could read from cached results
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

	fmt.Println("Scanning system...")

	result, err := s.Scan(ctx)
	if err != nil {
		fmt.Printf("Warning: %v\n", err)
	}

	// Scan for manual/ghost binaries
	manualMgr := managers.NewManual(executor)
	manualMgr.SetKnownBinaries(result.Binaries)

	manualBinaries, err := manualMgr.Scan(ctx)
	if err != nil {
		fmt.Printf("Warning: manual scan failed: %v\n", err)
	} else {
		for _, binary := range manualBinaries {
			result.AddBinary(binary)
		}
		result.DetectConflicts()
	}

	// Format and display results
	formatter := output.NewTableFormatter()

	if listOrphans {
		formatter.SetShowGhostsOnly(true)
	} else if listConflicts {
		formatter.SetShowConflictsOnly(true)
	}

	return formatter.Format(result)
}
