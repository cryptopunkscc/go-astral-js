package main

import (
	astraljs "github.com/cryptopunkscc/go-astral-js"
	"github.com/cryptopunkscc/go-astral-js/wails/app"
	"github.com/wailsapp/wails/v2/pkg/options"
)

func AppOptions() *options.App {
	return &options.App{
		Width:            1024,
		Height:           768,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Bind: []interface{}{
			&Adapter{*astraljs.NewAppHostFlatAdapter()},
		},
	}
}

type Adapter struct{ astraljs.AppHostFlatAdapter }

type FlagsPath struct {
	Path string `pos:"1" default:"."`
}

type FlagsApp struct{ FlagsPath }

func cliApplication(f *FlagsApp) error {
	return app.Run(f.Path, AppOptions())
}
