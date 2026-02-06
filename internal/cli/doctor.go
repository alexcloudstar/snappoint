package cli

import (
	"context"
	"fmt"

	"github.com/alexcloudstar/snappoint-cli/internal/managers"
	"github.com/alexcloudstar/snappoint-cli/pkg/system"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check system health and available package managers",
	Long:  `Run system diagnostics to check which package managers are available and identify potential issues.`,
	RunE:  runDoctor,
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}

func runDoctor(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	executor := system.NewExecutor()

	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Println("Running system diagnostics...")
	fmt.Println()

	// Check platform
	platform := system.GetPlatform()
	fmt.Printf("%s Platform: %s\n", cyan("ℹ"), platform.String())
	fmt.Println()

	// Check package managers
	fmt.Println("Package Managers:")

	mgrs := []struct {
		name    string
		manager interface{ IsAvailable(context.Context) bool }
	}{
		{"Homebrew", managers.NewHomebrew(executor)},
		{"NPM", managers.NewNPM(executor)},
		{"Pip", managers.NewPip(executor)},
	}

	for _, m := range mgrs {
		status := red("✗ Not available")
		if m.manager.IsAvailable(ctx) {
			status = green("✓ Available")
		}
		fmt.Printf("  %s: %s\n", m.name, status)
	}

	fmt.Println()

	// Check PATH
	paths := system.GetPATH()
	fmt.Printf("PATH directories: %d\n", len(paths))
	fmt.Println()

	// Check common binary directories
	fmt.Println("Common binary directories:")
	commonPaths := system.GetCommonBinaryPaths()
	for _, path := range commonPaths {
		fmt.Printf("  • %s\n", path)
	}

	return nil
}
