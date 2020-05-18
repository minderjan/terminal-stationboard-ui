package transport

import (
	"context"
	"fmt"
	"github.com/minderjan/opentransport-client/opentransport"
	"os"
	"strconv"
	"time"
)

func newClient() *opentransport.Client {
	//client, _ := opentransport.NewClientWithUrl(nil, "http://localhost:3001/v1")
	client := opentransport.NewClient()
	//client.EnableLogs(nil)
	return client
}

// Returns a formatted Number
func TransportNumber(c, n string) string {
	_, err := strconv.Atoi(n)
	if err == nil {
		return fmt.Sprintf("%s %s", c, n)
	} else {
		return fmt.Sprintf("%s", n)
	}
}

// Returns the length of the longest destination location
func DestinationLength(conn []opentransport.StationBoardJourney) int {
	longest := 0
	for _, c := range conn {
		if len(c.To) > longest {
			longest = len(c.To)
		}
	}
	return longest
}

// Returns a formatted transport number
func ShowPlatformCol(conn []opentransport.StationBoardJourney) bool {
	for _, c := range conn {
		if len(c.Stop.Platform) > 0 {
			return true
		}
	}
	return false
}

func LoadStationboard(station string, minConnectionDeparture time.Duration) *opentransport.StationboardResult {
	when := time.Now().Add(minConnectionDeparture)
	stbRes, err := newClient().Stationboard.SearchWithDate(context.Background(), station, when)
	if err != nil {
		fmt.Printf("Failed to load stationboard for %s: %s\n", station, err)
		os.Exit(1)
	}
	return stbRes
}

func LoadLocations(location string) []opentransport.Location {
	locations, err := newClient().Location.Search(context.Background(), location)
	if err != nil {
		fmt.Printf("Failed to load locations for %s: %s\n", location, err)
		os.Exit(1)
	}
	return locations
}

func StationName(station *opentransport.Location) string {
	return fmt.Sprintf("%s (%s)", station.Name, station.Id)
}
