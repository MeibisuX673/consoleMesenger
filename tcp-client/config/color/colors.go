package color

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var COLOR = map[string]string{
	"black":   "\u001b[30m",
	"red":     "\u001b[31m",
	"green":   "\u001b[32m",
	"yellow":  "\u001b[33m",
	"blue":    "\u001b[34m",
	"magenta": "\u001b[35m",
	"cyan":    "\u001b[36m",
	"white":   "\u001b[37m",
}

func GetColor() string {

	var inputColor string
	var colorUse string

	flag.StringVar(&inputColor, "color", "", "display colorized output")
	flag.Parse()

	if inputColor == "" {
		inputColor = os.Getenv("COLOR_MESSAGE")
	}

	colorUse, ok := COLOR[inputColor]
	if !ok {
		inputColor = "white"
		colorUse = COLOR[inputColor]
	}

	if err := godotenv.Write(
		map[string]string{
			"COLOR_MESSAGE": inputColor,
		},
		"./.env",
	); err != nil {
		log.Fatal(err)
	}

	return colorUse
}
