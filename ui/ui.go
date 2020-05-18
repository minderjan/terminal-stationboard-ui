package ui

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/minderjan/opentransport-client/opentransport"
	"github.com/minderjan/terminal-stationboard-ui/transport"
	"github.com/rivo/tview"
	"os"
	"time"
)

type UI struct {
	app   *tview.Application
	theme Theme

	table  *tview.Table
	main   *tview.Flex
	body   *tview.Frame
	footer *tview.Grid
	header *tview.Grid

	time    *tview.TextView
	station *tview.TextView
	updated *tview.TextView

	stations    []opentransport.Location
	connections []opentransport.StationBoardJourney

	selectedStation *opentransport.Location
	inputStation    string

	nextUpdate time.Time
	busy       bool

	minConnectionDeparture time.Duration
}

func NewUI(input string, station *opentransport.Location) *UI {
	return setupUI(themeBlue, input, station)
}

// Setup a new UI with a specific theme.
func NewUIWithTheme(input string, station *opentransport.Location, minConnectionDeparture time.Duration, theme Theme) *UI {
	ui := setupUI(theme, input, station)
	ui.inputStation = input
	ui.minConnectionDeparture = minConnectionDeparture
	return ui
}

func (u *UI) Run() error {
	return u.changeScreenSafe(u.main, true)
}

func setupUI(theme Theme, input string, station *opentransport.Location) *UI {
	u := &UI{
		app:     tview.NewApplication(),
		theme:   theme,
		table:   tview.NewTable(),
		main:    tview.NewFlex(),
		time:    tview.NewTextView(),
		station: tview.NewTextView(),
		updated: tview.NewTextView(),
	}

	u.inputStation = input
	u.selectedStation = station

	// Screen Size
	u.main.SetFullScreen(true)

	// Set Background Color
	u.main.SetBackgroundColor(u.theme.Bg)
	u.table.SetBackgroundColor(u.theme.Bg)

	// Set Foreground Colors
	u.time.SetTextColor(u.theme.Fg)
	u.station.SetTextColor(u.theme.Fg)
	u.updated.SetTextColor(u.theme.Fg)

	// Configure Text Labels
	u.time.SetTextAlign(tview.AlignRight).SetBackgroundColor(u.theme.Bg)
	u.station.SetTextAlign(tview.AlignLeft).SetBackgroundColor(u.theme.Bg)
	u.updated.SetTextAlign(tview.AlignRight).SetBackgroundColor(u.theme.Bg)

	// Setup Table
	u.table.SetBorder(false)
	u.table.SetTitle("Next connections")
	u.station.SetText(transport.StationName(station))

	// Add Widgets to Layout
	u.main.SetDirection(tview.FlexRow).
		AddItem(u.addHeader(u.station, u.time), 3, 0, false).
		AddItem(u.addBody(u.table), 0, 65, false).
		AddItem(u.addFooter(u.updated), 1, 0, false)

	u.registerEventHandlers()

	return u
}

// ----------- App

func (u *UI) registerEventHandlers() {

	inputHandler := func(event *tcell.EventKey) *tcell.EventKey {

		switch event.Key() {
		case tcell.KeyEscape:
			u.exit()
		case tcell.KeyCtrlC:
			u.exit()
		}

		switch event.Rune() {
		case 'q':
			u.exit()
			return nil
		case 'c':
			u.ChangeStation()
			return nil
		}

		return event
	}

	u.app.SetInputCapture(inputHandler)
}

func (u *UI) exit() {
	u.app.Stop()
	os.Exit(0)
}

func (u *UI) changeScreenSafe(root tview.Primitive, fullscreen bool) error {
	if err := u.app.SetRoot(root, fullscreen).SetFocus(root).Run(); err != nil {
		return err
	}
	u.registerEventHandlers()
	return nil
}

func (u *UI) changeScreen(root tview.Primitive, fullscreen bool) {
	if err := u.changeScreenSafe(root, fullscreen); err != nil {
		panic(err)
	}
}

// ----------- Layout

