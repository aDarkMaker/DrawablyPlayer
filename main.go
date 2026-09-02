package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var appIcon []byte

func main() {
	app := application.New(application.Options{
		Name:        "DrawablyPlayer",
		Description: "Drawably style music & video player",
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Icon: appIcon,
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:            "DrawablyPlayer",
		Width:            1280,
		Height:           800,
		URL:              "/",
		BackgroundColour: application.NewRGB(18, 18, 24),
	})

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
