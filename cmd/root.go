/*
Copyright © 2023 Snehil Shah snehilshah.989@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ServePort",
	Short: "A convenient CLI tool for Spawning Servers and managing Ports, IP-Addresses & Network Interfaces",
	Long: `ServePort is a simple CLI tool to Control & Manage Servers, Ports & Network Devices.

-> Spawn Static Servers over Local, Private & Public Networks allowing you to share your files with ease
-> Check Port Availability on different IP-Addresses
-> Lists all Network Interfaces on your machine along with their MAC, IPv4 & IPv6 Addresses

WARNING: Make sure you trust the Private or Public network you are serving over.
Strictly avoid Serving over Public Networks as you are more vulnerable Unrestricted Access to your Server
-> It's recommended to stick to Localhost for most menial tasks!`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
