package wails

import (
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func Run(path string, opt *options.App) (err error) {

	if opt.Title != "" {
		opt.Title = filepath.Base(path)
	}

	front := path
	path = path + "/dist"
	//path = path + "/build"

	bundleType, err := bundleType(path)
	if err != nil {
		return
	}

	if opt.AssetServer == nil {
		opt.AssetServer = &assetserver.Options{}
	}

	opt.AssetServer.Assets, err = BundleFS(bundleType, path)
	if err != nil {
		return
	}

	store, err := BundleStore(bundleType, path)
	if err != nil {
		return
	}
	opt.AssetServer.Handler = StoreHandler{store}

	log.Println("running wails")

	stopDevWatcher, url, _, err := runFrontendDevWatcherCommand(front, "npm run dev", true)
	if err != nil {
		return err
	}
	log.Println("url: ", url)
	go func() {
		quitChannel := make(chan os.Signal, 1)
		signal.Notify(quitChannel, os.Interrupt, syscall.SIGTERM)
		<-quitChannel
		stopDevWatcher()
	}()
	os.Setenv("devserver", "localhost:34115")
	os.Setenv("assetdir", path)
	os.Setenv("frontenddevserverurl", url)

	return wails.Run(opt)
}
