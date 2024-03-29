package main

import (
	astraljs "github.com/cryptopunkscc/go-astral-js/pkg/apphost"
	"github.com/cryptopunkscc/go-astral-js/pkg/clir"
	"io"
)

func main() {
	clir.Run(func() io.Closer {
		return &Adapter{FlatAdapter: *astraljs.NewFlatAdapter()}
	})
}

type Adapter struct{ astraljs.FlatAdapter }
