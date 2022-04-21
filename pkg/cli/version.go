package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/27149cheo/helmtool/version"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version",
	Long:  `Print the version.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		doVersion()
	},
}

func doVersion() {
	fmt.Printf("%#v\n", version.Get())
}
