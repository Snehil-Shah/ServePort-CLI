/*
Copyright © 2023 Snehil Shah <snehilshah.989@gmail.com>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

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

		go func() {
			http.Handle("/", http.FileServer(http.Dir(directory)))
			http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		}()
		fmt.Printf("\nServer Live on \033[34mhttp://localhost:%v\033[0m\n\n-> Press \033[31mEnter\033[0m to Stop the Server..", port)
		fmt.Scanln()
		fmt.Print("\nServer Stopped!")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntP("port", "p", 80, "Port to Serve on")
	serveCmd.Flags().StringP("directory", "d", ".", "Directory to Serve")
	serveCmd.Flags().BoolP("host", "h", false, "Select Server Host")
}
