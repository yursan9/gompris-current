//Copyright 2019 Yurizal Susanto

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at

//   http://www.apache.org/licenses/LICENSE-2.0

//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus"
)

func main() {
	// Get player name from command line args
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Need media player name. Usage: mpris-current [player name]")
		os.Exit(1)
	}
	player := os.Args[1]

	// Setup player status template
	template := SetupTemplate()

	// Connect to dbus
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	// Make new player status based on dbus object
	obj := conn.Object("org.mpris.MediaPlayer2."+player, "/org/mpris/MediaPlayer2")
	status := NewPlayerStatus(obj)
	if status == nil {
		fmt.Fprintln(os.Stderr, "Can't connect to dbus object", obj.Destination())
		fmt.Fprintln(os.Stderr, "Make sure", player, "is running")
		os.Exit(1)
	}

	// Subscibe to PropertiesChanged signal
	conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='/org/mpris/MediaPlayer2',interface='org.freedesktop.DBus.Properties',member=PropertiesChanged")

	// Make signal
	c := make(chan *dbus.Signal, 10)
	conn.Signal(c)

	fmt.Println("Start monitoring signal for org.mpris.MediaPlayer2." + player)
	fmt.Println("Press Ctrl+C to quit...")
	// Loop over the channel
	for v := range c {
		// Get signal's message body
		reply := v.Body[1].(map[string]dbus.Variant)

		// If Playback Status changed, update the player status
		if ps, ok := reply["PlaybackStatus"]; ok {
			status.Status = ps.Value().(string)
		}

		// If songs changed, update the artist and title
		if metadata, ok := reply["Metadata"]; ok {
			metadata := metadata.Value().(map[string]dbus.Variant)

			status.Artist, status.Title = GetArtistAndTitle(metadata)
		}

		// Open file for player status
		f := CreateFile()
		defer f.Close()

		// Merge player status with template, and write to file
		template.Execute(f, status)
	}
}
