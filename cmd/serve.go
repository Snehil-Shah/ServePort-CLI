/*
Copyright © 2023 Snehil Shah <snehilshah.989@gmail.com>
*/
package cmd

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type Host struct {
	IP   string
	Name string
}

func getHosts() []Host {
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

func selectHost() string {
	hosts := getHosts()
	var items = []Host{{IP: "127.0.0.1", Name: "127.0.0.1 - Localhost"}}
	for _, host := range hosts {
		items = append(items, Host{Name: fmt.Sprintf("%v - %v", host.IP, host.Name), IP: host.IP})
	}
	fmt.Printf("\n")
	prompt := promptui.Select{
		Label:    "Select Server Host",
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

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts a local HTTP Server",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		directory, _ := cmd.Flags().GetString("directory")
		chooseHost, _ := cmd.Flags().GetBool("host")
		hostName := "localhost"
		if chooseHost {
			hostName = selectHost()
			time.Sleep(100 * time.Millisecond)
		}
		go func() {
			http.Handle("/", http.FileServer(http.Dir(directory)))
			http.ListenAndServe(fmt.Sprintf("%s:%d", hostName, port), nil)
		}()
		fmt.Printf("\nServer Live on \033[34mhttp://%v:%v\033[0m\n\n-> Press \033[31mEnter\033[0m to Stop the Server..", hostName, port)
		fmt.Scanln()
		fmt.Print("\nServer Stopped!")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntP("port", "p", 80, "Port to Serve on")
	serveCmd.Flags().StringP("directory", "d", ".", "Directory to Serve")
	serveCmd.Flags().BoolP("host", "i", false, "Select Server Host")
}
