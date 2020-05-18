package main

import (
	"flag"
	"fmt"
	"github.com/minderjan/opentransport-client/opentransport"
	"github.com/minderjan/terminal-stationboard-ui/transport"
	"github.com/minderjan/terminal-stationboard-ui/ui"
	"os"
	"time"
)

// Interval to update the stationboard data from the api
const refreshInterval = 5 * time.Minute

// Interval of table updates (remove missed connections, based on cached data)
const tableRefresh = 20 * time.Second

// Interval to updated the timer in top right corner
const timeRefreshInterval = 1 * time.Second

// The minimum amount of time a connection must fit
const minConnectionDeparture = 3 * time.Minute

func main() {

	var theme string
	var station string

	flag.StringVar(&station,"station", "", "Name or the id from of a station")
	flag.StringVar(&theme, "theme", "dark", "blue, light, dark")

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(0)
	}

	flag.Parse()

	if len(station) == 0 {
		fmt.Println("Provide a station name or id. Use -station flag.")
		fmt.Println("Only stations within switzerland will be searched.")
		fmt.Println("Optional an id can be used to match a station.")
		fmt.Println("The id should match a location on https://transport.opendata.ch")
		os.Exit(1)
	}

	fmt.Printf("Search for a Station: %s\n", station)
	fmt.Printf("Requested Theme: %s\n", theme)

	// Load Data
	locs := transport.LoadLocations(station)
	stb := transport.LoadStationboard(station, minConnectionDeparture)

	if len(stb.Journeys) == 0 {
		fmt.Printf("No station found with name %s\n", station)

		var foundStations []opentransport.Location
		for _, l := range locs {
			if l.Station() {
				foundStations = append(foundStations, l)
			}
		}

		if len(foundStations) > 0 {
			fmt.Printf("Did you mean one of following location: \n")
			for _, s := range foundStations {
				fmt.Printf("%s (%s)\n", s.Name, s.Id)
			}
			fmt.Printf("You can also use the id of the station instead of the name.\n")
		}
		os.Exit(1)
	}

	// Start UI
	gui := ui.NewUIWithTheme(station, &stb.Station, minConnectionDeparture, ui.LoadTheme(theme))
	gui.AddStationboard(stb.Journeys)
	gui.AddLocations(locs)

	// Start update threads
	go gui.UpdateTime(timeRefreshInterval)
	go gui.Update(refreshInterval)
	go gui.UpdateStationboardTime(tableRefresh)
	go gui.UpdateIndicator()

	// Start the Dashboard
	err := gui.Run()
	if err != nil {
		panic(err)
	}
}

