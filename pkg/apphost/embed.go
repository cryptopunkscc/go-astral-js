package apphost

import "embed"

//go:embed apphost_base.js
var _jsBase string

func JsBaseString() string { return _jsBase }

//go:embed apphost.js
var _fs embed.FS

func JsFs() embed.FS { return _fs }
