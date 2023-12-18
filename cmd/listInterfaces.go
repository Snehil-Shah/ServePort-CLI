/*
Copyright © 2023 Snehil Shah <snehilshah.989@gmail.com>
*/
package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

var listInterfacesCmd = &cobra.Command{
	Use:   "list-interfaces",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		interfaces, err := net.Interfaces()
		if err != nil {
			fmt.Print(err)
			return
		}
		for _, i := range interfaces {
			fmt.Printf("\n\033[36m%v\033[0m : %v\n", i.Name, i.Flags)
			addrs, err := i.Addrs()
			if err != nil {
				fmt.Print(err)
				return
			}
			fmt.Printf("   MAC Address-> %v\n", i.HardwareAddr)
			fmt.Printf("   IPv6 Address-> %v\n", addrs[0])
			fmt.Printf("   IPv4 Address-> %v\n", addrs[1])
		}
	},
}

func init() {
	rootCmd.AddCommand(listInterfacesCmd)
}
