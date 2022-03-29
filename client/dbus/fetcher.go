package dbus

import (
	"fmt"

	"github.com/dawidd6/go-spotify-dbus"
	"github.com/godbus/dbus"
	"github.com/zext/spotify-cli/model"
)

type Fetcher struct {
	conn *dbus.Conn
}

func NewFetcher() (*Fetcher, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, fmt.Errorf("failed to create DBus session: %w", err)
	}

	return &Fetcher{conn: conn}, nil
}

func (f *Fetcher) CurrentTrack() (*model.Track, error) {
	metadata, err := spotify.GetMetadata(f.conn)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata: %w", err)
	}

	return &model.Track{
		Album:   metadata.Album,
		Artists: metadata.Artist,
		Title:   metadata.Title,
		URL:     metadata.URL,
	}, nil
}
