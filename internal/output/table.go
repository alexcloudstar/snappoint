package output

import (
	"fmt"
	"os"

	"github.com/alexcloudstar/snappoint-cli/internal/scanner"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

// TableFormatter formats output as a table
type TableFormatter struct {
	showGhostsOnly    bool
	showConflictsOnly bool
}

// NewTableFormatter creates a new table formatter
func NewTableFormatter() *TableFormatter {
	return &TableFormatter{}
}

// SetShowGhostsOnly sets whether to show only ghost binaries
func (t *TableFormatter) SetShowGhostsOnly(show bool) {
	t.showGhostsOnly = show
}

// SetShowConflictsOnly sets whether to show only conflicting binaries
func (t *TableFormatter) SetShowConflictsOnly(show bool) {
	t.showConflictsOnly = show
}

// Format outputs the scan results as a formatted table
func (t *TableFormatter) Format(result *scanner.ScanResult) error {
	binaries := result.Binaries

	// Filter based on options
	if t.showGhostsOnly {
		binaries = result.Ghosts
	} else if t.showConflictsOnly {
		var conflicting []*scanner.Binary
		for _, bins := range result.Conflicts {
			conflicting = append(conflicting, bins...)
		}
		binaries = conflicting
	}

	if len(binaries) == 0 {
		fmt.Println("No binaries found.")
		return nil
	}

	// Create table
	table := tablewriter.NewWriter(os.Stdout)
	table.Header("NAME", "PATH", "MANAGER", "VERSION")

	// Add rows
	for _, binary := range binaries {
		manager := binary.Manager
		if binary.IsGhost() {
			manager = "👻 ghost"
		}

		version := binary.Version
		if version == "" || version == "unknown" {
			version = "-"
		}

		if err := table.Append([]string{
			binary.Name,
			binary.Path,
			manager,
			version,
		}); err != nil {
			return err
		}
	}

	if err := table.Render(); err != nil {
		return err
	}

	// Print summary
	fmt.Println()
	t.printSummary(result)

	return nil
}

// printSummary prints a summary of the scan results
func (t *TableFormatter) printSummary(result *scanner.ScanResult) {
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Printf("%s Total binaries found: %d\n", cyan("ℹ"), result.TotalCount())

	if result.ConflictCount() > 0 {
		fmt.Printf("%s Found %d conflicts:\n", yellow("⚠️ "), result.ConflictCount())
		for name, bins := range result.Conflicts {
			fmt.Printf("  • %s: %d versions detected\n", name, len(bins))
			for _, bin := range bins {
				fmt.Printf("    - %s (%s)\n", bin.Path, bin.Manager)
			}
		}
		fmt.Println()
	}

	if result.GhostCount() > 0 {
		fmt.Printf("%s Found %d ghost binaries:\n", red("👻"), result.GhostCount())
		for _, ghost := range result.Ghosts {
			fmt.Printf("  • %s: No package manager claims this (%s)\n", ghost.Name, ghost.Path)
		}
	}
}
