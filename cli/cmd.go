package cli

import (
	"github.com/spf13/cobra"
	"log"
)

func Execute() {
	var rootCmd = &cobra.Command{Use: "microservice_order_history"}

	ExecuteServeHttp(rootCmd)

	err := rootCmd.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
