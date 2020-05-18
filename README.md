# Terminal Stationboard UI

The Terminal Stationboard UI is a showcase application for the package [OpenTransport-Client](https://github.com/minderjan/opentransport-client).

![overview](screenshot.png)

# Install on your System
You can install the application by sources or download a compiled binary form the [release page](https://github.com/minderjan/terminal-stationboard-ui/releases).
Download the Repository, navigate to the root folder and run:
```
make install
```

# Usage
You have to provide a valid station name or id which exists on [Swiss Public Transport API](https://transport.opendata.ch/docs.html#locations). Optionally you can choose an alternative theme.

```
Usage of stationboard:
  -station string
    	Name or the id from of a station
  -theme string
    	blue, light, dark (default "dark")
```

# Compile
You can use the predefined compile commands from the Makefile. 
This produces several binaries for each operating system.
```
make compile
```

# Run the application from sources
```
go run cmd/stationboard/main.go -station {station-name} -theme {theme}
```

# Dependencies
This application was build with following dependencies:

* [OpenTransport Client](https://github.com/minderjan/opentransport-client)
* [tview](https://github.com/rivo/tview)
* [tcell](https://github.com/gdamore/tcell)
