# gompris-current

gompris-current is a program that can get current song of media player that support
Mpris DBus protocol. It's output the status to `$XDG_CACHE_HOME/mpris-current`
or `$HOME/.cache/mpris-current` file. You can use the file as a text source
in OBS or other application.

## Installation

You can use the following command to install `gompris-current`:

```
go get -u github.com/yursan9/gompris-current
```

It will download the sources code and install the program into the `$GOPATH/bin`
directory. Then you can add `$GOPATH/bin` to your `$PATH` variabel.

## Running The Program

Execute the following command to run the program:

```
gompris-current [media player name]
```

You need to input the name of your media player as the command line argument.
Look at the table to see what name you need to input as argument:

| Program Name |  Input Name  |
|--------------|:------------:|
|Rhythmbox     | rhythmbox    |
|Spotify       | spotify      |
|VLC           | vlc          |
|Lollypop      | lollypop     |
|MPD           | mpd          |
|CMUS          | cmus         |

Example usage:

```
gompris-current rhythmbox
2019/05/15 13:12:57 Subscibe to org.freedesktop.DBus.Properties.PropertiesChanged signal
2019/05/15 13:12:57 Start monitoring signal for org.mpris.MediaPlayer2.rhythmbox
2019/05/15 13:12:57 Press Ctrl+C to quit...
```

## Reading The Current Song

The currently played song is listed in `$HOME/.cache/mpris-current`. You can
use the file with OBS by following these steps:

1. Click `+` in Sources window.
2. Click Text on the popup.
3. Enter the new source's name then click OK.
4. Tick the options "Read from file".
5. Then enter the name of text file by clicking the Browse button.
6. Finally, click OK to finish.

## Editing the Status Text

You can configure the outputted text by using template. Create the file in
`$HOME/.config/mpris-current/` directory and the file as `template`.
If there is no `template` file in `$HOME/.config/mpris-current`, the default
template is:

```
{{.Status}}: {{.Title}} - {{.Artist}}
```

The syntax follows the golang template's syntax.

## License

Copyright 2019 Yurizal Susanto

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
