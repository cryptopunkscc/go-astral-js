//go:build dev

package main

import (
	"github.com/cryptopunkscc/go-astral-js/wails/build"
	"github.com/cryptopunkscc/go-astral-js/wails/bundle"
	"github.com/cryptopunkscc/go-astral-js/wails/create"
	"github.com/cryptopunkscc/go-astral-js/wails/create/templates"
	"github.com/cryptopunkscc/go-astral-js/wails/dev"
	"github.com/leaanthony/clir"
	"github.com/pterm/pterm"
	"log"
)

func main() {
	cli := clir.NewCli("AstralJS", "JavaScript development environment for Astral.", "0.0.1")
	cli.NewSubCommandFunction("init", "Create production bundle.", cliInit)
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

func cliBundle(f *FlagsBundle) error {
	return bundle.Create(f.Path)
}

type FlagsInit struct {
	Template string `name:"t" description:"Name of built-in template to use, path to template or template url"`
	Name     string `name:"n" description:"Name of project"`
	Dir      string `name:"d" description:"Project directory"`
	Force    bool   `name:"f" description:"Force recreate project"`
	List     bool   `name:"l" description:"List available templates"`
}

func cliInit(f *FlagsInit) error {
	if f.List {
		return cliList()
	} else {
		return create.Run(f.Name, f.Dir, f.Template, f.Force)
	}
}

func cliList() error {
	templateList, err := templates.List()
	if err != nil {
		return err
	}

	pterm.DefaultSection.Println("Available templates")

	table := pterm.TableData{{"Template", "Short Name", "Description"}}
	for _, template := range templateList {
		table = append(table, []string{template.Name, template.ShortName, template.Description})
	}
	err = pterm.DefaultTable.WithHasHeader(true).WithBoxed(true).WithData(table).Render()
	pterm.Println()
	return err
}
