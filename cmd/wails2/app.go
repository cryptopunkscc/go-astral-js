//go:build !dev

package main

import (
	"github.com/cryptopunkscc/go-astral-js/wails"
	"github.com/cryptopunkscc/go-astral-js/wails/app"
)

var run wails.Run = app.Run
