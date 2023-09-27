package cmd

import (
	"Nie-Mand/pfender/utils"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringP("file", "f", "", "Email CSV File (required)")
	sendCmd.Flags().StringP("config", "c", "config.txt", "Config File (by default config.txt)")
	sendCmd.MarkFlagRequired("file")
}

var sendCmd = &cobra.Command{
  Use:   "send",
  Short: "Send Emails for PFEs",
  Long:  `Send Emails for PFEs`,
  Run: func(cmd *cobra.Command, args []string) {
	  utils.GetArtwork()
	  
	  file, _ := cmd.Flags().GetString("file")
	  config, _ := cmd.Flags().GetString("config")
	  godotenv.Load(config)

	  emails := utils.LoadEmails(file)
	  filtredEmails := utils.FilterRows(emails)
	  utils.Prepare(len(filtredEmails))

	  for _, email := range filtredEmails {
		go utils.Send(email)
		utils.AddHash(email)
	  }

	  log.Println("Waiting for all emails to be sent...")

	  utils.Wait()
	  utils.SaveNewHashes()
  },
}