package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zext/spotify-cli/app"
	"github.com/zext/spotify-cli/client/dbus"
)

var format string

func runFetch(cmd *cobra.Command, args []string) error {
	var fetcher app.Fetcher

	switch client {
	case "dbus":
		fet, err := dbus.NewFetcher()
		if err != nil {
			return fmt.Errorf("failed to initialize DBus fetcher: %w", err)
		}

		fetcher = fet

	case "web":
		return fmt.Errorf("not implemented")
	}

	service := app.FetcherService{
		Fetcher: fetcher,
		Writer:  os.Stdout,
	}

	if err := service.CurrentTrack(format); err != nil {
		return fmt.Errorf("failed to get current track: %w", err)
	}

	return nil
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch track information",
	RunE:  runFetch,
}

func init() {
	fetchCmd.Flags().StringVarP(&format, "format", "f", "{{ .Title }} by {{ slice .Artists 0 }}\n", "format as text/template")

	rootCmd.AddCommand(fetchCmd)
}
