package config

import (
	color "github.com/MeibisuX673/consoleMesenger/tcp-client/config/color"
)

var Flags *Flag

type Flag struct {
	UseColor   string
	Connection string
}

func InitFlags() {

	colorUse := color.GetColor()

	Flags = &Flag{
		UseColor: colorUse,
	}
}

func init() {
	InitFlags()
}
