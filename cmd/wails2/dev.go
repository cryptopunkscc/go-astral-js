//go:build dev

package main

import (
	"github.com/cryptopunkscc/go-astral-js/wails/build"
	"github.com/cryptopunkscc/go-astral-js/wails/bundle"
	"github.com/cryptopunkscc/go-astral-js/wails/dev"
	"github.com/leaanthony/clir"
	"log"
)

func main() {
	cli := clir.NewCli("AstralJS", "JavaScript development environment for Astral.", "0.0.1")
	cli.NewSubCommandFunction("dev", "Run development server for given dir.", cliDevelopment)
	cli.NewSubCommandFunction("run", "Execute app from bundle, dir, or file.", cliApplication)
	cli.NewSubCommandFunction("build", "Build application.", cliBuild)
	cli.NewSubCommandFunction("bundle", "Create production bundle.", cliBundle)
	if err := cli.Run(); err != nil {
		log.Fatalln(err)
	}
}

type FlagsDev struct{ FlagsPath }

func cliDevelopment(f *FlagsDev) error {
	return dev.Run(f.Path, AppOptions())
}

type FlagsBuild struct{ FlagsPath }

func cliBuild(f *FlagsBuild) error {
	return build.Run(f.Path)
}

type FlagsBundle struct{ FlagsBuild }

func cliBundle(f *FlagsBundle) (err error) {
	return bundle.Create(f.Path)
}
