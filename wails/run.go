package wails

import "github.com/wailsapp/wails/v2/pkg/options"

type Run func(path string, opt *options.App) error
