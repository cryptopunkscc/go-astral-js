//go:build dev

package main

import (
	"github.com/cryptopunkscc/go-astral-js/wails"
	"github.com/cryptopunkscc/go-astral-js/wails/dev"
)

var run wails.Run = dev.Run
