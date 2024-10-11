package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vitorsoratto/fbadmgenerator/config"
	"github.com/vitorsoratto/fbadmgenerator/firebase"
)

var generateCmd = &cobra.Command{
	Use:   "fbadmgen",
	Short: "Generates a credentials token for Firebase Admin SDK",
	Long: `Generates a credentials token for Firebase Admin SDK

This command will generate a credentials token for Firebase Admin SDK based on the JSON credentials file provided.

The token will be printed to stdout.
`,
}

func init() {
	logger := config.GetLogger()
	generateCmd.PersistentFlags().StringVarP(&firebase.JSONCredentials, "json", "j", "", "Path to the JSON credentials file")

	generateCmd.Run = func(cmd *cobra.Command, args []string) {
		if firebase.JSONCredentials == "" {
			logger.Error("JSON credentials file path is required\n")
			cmd.Usage()
			os.Exit(1)
		}

		token, err := firebase.NewTokenProvider()
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}
		logger.Infof("Firebase Cloud Messaging token: %v", *token)
	}
}

func Execute() {
	if err := generateCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
