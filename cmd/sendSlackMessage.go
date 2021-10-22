package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var sendSlackMessageCmd = &cobra.Command{
	Use:   "slack",
	Short: "Send Slack message",
	Long:  `Send Generic Slack message`,
	Run: func(cmd *cobra.Command, args []string) {
		err := sendSlackMessage(args)
		if err != nil {
			fmt.Println(err)
		}
	},
}

var params []string

func init() {
	rootCmd.AddCommand(sendSlackMessageCmd)
	sendSlackMessageCmd.PersistentFlags().StringArrayVar(&params, "message", params, "slackwebhook")
}

func sendSlackMessage(args []string) error {
	block := strings.NewReader(args[0])
	slackWebhook := args[1]

	fmt.Println(block)
	fmt.Println(slackWebhook)

	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, slackWebhook, block)
	if err != nil {
		log.Println("Http Error: ", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Http:", err)
	}

	fmt.Println(resp)

	return nil
}