func (u *UI) addBody(tableRef *tview.Table) *tview.Frame {
	u.body = tview.NewFrame(tableRef)
	u.body.SetBorder(false)
	u.body.SetBorderPadding(0, 0, 1, 1)
	u.body.SetBackgroundColor(u.theme.Bg)
	return u.body
}

func (u *UI) addHeader(stationRef *tview.TextView, timeRef *tview.TextView) *tview.Grid {
	u.header = tview.NewGrid().SetBorders(false).
		AddItem(stationRef, 0, 0, 1, 1, 0, 0, false).
		AddItem(timeRef, 0, 1, 1, 1, 0, 0, false)
	u.header.SetBorderPadding(1, 1, 1, 1)
	u.header.SetBackgroundColor(u.theme.Bg)
	return u.header
}

func (u *UI) addFooter(updates *tview.TextView) *tview.Grid {
	btnQuit := tview.NewButton("(q) quit").SetSelectedFunc(func() {
		u.app.Stop()
	})

	btnUpdate := tview.NewButton("(c) change").SetSelectedFunc(func() {
		u.app.Stop()
	})

	// Button Theme
	btnUpdate.SetBackgroundColor(u.theme.CmdBtnBg)
	btnQuit.SetBackgroundColor(u.theme.CmdBtnBg)

	cmds := tview.NewGrid().SetBorders(false).
		SetColumns(15, 16, -2).
		AddItem(btnQuit, 0, 0, 1, 1, 0, 0, false).
		AddItem(btnUpdate, 0, 1, 1, 1, 0, 0, false).
		AddItem(tview.NewBox().SetBackgroundColor(u.theme.Bg), 0, 3, 1, 1, 0, 0, false)

	u.footer = tview.NewGrid().SetBorders(false).
		AddItem(cmds, 0, 0, 1, 1, 0, 0, false).
		AddItem(updates, 0, 1, 1, 1, 0, 0, false)

	// Layout Theme
	u.footer.SetBackgroundColor(u.theme.Bg)
	cmds.SetBackgroundColor(u.theme.Bg)
	return u.footer
}

// ----------- Table

func (u *UI) tableHeader(values ...string) {
	for i, v := range values {
		u.cell(0, i, v)
	}
}

func (u *UI) cell(r, c int, text string) {
	u.table.SetCell(r, c,
		tview.NewTableCell(text).
			SetTextColor(u.theme.Fg).
			SetAlign(tview.AlignLeft))
}

func (u *UI) coloredCell(r, c int, text string) {
	u.table.SetCell(r, c,
		tview.NewTableCell(text).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft).
			SetBackgroundColor(u.theme.TblColoredCellBg).
			SetTextColor(u.theme.TblColoredCellFg))
}

func (u *UI) timeCell(r, c int, date time.Time) {

	colFg := u.theme.Fg
	diff := time.Until(date)
	diffMinutes := diff.Minutes()

	switch {
	case diffMinutes <= 5:
		colFg = u.theme.TimeAlert
		break
	case diffMinutes <= 7:
		colFg = u.theme.TimeWarning
		break
	default:
		colFg = u.theme.TimeOk
		break
	}

	u.table.SetCell(r, c,
		tview.NewTableCell(fmt.Sprintf("%3v ", fmt.Sprintf("%.0fm", diffMinutes))).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft).
			SetBackgroundColor(u.theme.Bg).
			SetTextColor(colFg))
}

// ----------- Updates

// Updates the time in upper right corner.
func (u *UI) UpdateTime(d time.Duration) {
	u.time.SetText(fmt.Sprintf("%s", time.Now().Format("15:04:05")))
	for {
		time.Sleep(d)
		u.app.QueueUpdateDraw(func() {
			u.time.SetText(fmt.Sprintf("%s", time.Now().Format("15:04:05")))
		})
	}
}

