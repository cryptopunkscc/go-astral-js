//go:build dev

package main

import (
	"github.com/cryptopunkscc/go-astral-js/pkg/backend/goja"
	"github.com/cryptopunkscc/go-astral-js/pkg/build"
	"github.com/cryptopunkscc/go-astral-js/pkg/bundle"
	"github.com/cryptopunkscc/go-astral-js/pkg/create"
	"github.com/cryptopunkscc/go-astral-js/pkg/create/templates"
	"github.com/cryptopunkscc/go-astral-js/pkg/frontend/wails/dev"
	"github.com/leaanthony/clir"
	"github.com/pterm/pterm"
	"log"
	"path"
	"sync"
)

func main() {
	cli := clir.NewCli(PortalName, PortalDevDescription, PortalVersion)
	cli.NewSubCommandFunction("create", "Create production bundle.", cliInit)
	cli.NewSubCommandFunction("dev", "Run development server for given dir.", cliDevelopment)
	cli.NewSubCommandFunction("open", "Execute app from bundle, dir, or file.", cliApplication)
	cli.NewSubCommandFunction("build", "Build application.", cliBuild)
	cli.NewSubCommandFunction("bundle", "Create production bundle.", cliBundle)
	if err := cli.Run(); err != nil {
		log.Fatalln(err)
	}
}

type FlagsDev struct{ FlagsApp }

func cliDevelopment(f *FlagsDev) (err error) {
	f.Setup()
	wait := sync.WaitGroup{}
	if f.Back {
		wait.Add(1)
		if err = goja.Run(path.Join(f.Path, "src")); err != nil {
			return
		}
	}
	if f.Front {
		wait.Add(1)
		return dev.Run(f.Path, AppOptions())
	}
	wait.Wait()
	return
}

type FlagsBuild struct{ FlagsPath }

func cliBuild(f *FlagsBuild) error {
	return build.Run(f.Path)
}

type FlagsBundle struct{ FlagsBuild }

func cliBundle(f *FlagsBundle) error {
	return bundle.Run(f.Path)
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
