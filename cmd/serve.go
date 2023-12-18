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
	Short: "Starts a quick HTTP Server",
	Long: `Boot up a Quick, Simple & Usable Static File Server over Local, connected Private/Public or All Networks!

Serve specific Directories on any Port & Network using designated flags to configure the server
Serving files over Private networks like Wifi/Ethernet makes Sharing files to other Computers quick & simple

WARNING: Make sure you trust the Private or Public network you are serving over.
Strictly avoid Serving over Public Networks like Public Wifi's as you are more vulnerable to Unrestricted Access to your Server
-> It's recommended to stick to Localhost for most menial tasks!`,

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
			if hostName == "0.0.0.0" || hostName == "127.0.0.1" || hostName == "localhost" {
				fmt.Printf("\nServer Live on \033[34mhttp://localhost:%v\033[0m\n\n-> Press \033[31mEnter\033[0m to Stop the Server..", port)
			} else {
				fmt.Printf("\nServer Live on \033[34mhttp://%v:%v\033[0m\n\n-> Press \033[31mEnter\033[0m to Stop the Server..", hostName, port)
			}
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
	serveCmd.Flags().BoolP("address", "a", false, "Select Server Host?")
}
