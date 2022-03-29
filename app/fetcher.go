package app

import (
	"fmt"
	"io"
	"text/template"

	"github.com/zext/spotify-cli/model"
)

type Fetcher interface {
	CurrentTrack() (*model.Track, error)
}

type FetcherService struct {
	Fetcher Fetcher
	Writer  io.Writer
}

func (s *FetcherService) CurrentTrack(format string) error {
	track, err := s.Fetcher.CurrentTrack()
	if err != nil {
		return fmt.Errorf("failed to get current track: %w", err)
	}

	tmpl, err := template.New("").Parse(format)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	if err := tmpl.Execute(s.Writer, track); err != nil {
		return fmt.Errorf("failed to execute template")
	}

	return nil
}
