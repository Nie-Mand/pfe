package cmd

import (
	"Nie-Mand/pfender/utils"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringP("file", "f", "", "Email CSV File (required)")
	sendCmd.MarkFlagRequired("file")
}

var sendCmd = &cobra.Command{
  Use:   "send",
  Short: "Send Emails for PFEs",
  Long:  `Send Emails for PFEs`,
  Run: func(cmd *cobra.Command, args []string) {
	  utils.GetArtwork()

	  file, _ := cmd.Flags().GetString("file")

	  emails := utils.LoadEmails(file)
	  
	  utils.Prepare(len(emails))

	  for _, email := range emails {
		go utils.Send(email)
	  }

	  log.Println("Waiting for all emails to be sent...")

	  utils.Wait()
  },
}