package cli

import (
	"github.com/spf13/cobra"
	"log"
)

func Execute() {
	var rootCmd = &cobra.Command{Use: "simpleapi"}

	ExecuteServeHttp(rootCmd)

	err := rootCmd.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
