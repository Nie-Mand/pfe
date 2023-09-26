package cmd

import (
	"Nie-Mand/pfender/utils"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of PFE",
  Long:  `All software has versions. This is PFE'`,
  Run: func(cmd *cobra.Command, args []string) {
	  utils.GetArtwork()
  },
}