package goja

import (
	astraljs "github.com/cryptopunkscc/go-astral-js/pkg/apphost"
	"github.com/cryptopunkscc/go-astral-js/pkg/assets"
	"github.com/dop251/goja"
	"io/fs"
	"log"
)

func Run(path string) (err error) {
	// identify app bundle type
	bundleType, err := assets.BundleType(path)
	if err != nil {
		return
	}

	bundleFs, err := assets.BundleFS(bundleType, path)
	if err != nil {
		return
	}

	bytes, err := fs.ReadFile(bundleFs, "service.js")
	if err != nil {
		return err
	}

	RunSource(string(bytes))
	return
}

func RunSource(app string) {

	vm := goja.New()

	err := Bind(vm, astraljs.NewFlatAdapter())
	if err != nil {
		log.Fatal(err)
	}

	// inject apphost client js lib
	_, err = vm.RunString(astraljs.JsBaseString())
	if err != nil {
		log.Fatal(err)
	}

	// start js application backend
	_, err = vm.RunString(app)
	if err != nil {
		log.Fatal(err)
	}
}
