package dbus

import (
	"fmt"

	"github.com/dawidd6/go-spotify-dbus"
	"github.com/godbus/dbus"
)

type Operator struct {
	conn *dbus.Conn
}

func NewDBusOperator() (*Operator, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, fmt.Errorf("failed to create DBus session: %w", err)
	}

	return &Operator{conn: conn}, nil
}

func (o *Operator) Next() error {
	if err := spotify.SendNext(o.conn); err != nil {
		return fmt.Errorf("failed to send next: %w", err)
	}

	return nil
}

func (o *Operator) Pause() error {
	if err := spotify.SendPause(o.conn); err != nil {
		return fmt.Errorf("failed to send pause: %w", err)
	}

	return nil
}

func (o *Operator) Play() error {
	if err := spotify.SendPlay(o.conn); err != nil {
		return fmt.Errorf("failed to send play: %w", err)
	}

	return nil
}

func (o *Operator) PlayPause() error {
	if err := spotify.SendPlayPause(o.conn); err != nil {
		return fmt.Errorf("failed to send play/pause: %w", err)
	}

	return nil
}

func (o *Operator) Previouse() error {
	if err := spotify.SendPlayPause(o.conn); err != nil {
		return fmt.Errorf("failed to send previouse: %w", err)
	}

	return nil
}
