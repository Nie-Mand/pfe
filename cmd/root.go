package cmd

import (
	"Nie-Mand/pfender/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "pfe",
	Short: "Send Emails for PFEs",
	Long: "End of Study Internship",
	Run: func(cmd *cobra.Command, args []string) {
		utils.GetArtwork()
	},
  }
  
  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }