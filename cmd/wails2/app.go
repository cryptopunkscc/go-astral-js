//go:build !dev

package main

import (
	"github.com/cryptopunkscc/go-astral-js/wails/app"
	"github.com/leaanthony/clir"
	"log"
)

func main() {
	cli := clir.NewCli("AstralJS", "JavaScript runtime environment for Astral.", "0.0.1")
	flags := &FlagsApp{}
	cli.AddFlags(flags)
	cli.Action(func() error { return app.Run(flags.Path, AppOptions()) })
	if err := cli.Run(); err != nil {
		log.Fatalln(err)
	}
}
