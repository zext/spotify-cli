package cmd

import (
	"github.com/spf13/cobra"
)

var (
	client  string
	Version string = "dev"
)

var rootCmd = &cobra.Command{
	Use:     "spotify-cli",
	Short:   "Simple Spotify CLI",
	Long:    `spotify-cli is a Simple Spotify CLI that supports DBus and web API.`,
	Version: Version,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&client, "client", "c", "dbus", "API client (dbus or web)")
}

func Execute() {
	rootCmd.Execute()
}
