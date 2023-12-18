/*
Copyright © 2023 Snehil Shah <snehilshah.989@gmail.com>
*/
package cmd

import (
	"fmt"
	"net"
	"strings"

	"github.com/spf13/cobra"
)

func GetHosts() []Host {
	var hosts []Host
	interfaces, _ := net.Interfaces()
	for _, hostName := range interfaces {
		addrs, _ := hostName.Addrs()
		for _, addr := range addrs {
			if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
				if ip.IP.To4() != nil && !strings.HasPrefix(ip.IP.String(), "169.254") {
					hosts = append(hosts, Host{IP: ip.IP.String(), Name: hostName.Name})
				}
			}
		}
	}
	return hosts
}

var listInterfacesCmd = &cobra.Command{
	Use:   "list-interfaces",
	Short: "List your Network Interfaces",
	Long: `List all Network Interfaces on your Computer with their Details.

Network Interfaces are Endpoints that connect your computer to various Networks.

Example of these Networks: A Wifi Router, or an Ethernet Connection etc.
that are all connected to you via a Network Interface with a unique IP-Address on that Network!

View MAC, IPv4, IPv6 Addresses and other information of all Interfaces on your Machine with just One Command.
`,
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
