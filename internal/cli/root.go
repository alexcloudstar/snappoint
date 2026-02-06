package cli

import (
	"fmt"
	"os"

	"github.com/alexcloudstar/snappoint/pkg/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "snappoint",
	Short: "SnapPoint - System auditor and package manager manager",
	Long: `SnapPoint is an open-source CLI tool that helps developers identify and manage
binaries on their system. It discovers packages from multiple package managers
(Homebrew, NPM, Pip) and identifies "ghost" binaries that aren't managed by any
package manager.`,
	Version: version.Version,
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate(`{{printf "SnapPoint v%s\n" .Version}}`)
}
