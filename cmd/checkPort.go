/*
Copyright © 2023 Snehil Shah <snehilshah.989@gmail.com>
*/
package cmd

import (
	"fmt"
	"net"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func PortAvailable(hostName string, port int) bool {
	_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", hostName, port))
	return err != nil && port < 65535
}

func SelectHost() string {
	hosts := GetHosts()
	var items = []Host{{IP: "127.0.0.1", Name: "127.0.0.1 - Localhost"}}
	for _, host := range hosts {
		items = append(items, Host{Name: fmt.Sprintf("%v - %v", host.IP, host.Name), IP: host.IP})
	}
	fmt.Printf("\n")
	prompt := promptui.Select{
		Label:    "Select Host Address",
		Items:    items,
		HideHelp: true,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}:",
			Active:   "> {{ .Name | magenta }}",
			Inactive: "  {{ .Name }}",
			Selected: "> {{ .Name | magenta }}",
		},
	}
	i, _, _ := prompt.Run()
	return items[i].IP
}

var checkPortCmd = &cobra.Command{
	Use:   "check-port",
	Short: "Checks Status of a Port",
	Long: `Check Port Availability of any/all of your IP-Addresses with ease

This can help you manage your services and apps better,
or maybe hunt down nefarious hidden services!`,

	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		chooseHost, _ := cmd.Flags().GetBool("address")
		hostName := "localhost"
		if chooseHost {
			hostName = SelectHost()
			time.Sleep(100 * time.Millisecond)
		}
		if PortAvailable(hostName, port) {
			fmt.Printf("\nPort %d is \033[32mavailable\033[0m\n", port)
		} else {
			fmt.Printf("\nPort %d is \033[31munavailable\033[0m\n", port)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkPortCmd)

	checkPortCmd.Flags().IntP("port", "p", 80, "Check Port for Availability")
	checkPortCmd.Flags().BoolP("address", "a", false, "Select IP-Address for Port Checking")
}
