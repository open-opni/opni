package commands

import (
	"fmt"
	"runtime/debug"

	"github.com/open-panoptes/opni/pkg/versions"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func BuildVersionCmd() *cobra.Command {
	var quiet bool
	var verbose bool
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show opni version information",
		Run: func(cmd *cobra.Command, args []string) {
			info, _ := debug.ReadBuildInfo()
			if verbose {
				fmt.Println(info.String())
				return
			}
			settings := lo.KeyBy(info.Settings, func(v debug.BuildSetting) string {
				return v.Key
			})
			var noun string
			var version string
			if versions.Version == "unversioned" {
				noun = "revision"
				version = settings["vcs.revision"].Value
			} else {
				noun = "version"
				version = versions.Version
			}
			if quiet {
				fmt.Println(version)
				return
			}
			fmt.Printf("Opni, %s %s\n", noun, version)
			fmt.Printf("  go version: %s\n", info.GoVersion)
		},
	}
	versionCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Only print version or revision")
	versionCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show verbose version and dependency information")
	return versionCmd
}