// Updates the update indicator in bottom right corner.
func (u *UI) UpdateIndicator() {
	for {
		time.Sleep(500 * time.Millisecond)
		if !u.busy {
			u.app.QueueUpdate(func() {

				minutes := time.Until(u.nextUpdate).Minutes()
				seconds := time.Until(u.nextUpdate).Seconds()

				var dateString string
				if int64(minutes) > 0 {
					dateString = fmt.Sprintf("%.0fm", minutes)
				} else {
					dateString = fmt.Sprintf("%.0fs", seconds)
				}

				u.updated.SetText(fmt.Sprintf("update in %s", dateString))
			})
		} else {
			u.updated.SetText(fmt.Sprintf("updating..."))
		}
	}
}

// Updates the stationboard data with an interval.
func (u *UI) Update(interval time.Duration) {
	u.nextUpdate = time.Now().Add(interval) // initial loading
	for {
		time.Sleep(interval)
		u.app.QueueUpdate(func() {
			u.updateStationboard()
			u.nextUpdate = time.Now().Add(interval)
		})
	}
}

func (u *UI) updateStationboard() {
	u.busy = true
	stbRes := transport.LoadStationboard(u.selectedStation.Id, u.minConnectionDeparture)
	u.table.Clear()
	u.AddStationboard(stbRes.Journeys)
	u.busy = false
}

// Removes Connections from Stationboard, which has been in past.
func (u *UI) UpdateStationboardTime(interval time.Duration) {
	for {
		time.Sleep(interval)
		u.busy = true
		u.table.Clear()

		// Remove unneeded connections
		var newConnections []opentransport.StationBoardJourney
		for _, c := range u.connections {
			if time.Until(c.Stop.Departure.Time) > u.minConnectionDeparture {
				newConnections = append(newConnections, c)
			}
		}
		u.AddStationboard(newConnections)
		u.busy = false
	}
}

// ---------- Add Data to UI

// Adds a list of stationboard journeys to the cache.
func (u *UI) AddStationboard(connections []opentransport.StationBoardJourney) {
	u.connections = connections
	showPlatform := transport.ShowPlatformCol(connections)
	if showPlatform {
		u.tableHeader("", " To", "Number", fmt.Sprintf("%10v", "Platform"))
	} else {
		u.tableHeader("", " To", "Number")
	}

	for i, c := range connections {
		i += 1 // table header offset
		u.timeCell(i, 0, c.Stop.Departure.Time)
		u.cell(i, 1, fmt.Sprintf(" %-*v", transport.DestinationLength(connections)+5, c.To))
		u.coloredCell(i, 2, fmt.Sprintf("%s", transport.TransportNumber(c.Category, c.Number)))

		if showPlatform {
			u.cell(i, 3, fmt.Sprintf("%10v", c.Stop.Platform))
		}
	}
}

// Adds a list of locations to the cache.
func (u *UI) AddLocations(locations []opentransport.Location) {
	for _, l := range locations {
		if l.Station() {
			u.stations = append(u.stations, l)
		}
	}
}

// ----------- Change

// Changes the location to a different one.
func (u *UI) SelectStation(location string) {
	u.inputStation = location

	// Get Stationboard Data
	stbRes := transport.LoadStationboard(location, u.minConnectionDeparture)
	u.AddStationboard(stbRes.Journeys)

	u.selectedStation = &stbRes.Station
	u.inputStation = u.selectedStation.Id
	u.station.SetText(transport.StationName(u.selectedStation))
}

// Opens a Modal Dialog to change the station.
func (u *UI) ChangeStation() {
	modal := tview.NewModal()

	var sButtons []string
	for _, s := range u.stations {
		sButtons = append(sButtons, s.Name)
	}

	if len(sButtons) > 1 {
		modal.SetText("Select a different location")
		modal.AddButtons(sButtons[:3])
		modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			u.SelectStation(buttonLabel)
			u.changeScreen(u.main, true)
		})
	} else {
		modal.AddButtons([]string{"cancel"})
		modal.SetText(fmt.Sprintf("Only one station found for: '%s'", u.inputStation))
		modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "cancel" {
				u.changeScreen(u.main, true)
			}
		})
	}

	u.changeScreen(modal, false)
}
