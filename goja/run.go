package goja

import (
	astraljs "github.com/cryptopunkscc/go-astral-js"
	"github.com/cryptopunkscc/go-astral-js/wails/assets"
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

	err := Bind(vm, astraljs.NewAppHostFlatAdapter())
	if err != nil {
		log.Fatal(err)
	}

	// inject apphost client js lib
	_, err = vm.RunString(astraljs.AppHostJsClient())
	if err != nil {
		log.Fatal(err)
	}

	// start js application backend
	_, err = vm.RunString(app)
	if err != nil {
		log.Fatal(err)
	}
}
