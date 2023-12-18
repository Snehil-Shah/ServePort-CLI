/*
Copyright © 2023 Snehil Shah <snehilshah.989@gmail.com>
*/
package cmd

import (
	"fmt"
	"net"
	"time"

	"github.com/spf13/cobra"
)

var checkPortCmd = &cobra.Command{
	Use:   "check-port",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		chooseHost, _ := cmd.Flags().GetBool("address")
		hostName := "localhost"
		if chooseHost {
			hostName = SelectHost()
			time.Sleep(100 * time.Millisecond)
		}
		connection, err := net.Dial("tcp", fmt.Sprintf("%s:%d", hostName, port))
		if err != nil {
			fmt.Printf("\nPort %d is \033[32mavailable\033[0m\n", port)
		} else {
			fmt.Printf("\nPort %d is \033[31munavailable\033[0m\n", port)
			connection.Close()
		}
	},
}

func init() {
	rootCmd.AddCommand(checkPortCmd)

	checkPortCmd.Flags().IntP("port", "p", 80, "Check Port Availability")
	checkPortCmd.Flags().BoolP("address", "a", false, "Select Host for Port Checking")
}
