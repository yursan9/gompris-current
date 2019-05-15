package main

import (
	"github.com/godbus/dbus"
)

type PlayerStatus struct {
	Status string
	Artist string
	Title  string
}

func NewPlayerStatus(obj dbus.BusObject) *PlayerStatus {
	// Get Playback Status from mpris
	reply, err := obj.GetProperty("org.mpris.MediaPlayer2.Player.PlaybackStatus")
	if err != nil {
		return nil
	}
	status := reply.Value().(string)

	// Get Metadata from mpris
	reply, err = obj.GetProperty("org.mpris.MediaPlayer2.Player.Metadata")
	if err != nil {
		return nil
	}
	meta := reply.Value().(map[string]dbus.Variant)
	// Extract Artist and Title from metadata
	artist, title := GetArtistAndTitle(meta)

	return &PlayerStatus{
		Status: status,
		Artist: artist,
		Title:  title,
	}
}

func GetArtistAndTitle(m map[string]dbus.Variant) (string, string) {
	if len(m) == 0 {
		return "", ""
	}

	artist := m["xesam:artist"].Value().([]string)[0]
	title := m["xesam:title"].Value().(string)

	return artist, title
}
