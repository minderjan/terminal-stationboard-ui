package ui

import "github.com/gdamore/tcell"

type Theme struct {
	Bg               tcell.Color
	Fg               tcell.Color
	TblColoredCellBg tcell.Color
	TblColoredCellFg tcell.Color
	TimeOk           tcell.Color
	TimeWarning      tcell.Color
	TimeAlert        tcell.Color
	CmdBtnBg         tcell.Color
	CmdBtnFq         tcell.Color
}

func LoadTheme(name string) Theme {
	switch name {
	case "blue":
		return themeBlue
	case "light":
		return themeLight
	case "dark":
		return themeBlack
	default:
		return themeBlack
	}
}

var themeLight = Theme{
	Bg:               tcell.ColorWhite,
	Fg:               tcell.ColorBlack,
	TblColoredCellBg: tcell.ColorWhite,
	TblColoredCellFg: tcell.ColorDarkCyan,
	TimeOk:           tcell.ColorGreen,
	TimeWarning:      tcell.ColorOrange,
	TimeAlert:        tcell.ColorRed,
	CmdBtnBg:         tcell.ColorDarkCyan,
	CmdBtnFq:         tcell.ColorWhite,
}

var themeBlue = Theme{
	Bg:               tcell.ColorDarkBlue,
	Fg:               tcell.ColorWhite,
	TblColoredCellBg: tcell.ColorDarkBlue,
	TblColoredCellFg: tcell.ColorYellow,
	TimeOk:           tcell.ColorLightGreen,
	TimeWarning:      tcell.ColorLightCoral,
	TimeAlert:        tcell.ColorRed,
	CmdBtnBg:         tcell.ColorDarkCyan,
	CmdBtnFq:         tcell.ColorWhite,
}

var themeBlack = Theme{
	Bg:               tcell.ColorBlack,
	Fg:               tcell.ColorWhite,
	TblColoredCellBg: tcell.ColorBlack,
	TblColoredCellFg: tcell.ColorYellow,
	TimeOk:           tcell.ColorLightGreen,
	TimeWarning:      tcell.ColorLightCoral,
	TimeAlert:        tcell.ColorRed,
	CmdBtnBg:         tcell.ColorDarkCyan,
	CmdBtnFq:         tcell.ColorWhite,
}
