package main

import (
	"fmt"
	astraljs "github.com/cryptopunkscc/go-astral-js"
	"github.com/wailsapp/wails/v2/pkg/options"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <app>\n", os.Args[0])
		os.Exit(0)
	}

	path := os.Args[1]
	opt := options.App{
		Width:            1024,
		Height:           768,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Bind: []interface{}{
			&Adapter{*astraljs.NewAppHostFlatAdapter()},
		},
	}

	if err := run(path, &opt); err != nil {
		log.Fatalln(err)
	}
}

type Adapter struct{ astraljs.AppHostFlatAdapter }
