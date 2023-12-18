/*
Copyright © 2023 Snehil Shah <snehilshah.989@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Host struct {
	IP   string
	Name string
}

func verifyDir(directory string) bool {
	_, err := os.Stat(directory)
	return !os.IsNotExist(err)
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
		chooseHost, _ := cmd.Flags().GetBool("address")
		hostName := "localhost"
		if chooseHost {
			hostName = SelectHost()
			time.Sleep(100 * time.Millisecond)
		}
		if !PortAvailable(hostName, port) {
			fmt.Printf("\nPort %d is \033[31munavailable\033[0m\n", port)
		} else if !verifyDir(directory) {
			fmt.Printf("\nDirectory does \033[31mNOT\033[0m Exist!\n")
		} else {
			server := &http.Server{
				Addr:    fmt.Sprintf("%s:%d", hostName, port),
				Handler: http.FileServer(http.Dir(directory)),
			}
			go func() {
				err := server.ListenAndServe()
				if err != nil && err != http.ErrServerClosed {
					fmt.Printf("Server \033[31mFailed\033[0m: %v", err)
				}
			}()
			fmt.Printf("\nServer Live on \033[34mhttp://%v:%v\033[0m\n\n-> Press \033[31mEnter\033[0m to Stop the Server..", hostName, port)
			fmt.Scanln()
			server.Shutdown(context.Background())
			fmt.Print("\nServer Stopped!")
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntP("port", "p", 80, "Port to Serve on")
	serveCmd.Flags().StringP("directory", "d", ".", "Directory to Serve")
	serveCmd.Flags().BoolP("address", "a", false, "Select Server Host")
}
