package main

import (
	"context"
	"github.com/cryptopunkscc/go-astral-js"
	"github.com/cryptopunkscc/go-astral-js/goja"
)

func main() {
	app := astraljs.ResolveWebApp()

	goja.RunSource(app.Source)

	ctx := context.Background()
	<-ctx.Done()
}
